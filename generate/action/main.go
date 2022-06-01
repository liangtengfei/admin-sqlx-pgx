package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"study.com/demo-sqlx-pgx/generate"
	"study.com/demo-sqlx-pgx/utils/strcase"
)

//const (
//	Business        = "Article"     //业务名称
//	BusinessComment = "通知公告"        //业务注解
//	TableName       = "cms_article" //业务表名
//	TableComment    = "文章信息表"       //业务表名注释
//)

type GenBusiness struct {
	Business        string
	BusinessComment string
	TableName       string
	TableComment    string
}

func (g GenBusiness) GetBusinessPath() string {
	return strcase.ToLowerCamel(g.Business)
}

func (g GenBusiness) GetBusinessSchema() string {
	return strcase.ToCamel(g.TableName)
}

func main() {
	business := GenBusiness{
		Business:        "File",
		BusinessComment: "文件上传",
		TableName:       "ago_file",
		TableComment:    "上传文件信息表",
	}
	var err error
	err = generate.ParseTemplateFiles()
	if err != nil {
		log.Fatal("加载模板文件错误：", err)
	}

	reqFileFormat := fmt.Sprintf("req.%s.go", business.GetBusinessPath())
	reqFilePath := filepath.Join("api/request", reqFileFormat)
	reqTemplateFile := "api.request.tmpl"
	err = fileCreate(reqFilePath, reqTemplateFile, business)
	failAndRemoveFile(reqFilePath, err)

	resFileFormat := fmt.Sprintf("res.%s.go", business.GetBusinessPath())
	resFilePath := filepath.Join("api/response", resFileFormat)
	resTemplateFile := "api.response.tmpl"
	err = fileCreate(resFilePath, resTemplateFile, business)
	failAndRemoveFile(resFilePath, err)

	ctrlFileFormat := fmt.Sprintf("ctrl.%s.go", business.GetBusinessPath())
	ctrlFilePath := filepath.Join("api", ctrlFileFormat)
	ctrlTemplateFile := "api.ctrl.tmpl"
	err = fileCreate(ctrlFilePath, ctrlTemplateFile, business)
	failAndRemoveFile(ctrlFilePath, err)

	serviceFileFormat := fmt.Sprintf("service.%s.go", business.GetBusinessPath())
	serviceFilePath := filepath.Join("service", serviceFileFormat)
	serviceTemplateFile := "db.service.tmpl"
	err = fileCreate(serviceFilePath, serviceTemplateFile, business)
	failAndRemoveFile(serviceFilePath, err)

	mapperFileFormat := fmt.Sprintf("mapper.%s.go", business.GetBusinessPath())
	mapperFilePath := filepath.Join("db/sqlx", mapperFileFormat)
	mapperTemplateFile := "db.mapper.tmpl"
	err = fileCreate(mapperFilePath, mapperTemplateFile, business)
	failAndRemoveFile(mapperFilePath, err)

	querierFileFormat := fmt.Sprintf("querier.%s.go", business.GetBusinessPath())
	querierFilePath := filepath.Join("db/sqlx", querierFileFormat)
	querierTemplateFile := "db.inter.tmpl"
	err = fileCreate(querierFilePath, querierTemplateFile, business)
	failAndRemoveFile(querierFilePath, err)

	schemaFileFormat := fmt.Sprintf("schema.%s.go", business.GetBusinessPath())
	schemaFilePath := filepath.Join("db/sqlx/internal", schemaFileFormat)
	schemaTemplateFile := "db.schema.tmpl"
	err = fileCreate(schemaFilePath, schemaTemplateFile, business)
	failAndRemoveFile(schemaFilePath, err)

	routerFileFormat := fmt.Sprintf("router.%s.go", business.GetBusinessPath())
	routerFilePath := filepath.Join("router/business", routerFileFormat)
	routerTemplateFile := "router.biz.tmpl"
	err = fileCreate(routerFilePath, routerTemplateFile, business)
	failAndRemoveFile(routerFilePath, err)

	// 需要插入生成的路由和接口
	interEntryFilePath := filepath.Join("db/sqlx", "entry.go")
	interEntryContent := fmt.Sprintf("\tQuerier%s", business.Business)
	err = fileInsert(interEntryFilePath, interEntryContent)
	if err != nil {
		log.Fatal(err)
	}

	routerEntryFilePath := filepath.Join("router", "entry.go")
	routerEntryContent := fmt.Sprintf("\tbusiness.%sRouter(root)", business.Business)
	err = fileInsert(routerEntryFilePath, routerEntryContent)
	if err != nil {
		log.Fatal(err)
	}
}

func fileCreate(filePath, templatePath string, business GenBusiness) error {
	f, err := os.OpenFile(filePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0o755)
	if err != nil {
		return err
	}
	defer f.Close()

	columns := generate.GetTableColumns(business.TableName)
	if len(columns) <= 0 {
		return errors.New("读取数据表列错误")
	}

	data := make(map[string]interface{}, 0)
	data["Biz"] = business
	data["Columns"] = columns
	data["ColumnPk"] = columns[0].ColumnName

	err = generate.TemplatesMap[templatePath].Execute(f, data)
	if err != nil {
		fileRemove(filePath)
	}
	return err
}

func fileInsert(filePath, content string) error {
	return generate.InsertStringToFileEnd(filePath, content, 1)
}

func failAndRemoveFile(filePath string, err error) {
	if err != nil {
		log.Fatal("生成文件错误：", err)
	}
}

// 清除生成的文件
func fileRemove(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Println("删除文件错误：", err)
	}
}
