package onofftype

func newOtherWays(name string, err error) (Variant, error) {
	val, has := nameToVariant[name]

	if !has {
		return Invalid, err
	}

	return Variant(val.ValueByte()), nil
}
