package validator

//gin > 1.4.0

//将验证器错误翻译成中文

import (
	"easy-gin/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var ValidTrans ut.Translator
var ValidObj *validator.Validate

func init() {
	ValidObj = validator.New()
	english := en.New()
	chinese := zh.New()
	uni := ut.New(chinese, english)
	ValidTrans, _ = uni.GetTranslator("zh")
	_ = zhTranslations.RegisterDefaultTranslations(ValidObj, ValidTrans)
}

//CheckParams 入参验证
func CheckParams(ctx *gin.Context, err error) {
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range err.(validator.ValidationErrors) {
				msg, _ := ValidTrans.T(fieldError.Tag(), fieldError.Field(), fieldError.Param())
				response.Error(ctx, msg)
				return
			}
		} else {
			response.TypeError(ctx, err.Error())
			return
		}
	}
}
