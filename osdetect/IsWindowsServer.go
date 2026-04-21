package osdetect

func IsWindowsServer() bool {
	osType := CurrentOsType()
	osDetail, err := osType.OsDetailWithError()
	isDetectFail := err != nil ||
		osDetail == nil ||
		osDetail.IsEmptyWindowsDetail()
	
	if isDetectFail {
		return false
	}
	
	return osDetail.IsWindows() &&
		osDetail.WindowsDetail.IsWindowsSever()
}

func IsWindowsServer2016() bool {
	osType := CurrentOsType()
	osDetail, err := osType.OsDetailWithError()
	isDetectFail := err != nil ||
		osDetail == nil ||
		osDetail.IsEmptyWindowsDetail()
	
	if isDetectFail {
		return false
	}
	
	return osDetail.IsWindows() &&
		osDetail.WindowsDetail.IsWindowsSever2016()
}

func IsWindowsServer2019() bool {
	osType := CurrentOsType()
	osDetail, err := osType.OsDetailWithError()
	isDetectFail := err != nil ||
		osDetail == nil ||
		osDetail.IsEmptyWindowsDetail()
	
	if isDetectFail {
		return false
	}
	
	return osDetail.IsWindows() &&
		osDetail.WindowsDetail.IsWindowsSever2019()
}
