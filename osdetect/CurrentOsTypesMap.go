package osdetect

func CurrentOsTypesMap() map[Variant]bool {
	casted, isSuccess := currentOsMixTypesMapOnce.
		Value().(map[Variant]bool)
	
	if isSuccess {
		return casted
	}
	
	return nil
}
