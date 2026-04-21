package conntrackstate

import "gitlab.com/auk-go/core/errcore"

func CreateMust(name string) Variant {
	newType, err := Create(name)
	errcore.HandleErr(err)

	return newType
}
