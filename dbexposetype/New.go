package dbexposetype

func New(name string) (Variant, error) {
	v, err := BasicEnumImpl.GetValueByName(name)

	if err != nil {
		return Invalid, err
	}

	return Variant(v), nil
}
