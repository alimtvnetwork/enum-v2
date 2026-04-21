package osgroupexecution

import "gitlab.com/auk-go/core/errcore"

func NewMust(name string) Precedence {
	newType, err := New(name)
	errcore.HandleErr(err)

	return newType
}
