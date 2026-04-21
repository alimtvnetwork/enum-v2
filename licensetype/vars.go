package licensetype

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:      "Invalid",
		PublicDomain: "PublicDomain",
		ByCc:         "ByCc",
		BySa:         "BySa",
		ByNc:         "ByNc",
		ByNcSa:       "ByNcSa",
		ByNd:         "ByNd",
		ByNcNd:       "ByNcNd",
	}

	RangesMap = map[string]Variant{
		"Invalid":      Invalid,
		"PublicDomain": PublicDomain,
		"ByCc":         ByCc,
		"BySa":         BySa,
		"ByNc":         ByNc,
		"ByNcSa":       ByNcSa,
		"ByNd":         ByNd,
		"ByNcNd":       ByNcNd,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
