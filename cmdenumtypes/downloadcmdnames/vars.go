package downloadcmdnames

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:            "Invalid",
		Help:               "Help",
		Log:                "Log",
		Status:             "Status",
		Install:            "Install",
		Uninstall:          "Uninstall",
		To:                 "To",
		Temp:               "Temp",
		Decompress:         "Decompress",
		TempDecompress:     "TempDecompress",
		Verify:             "Verify",
		DownloadVerify:     "DownloadVerify",
		Schedule:           "Schedule",
		ScheduleTemp:       "ScheduleTemp",
		ScheduleDecompress: "ScheduleDecompress",
		List:               "List",
		ListJson:           "ListJson",
		Backup:             "Backup",
		Import:             "Import",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
