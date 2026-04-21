package inttype

//goland:noinspection ALL
func GetSetVariant(
	isCondition bool,
	trueValue int,
	falseValue int,
) Variant {
	if isCondition {
		return Variant(trueValue)
	}

	return Variant(falseValue)
}
