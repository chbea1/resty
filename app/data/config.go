package data

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/chbea1/resty/app/mysqlspec"
)

type Config struct {
	Table TableConfig
}

type TableConfig struct {
	Host string
	User string
	Db   string
	Pass string
	Port int
	Name string
}

func (c *TableConfig) Valid() error {
	if c.Host == "" || c.User == "" || c.Db == "" ||
		c.Pass == "" || c.Port == 0 || c.Name == "" {
		return errors.New("Unable to create resty configuration.")
	}
	return nil
}

func (config *TableConfig) TableDescription() mysqlspec.MysqlTableSchema {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.Db,
	)
	conn, dbErr := sql.Open("mysql", connection)
	if dbErr != nil {
		panic(dbErr)
	}

	statement, stErr := conn.Prepare(fmt.Sprintf("describe %s", config.Name)) //needs to be changed
	if stErr != nil {
		fmt.Println(stErr)
		os.Exit(1)
	}
	rows, err := statement.Query() // execute select statement
	defer statement.Close()
	if err != nil {
		os.Exit(1)
	}
	var tableSchema mysqlspec.MysqlTableSchema
	var rowSchemas []mysqlspec.RowSchema
	for rows.Next() {
		var field string
		var typeV string
		var null string
		var key string
		var defaultV *string
		var extra string
		err := rows.Scan(&field, &typeV, &null, &key, &defaultV, &extra)
		if err != nil {
			panic(err.Error())
		}
		schema := mysqlspec.RowSchema{
			Field:   field,
			Type:    mysqlspec.GetType(typeV),
			Null:    mysqlspec.GetNull(null),
			Key:     mysqlspec.GetKey(key),
			Default: defaultV,
			Extra:   extra,
		}
		rowSchemas = append(rowSchemas, schema)
	}
	tableSchema.Fields = &rowSchemas
	tableSchema.Name = fmt.Sprintf("%s.%s", config.Db, config.Name)
	return tableSchema
}
