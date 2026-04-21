package sqljointype

const (
	// JoinTwoTableSyntaxFormat
	//
	// References:
	//  - JOINSyntax        : refers to the enum sql-syntax
	//  - Join Sample       : https://www.w3schools.com/sql/sql_join.asp
	//  - Inner Join Sample : https://www.w3schools.com/sql/sql_join_inner.asp
	//  - Joins Diagrams    : https://prnt.sc/26guaad
	//  - Code Samples      : https://prnt.sc/26gub4k
	JoinTwoTableSyntaxFormat = "%s %s ON %s = %s" // JOINSyntax, LeftTable, LeftTableWithField, RightTableWithField
	tableWithFieldNameFormat = "%s.%s"            // TableName, FieldName
)
