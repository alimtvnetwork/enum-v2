package sqliteconnpathtype

type sqliteConnectionCompiler struct{}

func (it sqliteConnectionCompiler) AllFormat() string {
	return sqliteConnectionFormats[AllSqlitePath]
}

func (it sqliteConnectionCompiler) AllWithTypeFormat() string {
	return sqliteConnectionFormats[AllSqlitePath]
}

func (it sqliteConnectionCompiler) AllWithDynamicTypeFormat() string {
	return sqliteConnectionFormats[AllWithTypeAndDynamicSqlitePath]
}

func (it sqliteConnectionCompiler) AllWithSequenceTypeFormat() string {
	return sqliteConnectionFormats[AllWithTypeAndSequenceSqlitePath]
}

func (it sqliteConnectionCompiler) PrefixFormat() string {
	return sqliteConnectionFormats[PrefixSqlitePath]
}

func (it sqliteConnectionCompiler) PrefixTypeFormat() string {
	return sqliteConnectionFormats[PrefixTypeSqlitePath]
}

func (it sqliteConnectionCompiler) DynamicSpecificFormat() string {
	return sqliteConnectionFormats[DynamicSpecificSqlitePath]
}

func (it sqliteConnectionCompiler) SequenceSpecificFormat() string {
	return sqliteConnectionFormats[SequenceSpecificSqlitePath]
}

func (it sqliteConnectionCompiler) DynamicSequenceSpecificFormat() string {
	return sqliteConnectionFormats[DynamicSequenceSpecificSqlitePath]
}

func (it sqliteConnectionCompiler) SpecificFormat() string {
	return sqliteConnectionFormats[SpecificSqlitePath]
}

func (it sqliteConnectionCompiler) CompileUsingType(
	pathType Variant,
	conn SqliteConnectionOption,
) string {
	return conn.Compile(pathType.PathFormat())
}

func (it sqliteConnectionCompiler) CompileUsingAllWithParams(
	root, prefix, db string,
) string {
	conn := SqliteConnectionOption{
		Root:   root,
		Prefix: prefix,
		DbName: db,
	}

	return it.CompileUsingType(AllSqlitePath, conn)
}

func (it sqliteConnectionCompiler) CompileUsingAllTypeWithParams(
	root, prefix, typeName, db string,
) string {
	conn := SqliteConnectionOption{
		Root:     root,
		Prefix:   prefix,
		DbName:   db,
		TypeName: typeName,
	}

	return it.CompileUsingType(
		AllWithTypeSqlitePath,
		conn)
}

func (it sqliteConnectionCompiler) CompileUsingAllTypeDynamicWithParams(
	root, prefix, typeName, dynamic, db string,
) string {
	conn := SqliteConnectionOption{
		Root:     root,
		Prefix:   prefix,
		DbName:   db,
		TypeName: typeName,
		Dynamic:  dynamic,
	}

	return it.CompileUsingType(
		AllWithTypeAndDynamicSqlitePath,
		conn)
}

func (it sqliteConnectionCompiler) CompileUsingAllTypeSequenceWithParams(
	root, prefix, typeName, sequence, db string,
) string {
	conn := SqliteConnectionOption{
		Root:     root,
		Prefix:   prefix,
		DbName:   db,
		TypeName: typeName,
		Sequence: sequence,
	}

	return it.CompileUsingType(
		AllWithTypeAndSequenceSqlitePath,
		conn)
}

func (it sqliteConnectionCompiler) CompileUsingPrefixWithParams(
	prefix, typeName, db string,
) string {
	conn := SqliteConnectionOption{
		Prefix:   prefix,
		DbName:   db,
		TypeName: typeName,
	}

	return it.CompileUsingType(
		PrefixSqlitePath,
		conn)
}

func (it sqliteConnectionCompiler) CompileUsingPrefixTypeWithParams(
	prefix, typeName, db string,
) string {
	conn := SqliteConnectionOption{
		Prefix:   prefix,
		DbName:   db,
		TypeName: typeName,
	}

	return it.CompileUsingType(
		PrefixTypeSqlitePath,
		conn)
}

func (it sqliteConnectionCompiler) CompileUsingDynamicSpecificWithParams(
	dynamic, db string,
) string {
	conn := SqliteConnectionOption{
		DbName:  db,
		Dynamic: dynamic,
	}

	return it.CompileUsingType(
		DynamicSpecificSqlitePath,
		conn)
}

func (it sqliteConnectionCompiler) CompileUsingSpecific(
	db string,
) string {
	conn := SqliteConnectionOption{
		DbName: db,
	}

	return it.CompileUsingType(
		SpecificSqlitePath,
		conn)
}

func (it sqliteConnectionCompiler) CompileUsingDb(
	db string,
) string {
	conn := SqliteConnectionOption{
		DbName: db,
	}

	return it.CompileUsingType(
		SpecificSqlitePath,
		conn)
}

func (it sqliteConnectionCompiler) CompileUsingSequenceSpecificWithParams(
	sequence, db string,
) string {
	conn := SqliteConnectionOption{
		DbName:   db,
		Sequence: sequence,
	}

	return it.CompileUsingType(
		AllWithTypeSqlitePath,
		conn)
}

func (it sqliteConnectionCompiler) CompileUsingDynamicSequenceSpecificWithParams(
	dynamic, sequence, db string,
) string {
	conn := SqliteConnectionOption{
		DbName:   db,
		Dynamic:  dynamic,
		Sequence: sequence,
	}

	return it.CompileUsingType(
		AllWithTypeSqlitePath,
		conn)
}

// CompileUsingFormat
//
// Format each meaning:
//  - "{root}"     :  Root,
//	- "{prefix}"   :  Prefix,
//	- "{db-name}"  :  DbName,
//	- "{type}"     :  TypeName,
//	- "{dynamic}"  :  Dynamic,
//	- "{sequence}" :  Sequence,
func (it sqliteConnectionCompiler) CompileUsingFormat(
	format string,
	conn SqliteConnectionOption,
) string {
	return conn.Compile(format)
}

func (it sqliteConnectionCompiler) PathFormat(
	pathType Variant,
) string {
	return sqliteConnectionFormats[pathType]
}
