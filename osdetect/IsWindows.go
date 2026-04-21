package osdetect

func IsWindows11() bool {
	return CurrentOsType().IsWindows11()
}

func IsWindows10() bool {
	osType := CurrentOsType()
	osDetail, err := osType.OsDetailWithError()
	isDetectFail := err != nil ||
		osDetail == nil ||
		osDetail.IsEmptyWindowsDetail()
	
	if isDetectFail {
		return false
	}
	
	return osDetail.IsWindows() &&
		osDetail.WindowsDetail.IsWindows10()
}

func IsWindows8() bool {
	osType := CurrentOsType()
	osDetail, err := osType.OsDetailWithError()
	isDetectFail := err != nil ||
		osDetail == nil ||
		osDetail.IsEmptyWindowsDetail()
	
	if isDetectFail {
		return false
	}
	
	return osDetail.IsWindows() &&
		osDetail.WindowsDetail.IsWindows8()
}
