package generate

import (
	"context"
	"database/sql"
	"log"
	"strings"
)

type TableColumn struct {
	DataType        string         `json:"dataType"`
	ColumnName      string         `json:"columnName"`
	DataTypeLong    sql.NullString `json:"dataTypeLong"`
	ColumnComment   sql.NullString `json:"columnComment"`
	IsNullable      string         `json:"isNullable"`
	OrdinalPosition int64          `json:"ordinalPosition"`
}

func GetTableColumns(tableName string) []TableColumn {
	ctx := context.Background()
	db := GetDBConn()
	defer db.Close()

	sql := `SELECT columns.COLUMN_NAME                                                                                      as column_name,
columns.is_nullable as is_nullable,
columns.ordinal_position,
		   columns.udt_name                                                                                        as data_type,
		   CASE
			   columns.udt_name
			   WHEN 'int8' THEN
				   concat_ws('', '', columns.CHARACTER_MAXIMUM_LENGTH)
			   WHEN 'varchar' THEN
				   concat_ws('', '', columns.CHARACTER_MAXIMUM_LENGTH)
			   WHEN 'smallint' THEN
				   concat_ws(',', columns.NUMERIC_PRECISION, columns.NUMERIC_SCALE)
			   WHEN 'decimal' THEN
				   concat_ws(',', columns.NUMERIC_PRECISION, columns.NUMERIC_SCALE)
			   WHEN 'integer' THEN
				   concat_ws('', '', columns.NUMERIC_PRECISION)
			   WHEN 'bigint' THEN
				   concat_ws('', '', columns.NUMERIC_PRECISION)
			   ELSE ''
			   END                                                                                                  AS data_type_long,
		   (select description.description
			from pg_description description
			where description.objoid = (select attribute.attrelid
										from pg_attribute attribute
										where attribute.attrelid =
											  (select oid from pg_class class where class.relname = '@TableName') and attname =columns.COLUMN_NAME )
			  and description.objsubid = (select attribute.attnum
										  from pg_attribute attribute
										  where attribute.attrelid =
												(select oid from pg_class class where class.relname = '@TableName') and attname =columns.COLUMN_NAME )) as column_comment
		FROM INFORMATION_SCHEMA.COLUMNS columns
		WHERE table_catalog = '@TableCatalog'
		  and table_schema = '@TableSchema'
		  and table_name = '@TableName'`

	sql = strings.ReplaceAll(sql, "@TableCatalog", TableCatalog)
	sql = strings.ReplaceAll(sql, "@TableSchema", TableSchema)
	sql = strings.ReplaceAll(sql, "@TableName", tableName)

	rows, err := db.Query(ctx, sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var items []TableColumn
	for rows.Next() {
		var i TableColumn
		if err := rows.Scan(&i.ColumnName, &i.IsNullable, &i.OrdinalPosition, &i.DataType, &i.DataTypeLong, &i.ColumnComment); err != nil {
			log.Fatal(err)
		}
		items = append(items, i)
	}

	return items
}
