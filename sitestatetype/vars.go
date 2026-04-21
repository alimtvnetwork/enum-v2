package sitestatetype

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:    "Invalid",
		NewlyAdded: "NewlyAdded",
		Removed:    "Removed",
		Unchanged:  "Unchanged",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.Default(
		Invalid,
		Ranges[:])
)
