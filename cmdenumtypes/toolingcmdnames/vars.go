package toolingcmdnames

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:       "Invalid",
		Help:          "Help",
		Log:           "Log",
		Update:        "Update",
		Upgrade:       "Upgrade",
		AutoFix:       "AutoFix",
		ImportAutoFix: "ImportAutoFix",
		Automate:      "Automate",
		Backup:        "Backup",
		HealthChecker: "HealthChecker",
		AutoUpdater:   "AutoUpdater",
		Import:        "Import",
		Export:        "Export",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.DefaultAllCases(
		Invalid,
		Ranges[:])
)
