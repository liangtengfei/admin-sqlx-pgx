package generate

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"study.com/demo-sqlx-pgx/utils/strcase"
	"text/template"
)

var TemplatesMap map[string]*template.Template

func LoadAllTemplateFiles() (temps map[string]string, err error) {
	//generate/templates
	files, err := ioutil.ReadDir("../generate/templates")
	if err != nil {
		return
	}
	temps = make(map[string]string, 0)
	for _, fi := range files {
		if strings.HasSuffix(fi.Name(), ".tmpl") {
			temps[fi.Name()] = filepath.Join("../generate/templates", fi.Name())
		}
	}
	return
}

func ParseTemplateFiles() error {
	TemplatesMap = make(map[string]*template.Template, 0)
	temps, err := LoadAllTemplateFiles()
	if err != nil {
		return err
	}
	for k, v := range temps {
		tmpl, err := template.New(k).Funcs(template.FuncMap{
			"inc":          incFunc,
			"dec":          decFunc,
			"comma":        commaFunc,
			"contains":     containsFunc,
			"containsNot":  notContainsFunc,
			"toLowerCamel": toLowerCamel,
			"toUpperCamel": toUpperCamel,
		}).ParseFiles(v)
		if err != nil {
			return err
		}
		TemplatesMap[k] = tmpl
	}
	return nil
}

func GetTemplate(name string) (*template.Template, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(wd, "generate/templates", "*.tmpl")

	//设置模板
	return template.New(name).Funcs(template.FuncMap{
		"inc": incFunc,
		"dec": decFunc,
		//处理SQL最后一个逗号问题
		"comma":        commaFunc,
		"contains":     containsFunc,
		"containsNot":  notContainsFunc,
		"toLowerCamel": toLowerCamel,
		"toUpperCamel": toUpperCamel,
	}).ParseGlob(path)
}

func incFunc(i int) int {
	return i + 1
}

func decFunc(i int) int {
	return i - 1
}

func commaFunc(length int) string {
	var s []string
	for i := 0; i < length; i++ {
		s = append(s, fmt.Sprintf("$%d", i+1))
	}
	return strings.Join(s, ", ")
}

func containsFunc(val, prefix string) bool {
	return strings.Contains(val, prefix)
}

func notContainsFunc(val, prefix string) bool {
	return !containsFunc(val, prefix)
}

func toLowerCamel(val string) string {
	return strcase.ToLowerCamel(val)
}

func toUpperCamel(val string) string {
	return strcase.ToCamel(val)
}
