package configcmdnames

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:             "Invalid",
		Help:                "Help",
		Log:                 "Log",
		Apply:               "Apply",
		Revert:              "Revert",
		Store:               "Store",
		DockerApply:         "DockerApply",
		ApplyDuringShutdown: "ApplyDuringShutdown",
		ApplyAfterReboot:    "ApplyAfterReboot",
		ApplyAfter:          "ApplyAfter",
		ApplyBefore:         "ApplyBefore",
		Histories:           "Histories",
		Backup:              "Backup",
		Import:              "Import",
		Export:              "Export",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
