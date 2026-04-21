package sqljointype

import "fmt"

type TableWithColumn struct {
	TableName, ColumnName string
}

// TableWithField
//
// Table.FieldName
func (it TableWithColumn) TableWithField() string {
	return fmt.Sprintf(
		tableWithFieldNameFormat,
		it.TableName,
		it.ColumnName)
}

func (it TableWithColumn) Clone() TableWithColumn {
	return TableWithColumn{
		TableName:  it.TableName,
		ColumnName: it.ColumnName,
	}
}
