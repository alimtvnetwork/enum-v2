package certaction

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid: "Invalid",
		Create:  "Create",
		Renew:   "Renew",
		Revoke:  "Revoke",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.Default(
		Invalid,
		Ranges[:])
)
