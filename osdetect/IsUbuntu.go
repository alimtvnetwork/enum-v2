package osdetect

func IsUbuntu() bool {
	return CurrentOsType().IsUbuntu()
}
