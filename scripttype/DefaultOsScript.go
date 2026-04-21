package scripttype

import "gitlab.com/auk-go/core/osconsts"

func DefaultOsScript() *ScriptDefault {
	if osconsts.IsWindows {
		return cmdDefaultScript
	}

	return bashDefaultScript
}
