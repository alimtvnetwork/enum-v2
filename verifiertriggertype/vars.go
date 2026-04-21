package verifiertriggertype

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:           "Invalid",
		AllComplete:       "AllComplete",
		AfterRestart:      "AfterRestart",
		AfterNetworkReset: "AfterNetworkReset",
	}

	RangesMap = map[string]Variant{
		"Invalid":           Invalid,
		"AllComplete":       AllComplete,
		"AfterRestart":      AfterRestart,
		"AfterNetworkReset": AfterNetworkReset,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
