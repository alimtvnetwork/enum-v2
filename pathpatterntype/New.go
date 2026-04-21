package pathpatterntype

// New
//
//  Variant gets created from Variant JSON name direct name or
//  curly name or path name also returns the variant.
//
// Example:
//  - "Id" or "\"Id\"" or {id}
//      or id or idValue as string("5") : should return Id
func New(name string) (Variant, error) {
	v, err := BasicEnumImpl.GetValueByName(name)

	if err != nil {
		return findUsingInternalMapping(name, err)
	}

	return Variant(v), nil
}
