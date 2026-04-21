package linuxservicestate

// NewCodeMapping (using RawMapping)
//
// Reference :
// https://gitlab.com/auk-go/os-manuals/uploads/a3fc906f4ea29a59ebf29490391d0f86/image.png
// https://t.ly/3jkY
func NewCodeMapping(rawExitCode byte) ExitCode {
	if rawExitCode >= rawMappingLength {
		return InvalidCode
	}

	return RawMapping[rawExitCode]
}
