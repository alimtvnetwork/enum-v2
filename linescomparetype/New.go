package linescomparetype

// New
//
//  Creates string to the type Variant
//
// Mapping (using @nameToVariant):
//   - "":    Invalid,
//   - "-1":  Invalid,
//   - "ask": Ask,
//   - "*":   Ask,
//   - "yes": On,
//   - "1":   On,
//   - "y":   On,
//   - "n":   Off,
//   - "no":  Off,
//   - "Off":  Off,
//   - "0":   Off,
func New(name string) (Variant, error) {
	val, err := BasicEnumImpl.GetValueByName(
		name)

	if err != nil {
		return Invalid, err
	}

	return Variant(val), nil
}
