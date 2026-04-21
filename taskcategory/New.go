package taskcategory

func New(taskCategoryName string) (Variant, error) {
	val, err := BasicEnumImpl.GetValueByName(
		taskCategoryName)

	if err != nil {
		return Invalid, err
	}

	return Variant(val), nil
}
