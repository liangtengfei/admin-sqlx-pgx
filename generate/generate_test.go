package generate

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestGetTableList(t *testing.T) {
	tables := GetTableList()

	require.NotEmpty(t, tables)
	require.Equal(t, 15, len(tables))
}

func TestLoadAllTemplateFiles(t *testing.T) {
	temps, err := LoadAllTemplateFiles()
	require.NoError(t, err)

	require.NotEmpty(t, temps)
	require.Equal(t, len(temps), 4)

	for _, s := range temps {
		t.Log(s)
		require.FileExists(t, s)
	}
}

func TestParseTemplateFiles(t *testing.T) {
	business := "Notice"
	tableName := "ago_notice"

	err := ParseTemplateFiles()
	require.NoError(t, err)

	filenameReq := filepath.Join("../temps", fmt.Sprintf("req.%s.go.txt", business))
	fReq, err := os.OpenFile(filenameReq, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o755)
	defer fReq.Close()
	require.NoError(t, err)

	filenameRes := filepath.Join("../temps", fmt.Sprintf("res.%s.go.txt", business))
	fRes, err := os.OpenFile(filenameRes, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o755)
	defer fRes.Close()
	require.NoError(t, err)

	columns := GetTableColumns(tableName)
	require.Equal(t, true, len(columns) > 0)

	data := make(map[string]interface{}, 0)
	data["Business"] = business
	data["TableName"] = tableName
	data["Columns"] = columns
	//主键默认第一个字段
	data["ColumnPk"] = columns[0].ColumnName

	err = TemplatesMap["api.request.tmpl"].Execute(fReq, data)
	require.NoError(t, err)

	err = TemplatesMap["api.response.tmpl"].Execute(fRes, data)
	require.NoError(t, err)
}
