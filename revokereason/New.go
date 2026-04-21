package revokereason

func New(name string) (Variant, error) {
	val, err := BasicEnumImpl.GetValueByName(
		name)

	if err != nil {
		return Unspecified, err
	}

	return Variant(val), nil
}
