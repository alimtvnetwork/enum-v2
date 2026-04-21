package osdetect

import (
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/osconsts"
	"gitlab.com/auk-go/core/ostype"
)

func getWinSysDetail() (windowsSystemDetailGetter, error) {
	if osconsts.IsWindows {
		return NewWindowsSystemDetailGetter()
	}
	
	return nil, errcore.NotSupportedType.Error(
		"Not supported other than windows system",
		corejson.NewPtr(ostype.CurrentGroupVariant).JsonString())
}
