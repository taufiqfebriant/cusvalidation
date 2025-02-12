package validator

import (
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate   *validator.Validate
	translator ut.Translator
)

func InitValidator() {
	validate = validator.New(validator.WithRequiredStructEnabled())

	eng := en.New()
	uni := ut.New(eng, eng)
	translator, _ = uni.GetTranslator("en")

	_ = en_translations.RegisterDefaultTranslations(validate, translator)
}

func ValidateStruct(obj interface{}) map[string]string {
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}

	errorsMap := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		jsonKey := getJSONTag(obj, e.StructField())
		errorsMap[jsonKey] = e.Translate(translator)
	}

	return errorsMap
}

func getJSONTag(obj interface{}, fieldName string) string {
	t := reflect.TypeOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	field, found := t.FieldByName(fieldName)
	if !found {
		return fieldName
	}

	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return fieldName
	}

	return jsonTag
}
