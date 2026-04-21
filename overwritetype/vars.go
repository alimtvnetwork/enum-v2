package overwritetype

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coreimpl/enumimpl"
)

var (
	Ranges = [...]string{
		Invalid:                            "Invalid",
		ForceWrite:                         "ForceWrite",
		SkipOnExistFiles:                   "SkipOnExistFiles",
		IgnoreRepeatInFolderNameExtraction: "IgnoreRepeatInFolderNameExtraction",
		Yes:                                "Yes",
		No:                                 "No",
	}

	RangesMap = map[string]Variant{
		"Invalid":                            Invalid,
		"ForceWrite":                         ForceWrite,
		"SkipOnExistFiles":                   SkipOnExistFiles,
		"IgnoreRepeatInFolderNameExtraction": IgnoreRepeatInFolderNameExtraction,
		"Yes":                                Yes,
		"No":                                 No,
	}

	overwriteMap = map[Variant]bool{
		ForceWrite:       true,
		ForceWriteRepeat: true,
		Yes:              true,
	}

	BasicEnumImpl = enumimpl.New.BasicByte.UsingTypeSlice(
		coredynamic.TypeName(Invalid),
		Ranges[:])
)
