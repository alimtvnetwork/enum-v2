package timeunit

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:     "Invalid",
		Millisecond: "Millisecond",
		Second:      "Second",
		Minute:      "Minute",
		Hour:        "Hour",
		Day:         "Day",
		Month:       "Month",
		Year:        "Year",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
