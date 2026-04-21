package sqljointype

// JoinerClause
//
//  Compile using JoinTwoTableSyntaxFormat
//
// References:
//  - Join Sample       : https://www.w3schools.com/sql/sql_join.asp
//  - Inner Join Sample : https://www.w3schools.com/sql/sql_join_inner.asp
//  - Joins Diagrams    : https://prnt.sc/26guaad
//  - Code Samples      : https://prnt.sc/26gub4k
type JoinerClause struct {
	JoinType Variant
	Left     TableWithColumn
	Right    TableWithColumn
}

func (it JoinerClause) SetLeftClone(
	left TableWithColumn,
) JoinerClause {
	return JoinerClause{
		JoinType: it.JoinType,
		Left:     left,
		Right:    it.Right,
	}
}

func (it JoinerClause) SetRightClone(
	right TableWithColumn,
) JoinerClause {
	return JoinerClause{
		JoinType: it.JoinType,
		Left:     it.Left,
		Right:    right,
	}
}

func (it JoinerClause) Clone() JoinerClause {
	return JoinerClause{
		JoinType: it.JoinType,
		Left:     it.Left,
		Right:    it.Right,
	}
}

// String
//
//  generate to string using JoinTwoTableSyntaxFormat
//
// References:
//  - Join Sample       : https://www.w3schools.com/sql/sql_join.asp
//  - Inner Join Sample : https://www.w3schools.com/sql/sql_join_inner.asp
//  - Joins Diagrams    : https://prnt.sc/26guaad
//  - Code Samples      : https://prnt.sc/26gub4k
func (it JoinerClause) String() string {
	return Compile(it)
}

// Compile
//
//  generate to string using JoinTwoTableSyntaxFormat
//
// References:
//  - Join Sample       : https://www.w3schools.com/sql/sql_join.asp
//  - Inner Join Sample : https://www.w3schools.com/sql/sql_join_inner.asp
//  - Joins Diagrams    : https://prnt.sc/26guaad
//  - Code Samples      : https://prnt.sc/26gub4k
func (it JoinerClause) Compile() string {
	return Compile(it)
}
