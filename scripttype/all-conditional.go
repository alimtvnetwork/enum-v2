package scripttype

func CondBool(
	isTrue bool,
	trueScript Variant,
	falseScript Variant,
) Variant {
	if isTrue {
		return trueScript
	}

	return falseScript
}

func CondFunc(
	isTrue bool,
	trueScriptFunc func() Variant,
) Variant {
	if isTrue {
		return trueScriptFunc()
	}

	return Invalid
}
