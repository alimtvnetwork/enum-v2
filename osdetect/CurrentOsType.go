package osdetect

func CurrentOsType() Variant {
	currentOsType := currentOsMixTypeOnce.
		Value()
	
	return Variant(currentOsType)
}
