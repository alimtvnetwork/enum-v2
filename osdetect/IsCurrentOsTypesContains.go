package osdetect

func IsCurrentOsTypesContains(items ...Variant) bool {
	currentTypesMap := CurrentOsTypesMap()
	
	for _, item := range items {
		_, has := currentTypesMap[item]
		
		if has {
			return true
		}
	}
	
	return false
}
