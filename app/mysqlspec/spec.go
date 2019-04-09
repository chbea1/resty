package mysqlspec

import "encoding/json"

type MysqlDatabaseSchema struct {
	Tables *[]MysqlTableSchema
}

type MysqlTableSchema struct {
	Name   string
	Fields *[]RowSchema
}

type RowSchema struct {
	Field   string
	Type    Type
	Null    NULL
	Key     KEY
	Default *string
	Extra   string
}

func (m MysqlDatabaseSchema) String() string {
	str, _ := json.MarshalIndent(m, "", "   ")
	return string(str)
}

func (m MysqlTableSchema) String() string {
	str, _ := json.MarshalIndent(m, "", "   ")
	return string(str)
}
