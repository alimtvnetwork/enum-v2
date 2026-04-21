package osdetect

func IsRedhat() bool {
	return CurrentOsType().
		IsRedHatEnterpriseLinux()
}
