package bot

import (
	"github.com/dimboknv/tg-stand-with-ukraine/app/store"
)

func (b *Bot) registerHandlers() {
	b.msgHandlers = map[store.Navigation]handler{
		store.UserNav:    b.userNavigation,
		store.Pass2faNav: b.pass2faNavigation,
		store.CodeNav:    b.codeNavigation,
		store.PhoneNav:   b.phoneNavigation,
	}

	b.cmdHandlers = map[string]handler{
		store.StartCommand:  b.handleStartCommand,
		store.LoginCommand:  b.handleLoginCommand,
		store.LogoutCommand: b.handleLogoutCommand,
	}
}
