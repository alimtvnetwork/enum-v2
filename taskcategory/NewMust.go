package taskcategory

import "github.com/alimtvnetwork/core-v9/errcore"

func NewMust(taskCategoryName string) Variant {
	newType, err := New(taskCategoryName)
	errcore.HandleErr(err)

	return newType
}
