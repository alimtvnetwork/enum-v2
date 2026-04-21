package strtype

import (
	"sync"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/ostype"
	"gitlab.com/auk-go/enum/osarchs"
)

var (
	invalidMaps = map[string]bool{
		"":              true,
		"Invalid":       true,
		"invalid":       true,
		"Unknown":       true,
		"unknown":       true,
		"Undefined":     true,
		"undefined":     true,
		"Uninitialized": true,
		"uninitialized": true,
		"None":          true,
		"none":          true,
		"Unspecified":   true,
		"unspecified":   true,
	}

	globalMutex = sync.Mutex{}

	// Arch
	// Current OS architecture
	Arch  = osarchs.CurrentArch
	Group = ostype.CurrentGroupVariant.Group
	// Type Current Os Type
	Type               = ostype.CurrentGroupVariant
	bytesSerializer    = corejson.Serialize.ToBytesErr
	stringDeserializer = corejson.Deserialize.BytesTo.String

	typeName = coredynamic.TypeName(Variant(""))
)
