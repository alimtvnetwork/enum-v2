package creationtests

import (
	"gitlab.com/auk-go/core/coredata/corerange"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
)

type EnumTestWrapper struct {
	Header                     string
	InitialBasicEnumer         enuminf.BasicEnumer
	TypeName                   string
	ExpectedEnumType           enuminf.EnumTyper
	ExpectedMapValues          map[string]interface{} // key - name, value - can be any
	ExpectedInvalidName        string
	ExpectedInvalidValueString string
	ExpectedRangesNamesCsv     string
	IntegerMinMax              corerange.MinMaxInt
	StringMin                  string
	StringMax                  string
}
