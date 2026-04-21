package osgroupexecution

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	ranges = [...]string{
		Invalid:            "Invalid",
		Create:             "Create",
		Delete:             "Delete",
		Update:             "Update",
		ManageByUsers:      "ManageByUsers",
		AddGroupsToSudoers: "AddGroupsToSudoers",
		GroupManage:        "GroupManage",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Create),
		ranges[:])
)
