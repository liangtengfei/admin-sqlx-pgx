package generate

import (
	"context"
	"log"
	"strings"
)

type TableInfo struct {
	TableCatalog string `json:"tableCatalog"`
	TableSchema  string `json:"tableSchema"`
	TableName    string `json:"tableName"`
	TableComment string `json:"tableComment"`
}

func GetTableList() []TableInfo {
	ctx := context.Background()

	db := GetDBConn()
	defer db.Close()

	sql := `SELECT
	it.table_catalog,
	it.table_schema,
	it. "table_name",
	(SELECT COALESCE(obj_description((it.table_schema || '.' || it. "table_name")::regclass), '')) AS table_comment
FROM
	information_schema.tables it
WHERE
	it.table_catalog = '@TableCatalog'
	AND it.table_schema = '@TableSchema'`

	sql = strings.ReplaceAll(sql, "@TableCatalog", TableCatalog)
	sql = strings.ReplaceAll(sql, "@TableSchema", TableSchema)
	rows, err := db.Query(ctx, sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var items []TableInfo
	for rows.Next() {
		var i TableInfo
		if err := rows.Scan(&i.TableCatalog, &i.TableSchema, &i.TableName, &i.TableComment); err != nil {
			log.Fatal(err)
		}
		items = append(items, i)
	}

	return items
}
