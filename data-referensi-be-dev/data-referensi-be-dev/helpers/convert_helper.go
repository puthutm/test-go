package helpers

import (
	"regexp"
	"strings"
	"time"
)

/* Convert Camel Case To Snake Case */
func ConvertCCToSC(str string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}

func StringToDate(stringDate string) (time.Time, error) {
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, stringDate)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
