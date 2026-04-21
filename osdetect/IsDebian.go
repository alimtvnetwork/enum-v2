package osdetect

func IsDebian() bool {
	return CurrentOsType().
		IsDebian()
}
