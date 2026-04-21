package instructiontype

func Is(
	rawString string,
	expected Variant,
) bool {
	conv, err := New(rawString)

	if err != nil {
		return false
	}

	return conv == expected
}
