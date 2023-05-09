package server

import "github.com/go-playground/validator/v10"

// FailedField 验证失败字段
type FailedField struct {
	FailedField string `json:"failed_field"` // 异常字段
	Tag         string `json:"tag"`          // 验证 tag
	Value       string `json:"value"`        // 验证值
}

var validate = validator.New()

// Validate 验证参数
func Validate[P any](params P) []*FailedField {
	err := validate.Struct(params)
	if err != nil {
		var errors []*FailedField
		for _, err := range err.(validator.ValidationErrors) {
			var element FailedField
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
		return errors
	}

	return nil
}
