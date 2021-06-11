package utils

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"strings"
)

var trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans() (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
	return
}

func ErrHandler(msg string, err error) string{
	if err := InitTrans(); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		log.Print(1)
		return fmt.Sprintf("%s,%s",msg,err.Error())
	}
	errs, ok := err.(validator.ValidationErrors)
	if !ok{
		return fmt.Sprintf("%s,%s",msg,err.Error())
	}
	res := removeTopStruct(errs.Translate(trans))
	log.Println(res)
	return fmt.Sprintf("%s,%s",msg,res)
}
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}