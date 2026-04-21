package linuxservicestate

import "gitlab.com/auk-go/core/errcore"

func NewMust(codeOrName string) ExitCode {
	exitCode, err := New(codeOrName)
	errcore.HandleErr(err)

	return exitCode
}
