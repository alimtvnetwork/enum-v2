package linuxtype

import "github.com/alimtvnetwork/core-v9/errcore"

func NewMust(name string) Variant {
	exitCode, err := New(name)
	errcore.HandleErr(err)

	return exitCode
}
