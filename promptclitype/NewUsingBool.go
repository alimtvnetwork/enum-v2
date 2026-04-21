package promptclitype

func NewUsingBool(isYes bool) Variant {
	if isYes {
		return Accept
	}

	return Reject
}
