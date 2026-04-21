package onofftype

import (
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/issetter"
)

func NewUsingSetter(value issetter.Value) (Variant, error) {
	mappedItem, has := isSetterWithVariantMap[value]

	if !has {
		return Invalid, errcore.
			KeyNotExistInMapType.
			Error(
				typeConvFailedPrefixMsg+coredynamic.TypeName(value),
				mapReferenceMessage)
	}

	return Variant(mappedItem.Value()), nil
}
