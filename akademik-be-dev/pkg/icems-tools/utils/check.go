package utils

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"unsia.ac.id/akademic_be/pkg/icems-tools/dto"
	msg "unsia.ac.id/akademic_be/pkg/icems-tools/dto/message"
	"unsia.ac.id/akademic_be/pkg/icems-tools/gateway/model"
)

func CheckRestError[T model.Response](
	log *logrus.Logger,
	data dto.Response[T],
	userID, serviceName, messageI string,
) error {
	if data.Error {
		log.WithFields(logrus.Fields{
			"service": serviceName,
			"message": messageI,
			"user-id": userID,
		}).Error(data.Message)
		return fiber.NewError(data.Status, data.Message)
	}

	if reflect.ValueOf(data.Data).IsZero() {
		log.WithFields(logrus.Fields{
			"service": serviceName,
			"message": messageI,
			"user-id": userID,
		}).Error(msg.ErrNotFound)
		return fiber.NewError(fiber.StatusNotFound, messageI+" data not found with user -> "+userID)
	}
	return nil
}

func CheckRestErrorForDataZero[T model.Response](
	log *logrus.Logger,
	data dto.Response[T],
	userID, serviceName, messageI string,
) error {
	if data.Error {
		log.WithFields(logrus.Fields{
			"service": serviceName,
			"message": messageI,
			"user-id": userID,
		}).Error(data.Message)
		return fiber.NewError(data.Status, data.Message)
	}
	return nil
}

var CheckDayEN = map[string]int{
	"monday":    0,
	"tuesday":   1,
	"wednesday": 2,
	"thursday":  3,
	"friday":    4,
	"saturday":  5,
	"sunday":    6,
}

var CheckDayIN = map[string]int{
	"senin":  0,
	"selasa": 1,
	"rabu":   2,
	"kamis":  3,
	"jumat":  4,
	"sabtu":  5,
	"minggu": 6,
}

func ReverseMap(m map[string]int) map[int]string {
	reversed := make(map[int]string)
	for k, v := range m {
		reversed[v] = k
	}
	return reversed
}

var (
	DayEN = ReverseMap(CheckDayEN)
	DayIN = ReverseMap(CheckDayIN)
)
