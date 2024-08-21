package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validator(values interface{}) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		if err := ctx.ShouldBindJSON(&values); err != nil {
			ctx.JSON(400, gin.H{
				"message": "Bad Request",
				"error":   err,
			})
			ctx.Abort()
			return
		}

		validate := validator.New()
		err := validate.Struct(values)
		if err != nil {
			println(err.Error(), "error")
			if _, ok := err.(*validator.InvalidValidationError); ok {

				fmt.Println(err.Error(), "error1")
				return
			}
			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.StructNamespace())
				fmt.Println(err.StructField())
				fmt.Println(err.Tag())
				fmt.Println(err.ActualTag())
				fmt.Println(err.Kind())
				fmt.Println(err.Type())
				fmt.Println(err.Value())
				fmt.Println(err.Param())
				fmt.Println()
			}
		}

		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "Bad Request",
				"error":   err,
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
