package envpathcmdnames

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:    "Invalid",
		Help:       "Help",
		Append:     "Append",
		Remove:     "Remove",
		TempAppend: "TempAppend",
		TempRemove: "TempRemove",
		Source:     "Source",
		Fix:        "Fix",
		OrderFix:   "OrderFix",
		HasIssues:  "HasIssues",
		List:       "List",
		ListJson:   "ListJson",
		Backup:     "Backup",
		Import:     "Import",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
