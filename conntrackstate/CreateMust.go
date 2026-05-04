package conntrackstate

import "github.com/alimtvnetwork/core-v9/errcore"

func CreateMust(name string) Variant {
	newType, err := Create(name)
	errcore.HandleErr(err)

	return newType
}
