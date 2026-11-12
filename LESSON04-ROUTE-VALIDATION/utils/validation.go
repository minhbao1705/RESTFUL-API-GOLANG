package utils

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func HandleValidationErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		log.Printf("%s", errors)

		for _, e := range validationError {
			root := strings.Split(e.Namespace(), ".")[0]

			rawPath := strings.TrimPrefix(e.Namespace(), root + ".")

			parts := strings.Split(rawPath, ".")

			for i, part := range parts {
				if strings.Contains("part", "[") {
					idx := strings.Index(part, "[")
					base := camelToSnake(part[:idx]) // từ 0 đến trước [
					index := part[idx:]
					parts[i] = base + index
				} else {
					parts[i] = camelToSnake(part)
				}
			}

			fieldPath := strings.Join(parts, ".")

			switch e.Tag() {
			case "gt":
				errors[fieldPath] = fmt.Sprintf("%s phải lớn hơn %s", fieldPath, e.Param())
			case "lt":
				errors[fieldPath] = fmt.Sprintf("%s phải nhỏ hơn %s", fieldPath, e.Param())
			case "gte":
				errors[fieldPath] = fmt.Sprintf("%s phải lớn hơn hoặc bằng %s", fieldPath, e.Param())
			case "lte":
				errors[fieldPath] = fmt.Sprintf("%s phải nhỏ hơn hoặc bằng %s", fieldPath, e.Param())
			case "uuid":
				errors[fieldPath] = fmt.Sprintf("%s phải là UUID hợp lệ", fieldPath)
			case "slug":
				errors[fieldPath] = fmt.Sprintf("%s chỉ được chứa chữ thường, số, dấu gạch ngang hoặc dấu chấm", fieldPath)
			case "min":
				errors[fieldPath] = fmt.Sprintf("%s phải nhiều hơn %s ký tự", fieldPath, e.Param())
			case "max":
				errors[fieldPath] = fmt.Sprintf("%s phải ít hơn %s ký tự", fieldPath, e.Param())
			case "min_int":
				errors[fieldPath] = fmt.Sprintf("%s phải có giá trị lớn hơn %s", fieldPath, e.Param())
			case "max_int":
				errors[fieldPath] = fmt.Sprintf("%s phải có giá trị bé hơn %s", fieldPath, e.Param())
			case "oneof":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ",")
				errors[fieldPath] = fmt.Sprintf("%s chỉ được chứa các ký tự: %s", fieldPath, allowedValues)
			case "required":
				errors[fieldPath] = fmt.Sprintf("%s là bắt buộc", fieldPath)
			case "search":
				errors[fieldPath] = fmt.Sprintf("%s chỉ được chứa chữ thường, in hoa, số, và khoảng trắng", fieldPath)
			case "email":
				errors[fieldPath] = fmt.Sprintf("%s phải đúng điều kiện email", fieldPath)
			case "datetime":
				errors[fieldPath] = fmt.Sprintf("%s phải theo đúng định dạng YYYY-MM-DD", fieldPath)
			case "file_ext":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ",")
				errors[fieldPath] = fmt.Sprintf("%s chỉ cho phép những file có extension: %s", fieldPath, allowedValues)
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

	v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
		minStr := fl.Param()
		// Base
		// 10: Hệ thập phân (decimal)
		// 16: Hệ thập lục phân (hex, ví dụ "FF" => 255)
		// 2: Hệ nhị phân (binary, ví dụ "1010" => 10)
		minVal, err := strconv.ParseInt(minStr, 10, 64)
		if err != nil {
			return false
		}

		return fl.Field().Int() >= minVal
	})

	v.RegisterValidation("max_int", func(fl validator.FieldLevel) bool {
		maxStr := fl.Param()
		// Base
		// 10: Hệ thập phân (decimal)
		// 16: Hệ thập lục phân (hex, ví dụ "FF" => 255)
		// 2: Hệ nhị phân (binary, ví dụ "1010" => 10)
		maxVal, err := strconv.ParseInt(maxStr, 10, 64)
		if err != nil {
			return false
		}

		return fl.Field().Int() >= maxVal
	})

	v.RegisterValidation("file_ext", func(fl validator.FieldLevel) bool {
		filename := fl.Field().String()

		allowedStr := fl.Param()
		if allowedStr == "" {
			return false
		}

		allowedExt := strings.Fields(allowedStr)
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filename)), ".")

		for _, allowed := range allowedExt {
			if ext == strings.ToLower(allowed) {
				return true
			}
		}
		return false
	})

	return nil
}


