package utils

import (
	"fmt"
	"log"

	"unsia.ac.id/akademic_be/internal/dto"
)

func CreateMsgDebuging(err string, id string, location string) string {
	return fmt.Sprintf("%s : %s location %s", err, id, location)
}

func PrintMsgDebuging(msg string) error {
	if dto.APP_DEBUG {
		log.Print(msg)
	}

	return nil
}
