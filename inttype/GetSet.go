package inttype

//goland:noinspection ALL
func GetSet(
	isCondition bool,
	trueValue Variant,
	falseValue Variant,
) Variant {
	if isCondition {
		return trueValue
	}

	return falseValue
}
