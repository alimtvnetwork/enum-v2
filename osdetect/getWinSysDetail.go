package osdetect

import (
	"github.com/alimtvnetwork/core-v9/coredata/corejson"
	"github.com/alimtvnetwork/core-v9/errcore"
	"github.com/alimtvnetwork/core-v9/osconsts"
	"github.com/alimtvnetwork/core-v9/ostype"
)

func getWinSysDetail() (windowsSystemDetailGetter, error) {
	if osconsts.IsWindows {
		return NewWindowsSystemDetailGetter()
	}
	
	return nil, errcore.NotSupportedType.Error(
		"Not supported other than windows system",
		corejson.NewPtr(ostype.CurrentGroupVariant).JsonString())
}
