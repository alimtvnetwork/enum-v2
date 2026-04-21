package osarchs

func Is(
	rawString string,
	expected Architecture,
) bool {
	conv, err := New(rawString)

	if err != nil {
		return false
	}

	return conv == expected
}
