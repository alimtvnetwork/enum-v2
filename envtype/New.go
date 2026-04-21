package envtype

func New(name string) (Variant, error) {
	val, err := BasicEnumImpl.GetValueByName(
		name)

	if err != nil {
		return Uninitialized, err
	}

	return Variant(val), nil
}
