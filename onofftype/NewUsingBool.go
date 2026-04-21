package onofftype

func NewUsingBool(isYes bool) Variant {
	if isYes {
		return On
	}

	return Off
}
