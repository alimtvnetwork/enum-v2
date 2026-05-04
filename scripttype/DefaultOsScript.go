package scripttype

import "github.com/alimtvnetwork/core-v9/osconsts"

func DefaultOsScript() *ScriptDefault {
	if osconsts.IsWindows {
		return cmdDefaultScript
	}

	return bashDefaultScript
}
