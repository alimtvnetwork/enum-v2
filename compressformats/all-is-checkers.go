package compressformats

func Is(
	rawString string,
	expectedEnum Variant,
) bool {
	ipConv, err := New(rawString)

	if err != nil {
		return false
	}

	return ipConv == expectedEnum
}
