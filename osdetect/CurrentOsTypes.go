package osdetect

func CurrentOsMixTypes() []Variant {
	casted, isSuccess := currentOsMixTypesOnce.
		Value().([]Variant)
	
	if isSuccess {
		return casted
	}
	
	return nil
}
