package osdetect

func CurrentOsTypesMustBePresent(items ...Variant) {
	err := CurrentOsTypesNotContainsError(items...)
	
	if err != nil {
		panic(err)
	}
}
