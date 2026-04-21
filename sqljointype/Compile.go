package sqljointype

import "fmt"

func Compile(clause JoinerClause) string {
	return fmt.Sprintf(
		JoinTwoTableSyntaxFormat,
		clause.JoinType.ToSqlName(),
		clause.Left.TableName,
		clause.Left.TableWithField(),
		clause.Right.TableWithField(),
	)
}
