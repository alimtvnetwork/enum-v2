package hostingplancmdnames

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:       "Invalid",
		Help:          "Help",
		Add:           "Add",
		Assign:        "Assign",
		AddAssign:     "AddAssign",
		AddOrUpdate:   "AddOrUpdate",
		Update:        "Update",
		Remove:        "Remove",
		RemoveOnExist: "RemoveOnExist",
		List:          "List",
		Search:        "Search",
		Histories:     "Histories",
		StateChange:   "StateChange",
		Backup:        "Backup",
		Import:        "Import",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
