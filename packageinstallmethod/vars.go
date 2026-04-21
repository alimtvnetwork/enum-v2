package packageinstallmethod

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:       "Invalid",
		Url:           "Url",
		OsPackages:    "OsPackages",
		AdvanceScript: "AdvanceScript",
	}

	RangesMap = map[string]Variant{
		"Invalid":       Invalid,
		"Url":           Url,
		"OsPackages":    OsPackages,
		"AdvanceScript": AdvanceScript,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
