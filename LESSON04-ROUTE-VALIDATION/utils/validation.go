package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func HandleValidationErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationError {
			switch e.Tag() {
			case "gt":
				errors[e.Field()] = e.Field() + " phải lớn hơn giá trị tối thiểu"
			case "lt":
				errors[e.Field()] = e.Field() + " phải nhỏ hơn giá trị tối thiểu"
			case "gte":
				errors[e.Field()] = e.Field() + " phải lớn hơn hoặc bằng giá trị tối thiểu"
			case "lte":
				errors[e.Field()] = e.Field() + " phải nhỏ hơn hoặc bằng giá trị tối thiểu"
			case "uuid":
				errors[e.Field()] = e.Field() + " phải là UUID hợp lệ"
			case "slug":
				errors[e.Field()] = e.Field() + " chỉ được chứa chữ thường, số, dấu gạch ngang hoặc dấu chấm"
			case "min":
				errors[e.Field()] = fmt.Sprintf("%s phải nhiều hơn %s ký tự", e.Field(), e.Param())
			case "max":
				errors[e.Field()] = fmt.Sprintf("%s phải ít hơn %s ký tự", e.Field(), e.Param())
			case "oneof":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ",")
				errors[e.Field()] = fmt.Sprintf("%s chỉ được chứa các ký tự: %s", e.Field(), allowedValues)
			case "required":
				errors[e.Field()] = e.Field() + " là bắt buộc"
			case "search":
				errors[e.Field()] = e.Field() + " chỉ được chứa chữ thường, in hoa, số, và khoảng trắng"
			case "email":
				errors[e.Field()] = e.Field() + " phải đúng điều kiện email"
			case "datetime":
				errors[e.Field()] = e.Field() + " phải theo đúng định dạng YYYY-MM-DD"
			}
		}
		return gin.H{"error": errors}
	}

	return gin.H{"error": "Yêu cầu không hợp lệ " + err.Error()}
}

func RegisterValidators() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("Failed to get validator engine")
	}

	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	})

	var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
		return searchRegex.MatchString(fl.Field().String())
	})

	return nil
}

// func MinMaxPrice() error {

// }
