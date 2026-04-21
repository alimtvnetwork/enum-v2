package dbexposetype

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:    "Invalid",
		AnyIp:      "AnyIp",
		SpecificIp: "SpecificIp",
	}

	RangesMap = map[string]Variant{
		"Invalid":    Invalid,
		"AnyIp":      AnyIp,
		"SpecificIp": SpecificIp,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
