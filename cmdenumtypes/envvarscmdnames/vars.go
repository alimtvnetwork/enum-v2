package envvarscmdnames

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:       "Invalid",
		Help:          "Help",
		AppendPath:    "AppendPath",
		Set:           "Set",
		Remove:        "Remove",
		TempSet:       "TempSet",
		TempRemove:    "TempRemove",
		Source:        "Source",
		HasIssues:     "HasIssues",
		List:          "List",
		ListJson:      "ListJson",
		ListPaths:     "ListPaths",
		ListPathsJson: "ListPathsJson",
		Backup:        "Backup",
		Import:        "Import",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
