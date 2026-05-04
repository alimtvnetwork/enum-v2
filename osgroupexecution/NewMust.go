package osgroupexecution

import "github.com/alimtvnetwork/core-v9/errcore"

func NewMust(name string) Precedence {
	newType, err := New(name)
	errcore.HandleErr(err)

	return newType
}
