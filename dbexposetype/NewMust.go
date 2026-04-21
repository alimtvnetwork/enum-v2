package dbexposetype

import "gitlab.com/auk-go/core/errcore"

func NewMust(name string) Variant {
	exitCode, err := New(name)
	errcore.HandleErr(err)

	return exitCode
}
