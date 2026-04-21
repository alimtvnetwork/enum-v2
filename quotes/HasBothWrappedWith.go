package quotes

func HasBothWrappedWith(str string, quote Quote) bool {
	if len(str) <= 1 {
		return false
	}

	// has at least 2 chars
	length := len(str)
	firstChar := str[0]
	lastChar := str[length-1]

	return quote.IsEqual(firstChar) && quote.IsEqual(lastChar)
}
