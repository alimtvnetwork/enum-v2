package dnscmdnames

import "gitlab.com/auk-go/core/errcore"

func NewMust(name string) Variant {
	newType, err := New(name)
	errcore.HandleErr(err)

	return newType
}
