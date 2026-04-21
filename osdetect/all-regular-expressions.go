package osdetect

import (
	"gitlab.com/auk-go/core/regexnew"
)

var (
	prettyNameLazyRegex           = regexnew.PrettyNameRegex
	exactIdFieldMatchLazyRegex    = regexnew.ExactIdFieldMatchingRegex
	versionIdLazyRegex            = regexnew.ExactVersionIdFieldMatchingRegex
	ubuntuLazyRegex               = regexnew.UbuntuNameCheckerRegex
	centOSLazyRegex               = regexnew.CentOsNameCheckerRegex
	redHatLazyRegex               = regexnew.RedHatNameCheckerRegex
	windowsVersionNumberLazyRegex = regexnew.WindowsVersionNumberCheckerRegex
)
