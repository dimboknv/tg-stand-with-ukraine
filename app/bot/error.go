package bot

import "fmt"

type userError struct {
	Err     error
	UserMsg string
}

func (u *userError) Error() string {
	// nolint:gocritic // u.Err can be nil
	return fmt.Sprint(u.Err)
}
