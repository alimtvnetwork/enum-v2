package brackets

// UnWrapWith
//
// Note : It doesn't care about in brackets exist in middle of (str),
// it just un-wrap from both sides if brackets are there.
func UnWrapWith(str string, category Category) string {
	if str == "" {
		return str
	}

	if HasBothWrappedWith(str, category) {
		// no need to modify
		return unWrapBoth(str)
	}

	status := getSingleBracketStatus(str)

	if status.IsBracketFound {
		return unWrapSingle(str, status.IsLeft)
	}

	// return as is, there is nothing to modify
	return str
}
