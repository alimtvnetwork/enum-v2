package servicestate

import "gitlab.com/auk-go/core/errcore"

func NewMust(name string) Action {
	exitCode, err := New(name)
	errcore.HandleErr(err)

	return exitCode
}
