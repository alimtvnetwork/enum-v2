package linuxservicestate

func New(codeOrName string) (ExitCode, error) {
	v, err := BasicEnumImpl.GetValueByName(codeOrName)

	if err != nil {
		return InvalidCode, err
	}

	return ExitCode(v), nil
}
