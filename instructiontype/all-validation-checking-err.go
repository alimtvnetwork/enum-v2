package instructiontype

import "gitlab.com/auk-go/core/errcore"

func ValidationError(
	rawString string,
	expectedEnum Variant,
) error {
	return BasicEnumImpl.ExpectingEnumValueError(
		rawString,
		expectedEnum)
}

func StringMustBe(
	rawString string,
	expected Variant,
) {
	err := ValidationError(rawString, expected)
	errcore.MustBeEmpty(err)
}
