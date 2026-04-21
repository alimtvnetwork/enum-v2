package linescomparetype

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:      "Invalid",
		Equal:        "Equal",
		Less:         "Less",
		LessEqual:    "LessEqual",
		Greater:      "Greater",
		GreaterEqual: "GreaterEqual",
		NotEqual:     "NotEqual",
	}

	RangesMap = map[string]Variant{
		"Invalid":      Invalid,
		"Equal":        Equal,
		"Less":         Less,
		"LessEqual":    LessEqual,
		"Greater":      Greater,
		"GreaterEqual": GreaterEqual,
		"NotEqual":     NotEqual,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
