package strtype

const (
	CurlyStringWrapFormat                                = "{%s}"
	BracketStringWrapFormat                              = "[%s]"
	TitleValueQuotationParenthesisRefWrapReferenceFormat = "%s: %q (%s)" // Title, QuotationValue, Reference string
	TitleBracketWrapFormat                               = "%s: [%s]"    // Title, Value - (value bracket wrapped)
	TitleCurlyWrapFormat                                 = "%s: {%s}"    // Title, Value - (value curly wrapped)
	TitleQuotationWrapFormat                             = "%s: %q"      // Title, Value - (value Quotation wrapped)
)
