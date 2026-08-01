package dto

import (
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
)

var APP_DEBUG bool

type Response[t any] struct {
	Error     bool   `json:"error"`
	Data      t      `json:"data"`
	Message   string `json:"message"`
	Status    int    `json:"status"`
	CodeError string `json:"code_error,omitempty"`
}

func CreatePageableData[T any](page int, limit int, totalPage int, totalData int64, data []T) pageable.PageableResponse[T] {
	return pageable.PageableResponse[T]{
		Data: data,
		Metadata: pageable.Metadata{
			TotalData: totalData,
			TotalPage: totalPage,
			Page:      page,
			Size:      limit,
		},
	}
}

func CreateSuccess[T any](status int, message string, data T) Response[T] {
	return Response[T]{
		Error:   false,
		Data:    data,
		Message: message,
		Status:  status,
	}
}

func CreateError(status int, code string, message string) Response[string] {
	var res Response[string]
	if !APP_DEBUG && status >= 500 {
		res = Response[string]{
			Error:     true,
			Data:      "",
			Message:   msg.ErrInternalServer.Error(),
			Status:    status,
			CodeError: code,
		}
		return res
	}
	res = Response[string]{
		Error:     true,
		Data:      "",
		Message:   message,
		Status:    status,
		CodeError: code,
	}

	return res
}

func CreateErrorData(status int, message string, data map[string]string) Response[map[string]string] {
	return Response[map[string]string]{
		Error:   true,
		Data:    data,
		Message: message,
		Status:  status,
	}
}

func CreateErrorValidation(data map[string]string) Response[map[string]string] {
	return Response[map[string]string]{
		Error:   true,
		Data:    data,
		Message: "Validation failed.",
		Status:  422, // fiber.StatusUnprocessableEntity
	}
}
