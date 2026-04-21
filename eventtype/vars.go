package eventtype

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid: "Invalid",
		Log:     "Log",
		Success: "Success",
		Error:   "Error",
		Failure: "Failure",
		File:    "File",
		Custom:  "Custom",
	}

	ErrorMap = map[Variant]bool{
		Failure: true,
		Error:   true,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Success),
		Ranges[:])
)
