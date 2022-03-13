package reporter

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/dimboknv/tg-stand-with-ukraine/app/hub"
	"github.com/dimboknv/tg-stand-with-ukraine/app/store"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Reporter struct {
	db                   store.Store
	hub                  *hub.Hub
	log                  *zap.Logger
	unresolvedRashistsCh chan store.Rashist
	message              string
	intervalMaxReps      int
	interval             time.Duration
}

type Opts struct {
	DB                 store.Store
	Hub                *hub.Hub
	Logger             *zap.Logger
	Message            string
	Interval           time.Duration
	IntervalMaxReports int
}

func New(opts Opts) *Reporter {
	rand.Seed(time.Now().UnixNano())
	return &Reporter{
		db:                   opts.DB,
		hub:                  opts.Hub,
		log:                  opts.Logger,
		message:              opts.Message,
		interval:             opts.Interval,
		intervalMaxReps:      opts.IntervalMaxReports,
		unresolvedRashistsCh: make(chan store.Rashist, 1),
	}
}

func (r *Reporter) clientSendReports(ctx context.Context, user store.User, phone string, rashists []store.Rashist) []store.Report {
	max := r.intervalMaxReps
	reports := make([]store.Report, 0)

LOOP:
	for i := range rashists {
		url := rashists[i].URL
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
		if max == 0 {
			break
		}
		if _, has := user.Clients[phone].SentReports[url]; has {
			continue
		}

		rep, err := r.hub.SendReport(ctx, user, phone, url, r.message)
		if err != nil {
			r.log.Error("failed to send report report",
				zap.Int64("userID", user.ID), zap.String("phone", phone), zap.String("url", url), zap.Error(err))
			resolveErr := &hub.ResolveURLErr{}
			switch {
			case errors.Is(err, hub.NotAuthorizedErr):
				break LOOP
			case errors.As(err, &resolveErr):
				r.unresolvedRashistsCh <- store.Rashist{
					URL:        resolveErr.URL,
					ResolveErr: resolveErr.Error(),
				}
			}
			continue
		}
		max--
		reports = append(reports, rep)
		r.log.Info("success to send report report", zap.Int64("userID", user.ID), zap.String("phone", phone), zap.Any("report", rep))
		// nolint:gosec // use rand just for jitter
		time.Sleep(10*time.Second + time.Duration(rand.Int63n(int64(5*time.Second))))
	}
	return reports
}

func (r *Reporter) userSendReports(ctx context.Context, user store.User, rashists []store.Rashist) error {
	mu, wg, clientsReps := &sync.Mutex{}, &sync.WaitGroup{}, map[string]*store.Client{}
	wg.Add(len(user.Clients))

	// send report from all user`s clients and collect sent reports
	for phone := range user.Clients {
		mu.Lock()
		clientsReps[phone] = &store.Client{SentReports: map[string]store.Report{}}
		mu.Unlock()
		go func(phone string) {
			defer wg.Done()
			reps := r.clientSendReports(ctx, user, phone, rashists)

			mu.Lock()
			if clientsReps[phone].SentReports == nil {
				clientsReps[phone].SentReports = map[string]store.Report{}
			}
			for i := range reps {
				clientsReps[phone].SentReports[reps[i].URL] = reps[i]
			}
			mu.Unlock()
		}(phone)
	}
	wg.Wait()

	// update user
	user, err := r.db.GetUser(user.ID)
	if err != nil {
		return err
	}
	for phone := range clientsReps {
		if user.Clients[phone].SentReports == nil {
			user.Clients[phone].SentReports = map[string]store.Report{}
		}
		for url := range clientsReps[phone].SentReports {
			user.Clients[phone].SentReports[url] = clientsReps[phone].SentReports[url]
		}

	}
	return r.db.PutUser(user)
}

func (r *Reporter) sendReports(ctx context.Context) error {
	users, err := r.db.GetUsers()
	if err != nil {
		return errors.Wrap(err, "can`t get users")
	}

	// get rashists and filter by ResolveErr
	rashists, err := r.db.GetRashists()
	if err != nil {
		return errors.Wrap(err, "can`t get rashists")
	}
	noMercy := make([]store.Rashist, 0)
	for i := range rashists {
		if rashists[i].ResolveErr == "" {
			noMercy = append(noMercy, rashists[i])
		}
	}
	rashists = noMercy

	resCh := r.collectUnresolvedRashists()
	wg := &sync.WaitGroup{}
	wg.Add(len(users))
	for _, user := range users {
		go func(user store.User) {
			defer wg.Done()
			if err := r.userSendReports(ctx, user, rashists); err != nil {
				r.log.Error("user failed to send reports", zap.Int64("userID", user.ID), zap.Error(err))
				return
			}
		}(user)
	}
	wg.Wait()
	close(r.unresolvedRashistsCh)
	return r.db.PutRashists(<-resCh)
}

func (r *Reporter) collectUnresolvedRashists() chan []store.Rashist {
	resCh := make(chan []store.Rashist)
	r.unresolvedRashistsCh = make(chan store.Rashist, 1)

	go func() {
		unresolvedRashistsMap := map[string]store.Rashist{}
		for rashist := range r.unresolvedRashistsCh {
			unresolvedRashistsMap[rashist.URL] = rashist
		}
		rashists := make([]store.Rashist, 0, len(unresolvedRashistsMap))
		for key := range unresolvedRashistsMap {
			rashists = append(rashists, unresolvedRashistsMap[key])
		}
		resCh <- rashists
		close(resCh)
	}()

	return resCh
}

func (r *Reporter) AddRashists(_ context.Context, urls []string) error {
	t := time.Now()
	for _, url := range urls {
		_, err := r.db.GetRashist(url)
		if err == nil {
			continue
		}
		if !errors.Is(err, store.NotFoundError) {
			return err
		}

		if err := r.db.PutRashist(store.Rashist{URL: url, CreatedAt: t}); err != nil {
			return err
		}
	}
	return nil
}

func (r *Reporter) Run(ctx context.Context) {
	ticker := time.NewTicker(r.interval)
	defer ticker.Stop()
	r.log.Info("reporter is started", zap.Duration("interval", r.interval))

	for {
		select {
		case <-ticker.C:
			go func() {
				r.log.Info("start reporting...")
				c, cancel := context.WithTimeout(ctx, r.interval-time.Minute)

				err := r.sendReports(c)
				cancel()
				if err != nil {
					r.log.Error("reporting is failed", zap.Error(err))
					return
				}
				r.log.Info("reporting is done")
			}()
		case <-ctx.Done():
			return
		}
	}
}
