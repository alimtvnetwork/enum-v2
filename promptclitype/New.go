package promptclitype

// New
//
//  Creates string to the type Variant
//
// Mapping (using @nameToVariant):
//   - "":       Invalid,
//   - "-1":     Invalid,
//   - "ask":    Later,
//   - "*":      Later,
//   - "yes":    Accept,
//   - "1":      Accept,
//   - "y":      Accept,
//   - "n":      Reject,
//   - "no":     Reject,
//   - "Reject": Reject,
//   - "0":      Reject,
//   - "3":      Review,
//   - "review": Review,
//   - "Review": Review,
//   - "A":      Accept,
//   - "R":      Reject,
//   - "L":      Later,
//   - "C":      Review,
//   - "a":      Accept,
//   - "r":      Reject,
//   - "l":      Later,
//   - "c":      Review,
func New(name string) (Variant, error) {
	val, err := BasicEnumImpl.GetValueByName(
		name)

	if err != nil {
		return newOtherWays(name, err)
	}

	return Variant(val), nil
}
