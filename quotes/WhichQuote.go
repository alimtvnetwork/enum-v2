package quotes

func WhichQuote(char byte, isLeft bool) QuoteStatus {
	otherQuoteStatus, has := otherQuoteCharsMaps[char]

	if !has {
		return EmptyQuoteStatus()
	}

	otherQuoteStatus.IsLeft = isLeft

	return otherQuoteStatus
}
