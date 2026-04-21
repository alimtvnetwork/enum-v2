package osdetect

func IsCentOs() bool {
	return CurrentOsType().
		IsCentos()
}
