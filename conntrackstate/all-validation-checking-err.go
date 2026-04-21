package conntrackstate

func ValidationError(
	rawString string,
	expectedEnum Variant,
) error {
	return BasicEnumImpl.ExpectingEnumValueError(
		rawString,
		expectedEnum)
}
