package utils

import (
	"go.uber.org/zap"
	"log"
	"study.com/demo-sqlx-pgx/global"
)
import "github.com/jinzhu/copier"

func StructCopyMust(toVal interface{}, fromVal interface{}) {
	err := copier.Copy(toVal, fromVal)
	if err != nil {
		log.Println("拷贝结构体信息错误：", err.Error())
	}
}

func StructCopy(toVal interface{}, fromVal interface{}) error {
	err := copier.Copy(toVal, fromVal)
	if err != nil {
		global.Log.Error("结构体复制错误", zap.String("Msg", err.Error()))
		return err
	}
	return nil
}
