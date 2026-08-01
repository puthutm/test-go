package utils

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func GetDataReference[T any](viperCnf *viper.Viper, token, apiUrl string) *T {
	baseUrl := viperCnf.GetString("BASE_URL_DATA_REFERENSI")
	url := fmt.Sprintf("%s%s", baseUrl, apiUrl)

	agent := fiber.Get(url)
	agent = agent.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return nil
	}
	if statusCode != fiber.StatusOK {
		return nil
	}
	var something T
	err := json.Unmarshal(body, &something)
	if err != nil {
		return nil
	}
	return &something
}

type LogParam struct{}

func CreateLogSSO(logParam LogParam) {
}
