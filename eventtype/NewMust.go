package eventtype

import "github.com/alimtvnetwork/core-v9/errcore"

func NewMust(name string) Variant {
	newType, err := New(name)
	errcore.HandleErr(err)

	return newType
}
