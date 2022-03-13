package hub

import "fmt"

type ResolveURLErr struct {
	Err error
	URL string
}

func (r *ResolveURLErr) Error() string {
	// nolint:gocritic // r.Err can be nil
	return fmt.Sprint(r.Err)
}
