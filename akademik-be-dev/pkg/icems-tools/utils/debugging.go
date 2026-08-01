package utils

import (
	"fmt"
	"log"
)

var APP_DEBUG bool

func UpdateAppDebug(debug bool) {
	APP_DEBUG = debug
}

func CreateMsgDebuging(err string, id string, location string) string {
	return fmt.Sprintf("%s : %s location %s", err, id, location)
}

func PrintMsgDebuging(msg string) error {
	if APP_DEBUG {
		log.Print(msg)
	}
	return nil
}
