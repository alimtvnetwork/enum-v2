package scripttype

import "gitlab.com/auk-go/core/osconsts"

func OsDefaultScriptType() Variant {
	if osconsts.IsWindows {
		return Powershell
	}

	return Bash
}
