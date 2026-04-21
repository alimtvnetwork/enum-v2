package osarchs

// New
//
//  Creates string to the type Architecture
func New(name string) (Architecture, error) {
	val, err := BasicEnumImpl.GetValueByName(
		name)

	if err != nil {
		return Invalid, err
	}

	return Architecture(val), nil
}
