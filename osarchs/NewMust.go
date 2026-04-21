package osarchs

import "gitlab.com/auk-go/core/errcore"

// NewMust
//
//	Creates string to the type Variant
func NewMust(name string) Architecture {
	newType, err := New(name)
	errcore.HandleErr(err)

	return newType
}
