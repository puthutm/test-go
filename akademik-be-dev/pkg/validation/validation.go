package validation

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"unsia.ac.id/akademic_be/pkg/utils"
)

func Validate[T any](validate *validator.Validate, data T) map[string]string {
	err := validate.Struct(data)
	res := make(map[string]string)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			switch v.ActualTag() {
			case "required":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s is required", utils.SplitCamelCase(v.Field()))
				continue
			case "len":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must have length between %s and %s", utils.SplitCamelCase(v.Field()), v.Param(), v.Value())
				continue
			case "email":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must be a valid email address", utils.SplitCamelCase(v.Field()))
				continue
			case "min":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must be between %s and %s", utils.SplitCamelCase(v.Field()), v.Param(), v.Value())
				continue
			case "max":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must be less than or equal to %s", utils.SplitCamelCase(v.Field()), v.Value())
				continue
			case "stringMax":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must be at most %s characters long", utils.SplitCamelCase(v.Field()), v.Param())
				continue
			case "eq":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must be equal to %s", utils.SplitCamelCase(v.Field()), v.Param())
				continue
			case "gt":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must be greater than %s", utils.SplitCamelCase(v.Field()), v.Param())
				continue
			case "lt":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must be less than %s", utils.SplitCamelCase(v.Field()), v.Value())
				continue
			case "contains":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must contain %s", utils.SplitCamelCase(v.Field()), v.Param())
				continue
			case "isdivisibleby":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must be divisible by %s", utils.SplitCamelCase(v.Field()), v.Param())
				continue
			case "numeric":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must be a numeric value", utils.SplitCamelCase(v.Field()))
				continue
			case "isalpha":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must be a alphabetic string", utils.SplitCamelCase(v.Field()))
				continue
			case "parent":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must be either 'father' or 'mother'", utils.SplitCamelCase(v.Field()))
				continue
			case "trueFalse":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must be either 'true' or 'false'", utils.SplitCamelCase(v.Field()))
				continue
			case "uuid":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s Ensures the string is a valid UUID format", utils.SplitCamelCase(v.Field()))
				continue
			case "bool":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s must be a boolean format", utils.SplitCamelCase(v.Field()))
				continue
			case "uuid_valid_not_nil":
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("%s Ensures the string is a valid UUID format", utils.SplitCamelCase(v.Field()))
				continue
			default:
				res[utils.CamelCaseToSnakeCase(v.StructField())] = fmt.Sprintf("Validation error for field %s: %s", CamelCaseToReadable(v.Field()), v.ActualTag())
				continue
			}
		}
	}
	return res
}

func CamelCaseToReadable(input string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")

	output := re.ReplaceAllString(input, "${1} ${2}")

	words := strings.Split(cases.Lower(language.Tag{}).String(output), " ")

	if len(words) > 0 {
		words[0] = cases.Title(language.Tag{}).String(words[0])
	}

	return strings.Join(words, " ")
}

func ValidatePhoneNumber(phone string) (bool, error) {
	pattern := `^(\+62|62|0)(\d{3,4})(\d{3,4})(\d{4})$`
	return regexp.MatchString(pattern, phone)
}
