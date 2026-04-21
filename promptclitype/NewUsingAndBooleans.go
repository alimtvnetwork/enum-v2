package promptclitype

func NewUsingAndBooleans(allBooleans ...bool) Variant {
	for _, isYes := range allBooleans {
		if !isYes {
			return Reject
		}
	}

	return Accept
}
