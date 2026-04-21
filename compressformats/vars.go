package compressformats

import (
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	ranges = [...]string{
		Zip:     "Zip",
		Tar:     "Tar",
		TarGZ:   "TarGZ",
		TarXZ:   "TarXZ",
		TarBz2:  "TarBz2",
		Invalid: "Invalid",
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingFirstItemSliceAllCases(
		Zip,
		ranges[:])
)
