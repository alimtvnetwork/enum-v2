package onofftype

func NewUsingAndBooleans(allBooleans ...bool) Variant {
	for _, isYes := range allBooleans {
		if !isYes {
			return Off
		}
	}

	return On
}
