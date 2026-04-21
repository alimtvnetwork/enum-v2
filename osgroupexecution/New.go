package osgroupexecution

func New(name string) (Precedence, error) {
	val, err := BasicEnumImpl.GetValueByName(
		name)

	if err != nil {
		return Invalid, err
	}

	return Precedence(val), nil
}
