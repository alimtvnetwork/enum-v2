package conntrackstate

func Is(
	rawStr string,
	expectedVersion Variant,
) bool {
	convData, err := Create(rawStr)

	if err != nil {
		return false
	}

	return convData == expectedVersion
}
