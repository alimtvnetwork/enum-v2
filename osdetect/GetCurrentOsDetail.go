package osdetect

import (
	"gitlab.com/auk-go/core/codestack"
	"gitlab.com/auk-go/core/errcore"
)

func GetCurrentOsDetail() (*OperatingSystemDetail, error) {
	osDetailWithErr := currentOsDetailGeneratorOnce.
		Value().(*OsDetailWithErr)
	
	if osDetailWithErr != nil {
		return osDetailWithErr.OperatingSystemDetail, errcore.ToError(osDetailWithErr.Error)
	}
	
	return nil, errcore.NotSupportedType.Error(
		"couldn't cast or generate os details!",
		codestack.StacksStringDefault())
}
