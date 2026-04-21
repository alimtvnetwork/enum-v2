package inttype

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coredata/corejson"
)

var (
	typeName = coredynamic.TypeName(Variant(-1))

	bytesToDeserializer = corejson.
				Deserialize.
				BytesTo
)
