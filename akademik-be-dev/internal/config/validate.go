package config

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"unsia.ac.id/akademic_be/pkg/utils"
)

func NewValidateConfig() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("valid_number_phone", phoneNumberValidator)
	validate.RegisterValidation("nik", nikValidator)
	validate.RegisterValidation("parent", parentValidator)
	validate.RegisterValidation("stringMax", stringMaxValidator)
	validate.RegisterValidation("trueFalse", trueFalseValidator)
	validate.RegisterValidation("bool", booleanCheck)
	validate.RegisterValidation("dayNameEN", dayNameEN)
	validate.RegisterValidation("dayNameIN", dayNameIN)
	validate.RegisterValidation("uuid_valid_not_nil", uuidValidAndNotNil)
	return validate
}

func phoneNumberValidator(fl validator.FieldLevel) bool {
	pattern := `^(\+62|62|0)(\d{3,4})(\d{3,4})(\d{4})$`
	bo, _ := regexp.MatchString(pattern, fl.Field().String())
	return bo
}

func nikValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	pattern := `^(1[1-9]|21|[37][1-6]|5[1-3]|6[1-5]|[89][12])\d{2}\d{2}([04][1-9]|[1256][0-9]|[37][01])(0[1-9]|1[0-2])\d{2}\d{4}$`
	match, _ := regexp.MatchString(pattern, value)
	return match
}

func parentValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return value == "father" || value == "mother"
}

func stringMaxValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	maxLength := fl.Param()

	if len(value) > 0 && len(value) <= utils.StringToInt(maxLength) {
		return true
	}
	return false
}

func uuidValidAndNotNil(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		value := field.String()
		id, err := uuid.Parse(value)
		if err != nil || id == uuid.Nil {
			return false
		}
		return true

	case reflect.Array:
		if field.Type() == reflect.TypeOf(uuid.UUID{}) {
			id := field.Interface().(uuid.UUID)
			return id != uuid.Nil
		}
	}
	return false
}

func trueFalseValidator(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return value == "true" || value == "false"
}

func booleanCheck(fl validator.FieldLevel) bool {
	return fl.Field().Kind() == reflect.Bool && fl.Field().Interface() != nil
}

var checkEN = map[string]bool{
	"monday":    true,
	"tuesday":   true,
	"wednesday": true,
	"thursday":  true,
	"friday":    true,
	"saturday":  true,
	"sunday":    true,
}

func dayNameEN(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	value = strings.ToLower(value)

	_, ok := checkEN[value]

	return ok
}

var checkIN = map[string]bool{
	"senin":  true,
	"selasa": true,
	"rabu":   true,
	"kamis":  true,
	"jumat":  true,
	"sabtu":  true,
	"minggu": true,
}

func dayNameIN(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	value = strings.ToLower(value)

	_, ok := checkIN[value]

	return ok
}
