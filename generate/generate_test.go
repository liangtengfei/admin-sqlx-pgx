package generate

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

const (
	Business     = "Article"     //业务名称
	BusinessFile = "article"     //业务名称
	BusinessNote = "通知公告"        //业务注解
	BusinessPath = "article"     //业务路径
	TableName    = "cms_article" //业务表名
	TableComment = "文章信息表"       //业务表名
)

//请求和返回信息
func TestParseTemplateFiles_ReqRes(t *testing.T) {
	err := ParseTemplateFiles()
	require.NoError(t, err)

	filenameReq := filepath.Join("../api/request", fmt.Sprintf("req.%s.go", BusinessFile))
	fReq, err := os.OpenFile(filenameReq, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o755)
	defer fReq.Close()
	require.NoError(t, err)

	filenameRes := filepath.Join("../api/response", fmt.Sprintf("res.%s.go", BusinessFile))
	fRes, err := os.OpenFile(filenameRes, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o755)
	defer fRes.Close()
	require.NoError(t, err)

	columns := GetTableColumns(TableName)
	require.Equal(t, true, len(columns) > 0)

	data := make(map[string]interface{}, 0)
	data["Business"] = Business
	data["BusinessNote"] = BusinessNote
	data["BusinessPath"] = BusinessPath
	data["TableName"] = TableName
	data["Columns"] = columns
	//主键默认第一个字段
	data["ColumnPk"] = columns[0].ColumnName

	// controller是否生成List方法
	data["ListAll"] = false

	err = TemplatesMap["api.request.tmpl"].Execute(fReq, data)
	require.NoError(t, err)

	err = TemplatesMap["api.response.tmpl"].Execute(fRes, data)
	require.NoError(t, err)
}

func TestParseTemplateFiles_Schema(t *testing.T) {
	err := ParseTemplateFiles()
	require.NoError(t, err)

	filenameReq := filepath.Join("../db/sqlx", fmt.Sprintf("db.schema.%s.go", BusinessFile))
	fReq, err := os.OpenFile(filenameReq, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o755)
	defer fReq.Close()
	require.NoError(t, err)

	columns := GetTableColumns(TableName)
	require.Equal(t, true, len(columns) > 0)

	data := make(map[string]interface{}, 0)
	data["Business"] = Business
	data["BusinessNote"] = BusinessNote
	data["BusinessPath"] = BusinessPath
	data["TableName"] = TableName
	data["TableComment"] = TableComment
	data["Columns"] = columns
	//主键默认第一个字段
	data["ColumnPk"] = columns[0].ColumnName

	err = TemplatesMap["db.schema.tmpl"].Execute(fReq, data)
	require.NoError(t, err)
}

// controller、service、mapper等
func TestParseTemplateFiles_CtrlServiceMapper(t *testing.T) {
	err := ParseTemplateFiles()
	require.NoError(t, err)

	filenameCtrl := filepath.Join("../api", fmt.Sprintf("ctrl.%s.go", BusinessFile))
	fReq, err := os.OpenFile(filenameCtrl, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o755)
	defer fReq.Close()
	require.NoError(t, err)

	filenameService := filepath.Join("../service", fmt.Sprintf("service.%s.go", BusinessFile))
	fService, err := os.OpenFile(filenameService, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o755)
	defer fService.Close()
	require.NoError(t, err)

	filenameMapper := filepath.Join("../db/sqlx", fmt.Sprintf("mapper.%s.go", BusinessFile))
	fMapper, err := os.OpenFile(filenameMapper, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o755)
	defer fMapper.Close()
	require.NoError(t, err)

	filenameInter := filepath.Join("../db/sqlx", fmt.Sprintf("inter.%s.go", BusinessFile))
	fInter, err := os.OpenFile(filenameInter, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o755)
	defer fInter.Close()
	require.NoError(t, err)

	filenameRouter := filepath.Join("../router/business", fmt.Sprintf("router.%s.go", BusinessFile))
	fRouter, err := os.OpenFile(filenameRouter, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o755)
	defer fRouter.Close()
	require.NoError(t, err)

	columns := GetTableColumns(TableName)
	require.Equal(t, true, len(columns) > 0)

	//
	data := make(map[string]interface{}, 0)
	data["Business"] = Business
	data["BusinessNote"] = BusinessNote
	data["BusinessPath"] = BusinessPath
	data["TableName"] = TableName
	data["Columns"] = columns
	//主键默认第一个字段
	//data["ColumnPk"] = columns[0].ColumnName

	// controller是否生成List方法
	data["ListAll"] = true

	err = TemplatesMap["api.ctrl.tmpl"].Execute(fReq, data)
	require.NoError(t, err)

	err = TemplatesMap["db.service.tmpl"].Execute(fService, data)
	require.NoError(t, err)

	err = TemplatesMap["db.mapper.tmpl"].Execute(fMapper, data)
	require.NoError(t, err)

	err = TemplatesMap["db.inter.tmpl"].Execute(fInter, data)
	require.NoError(t, err)

	err = TemplatesMap["router.biz.tmpl"].Execute(fRouter, data)
	require.NoError(t, err)
}

// 2022年05月30日 在inter.business.go和business.go 中加入生成的 mapper接口和路由信息
func TestParseTemplateFiles_InterRouter(t *testing.T) {
	interFilePath := filepath.Join("../db/sqlx/internal", "entry.go")
	interFile, err := os.Open(interFilePath)
	defer interFile.Close()
	require.NoError(t, err)

	querierStr := fmt.Sprintf("\tQuerier%s2", Business)
	err = insertStringToFileEnd(interFilePath, querierStr, 2)
	require.NoError(t, err)

	routerFilePath := filepath.Join("../router", "entry.go")
	routerFile, err := os.Open(routerFilePath)
	defer routerFile.Close()
	require.NoError(t, err)

}

func file2Lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func insertStringToFile(path, content string, index int) error {
	lines, err := file2Lines(path)
	if err != nil {
		return err
	}

	var builder strings.Builder
	for i, line := range lines {
		if i == index {
			builder.WriteString(content)
		}
		builder.WriteString(line)
		builder.WriteString("\n")
	}

	return ioutil.WriteFile(path, []byte(builder.String()), 0666)
}

// 插入末尾第几行 offset 跳过几行
func insertStringToFileEnd(path, content string, offset int) error {
	lines, err := file2Lines(path)
	if err != nil {
		return err
	}

	index := len(lines) - offset

	var builder strings.Builder
	for i, line := range lines {
		if i == index {
			builder.WriteString(content)
			builder.WriteString("\n")
		}
		builder.WriteString(line)
		builder.WriteString("\n")
	}

	return ioutil.WriteFile(path, []byte(builder.String()), 0666)
}
