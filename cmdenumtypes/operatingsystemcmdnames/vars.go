package operatingsystemcmdnames

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:       "Invalid",
		Help:          "Help",
		Log:           "Log",
		Upgrade:       "Upgrade",
		ImportAutoFix: "ImportAutoFix",
		Info:          "Info",
		Update:        "Update",
		AutoFix:       "AutoFix",
		DefaultConfig: "DefaultConfig",
		Install:       "Install",
		ServiceCreate: "ServiceCreate",
		Services:      "Services",
		SourceList:    "SourceList",
		ChangeSource:  "ChangeSource",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(
		Invalid,
		Ranges[:])
)
