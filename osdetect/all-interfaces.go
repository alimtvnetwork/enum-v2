package osdetect

import (
	"gitlab.com/auk-go/enum/inttype"
	"gitlab.com/auk-go/enum/strtype"
)

type windowsSysDetailDefiner interface {
	Value(
		name string,
	) strtype.Variant
	ValueInt(
		name string,
	) inttype.Variant
	CloseRegKeyRead()
	CompiledErrorWithStackTraces() error
	Finalize() error
	windowsSystemDetailGetter
}

type windowsSystemDetailGetter interface {
	SystemDetail() (*OperatingSystemDetail, error)
}
