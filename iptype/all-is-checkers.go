package iptype

func IsV4(rawIpVersion string) bool {
	return Is(rawIpVersion, V4)
}

func IsV6(rawIpVersion string) bool {
	return Is(rawIpVersion, V6)
}

func Is(
	rawIpVersion string,
	expectedVersion Variant,
) bool {
	ipConv, err := New(rawIpVersion)

	if err != nil {
		return false
	}

	return ipConv == expectedVersion
}
