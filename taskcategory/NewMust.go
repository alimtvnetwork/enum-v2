package taskcategory

import "gitlab.com/auk-go/core/errcore"

func NewMust(taskCategoryName string) Variant {
	newType, err := New(taskCategoryName)
	errcore.HandleErr(err)

	return newType
}
