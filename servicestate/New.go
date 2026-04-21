package servicestate

func New(name string) (Action, error) {
	v, err := BasicEnumImpl.GetValueByName(
		name)

	if err != nil {
		return Invalid, err
	}

	return Action(v), nil
}
