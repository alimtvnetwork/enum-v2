package quotes

// UnWrapWith
//
// Note : It doesn't care about in quotes exist in middle of (str),
// it just un-wrap from both sides if quotes are there.
func UnWrapWith(str string, quote Quote) string {
	if str == "" {
		return str
	}

	// skip on exist part
	if HasBothWrappedWith(str, quote) {
		return unWrapBoth(str)
	}

	singleQuoteStatus := getQuoteStatus(str)

	if singleQuoteStatus.IsQuoteFound {
		return unWrapSingle(str, singleQuoteStatus.IsLeft)
	}

	// return as is, there is nothing to modify
	return str
}
