package restapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"unsia.ac.id/akademic_be/pkg/icems-tools/dto"
	"unsia.ac.id/akademic_be/pkg/icems-tools/gateway/model"
	"unsia.ac.id/akademic_be/pkg/icems-tools/utils"
)

type ContentType string
type RestError error
type MethodHttpSend string

const (
	ContentTypeJson     ContentType = "application/json"
	ContentTypeForm     ContentType = "application/x-www-form-urlencoded"
	ContentTypeMulipart ContentType = "multipart/form-data"
)

const (
	PostSend  MethodHttpSend = http.MethodPost
	PutSend   MethodHttpSend = http.MethodPut
	PatchSend MethodHttpSend = http.MethodPatch
)

type Rest[T model.Response] struct {
	ApiKey string
	URL    string
	Log    *logrus.Logger
}

func (r *Rest[T]) GetData(header map[string]string) (dto.Response[T], error) {
	var dataResponse dto.Response[T]
	agent := fiber.Get(r.URL)

	if r.ApiKey != "" {
		agent = agent.Set("X-API-KEY", r.ApiKey)
	}

	for k, v := range header {
		agent = agent.Set(k, v)
	}

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
			"message":     "get data failed",
		}).Error(errs)
		return dataResponse, errs[0]
	}
	if statusCode > 500 {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
			"status":      statusCode,
		}).Error("unexpected response status")
	}
	err := json.Unmarshal(body, &dataResponse)
	if err != nil {
		return dataResponse, err
	}
	return dataResponse, nil
}

func (r *Rest[T]) GetDataWithParamsValue(header, query map[string]string, argsParam ...any) (dto.Response[T], error) {
	var dataResponse dto.Response[T]
	var agent *fiber.Agent

	countArgs := strings.Count(r.URL, "%s")
	if countArgs > 0 {
		if countArgs == len(argsParam) {
			agent = fiber.Get(fmt.Sprintf(r.URL, argsParam...))
		} else {
			return dataResponse, fmt.Errorf("url have argumen for params url, url have %d args", countArgs)
		}
	} else {
		agent = fiber.Get(r.URL)
	}

	if r.ApiKey != "" {
		agent = agent.Set("X-API-KEY", r.ApiKey)
	}

	for k, v := range header {
		agent = agent.Set(k, v)
	}

	values := url.Values{}
	for k, v := range query {
		values.Set(k, v)
	}

	agent = agent.QueryString(values.Encode())

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
			"message":     "get data failed",
		}).Error(errs)
		return dataResponse, errs[0]
	}
	if statusCode > 500 {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
			"status":      statusCode,
		}).Error("unexpected response status")
	}
	err := json.Unmarshal(body, &dataResponse)
	if err != nil {
		return dataResponse, err
	}
	return dataResponse, nil
}

func (r *Rest[T]) GetDataWithParamsOrQuery(
	header, query map[string]string, urlParams string,
) (dto.Response[T], error) {
	var dataResponse dto.Response[T]
	var agent *fiber.Agent
	if urlParams != "" {
		agent = fiber.Get(r.URL + urlParams)
	} else {
		agent = fiber.Get(r.URL)
	}

	if r.ApiKey != "" {
		agent = agent.Set("X-API-KEY", r.ApiKey)
	}

	for k, v := range header {
		agent = agent.Set(k, v)
	}

	values := url.Values{}
	for k, v := range query {
		values.Set(k, v)
	}

	agent = agent.QueryString(values.Encode())

	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
			"message":     "get data failed",
		}).Error(errs)
		return dataResponse, errs[0]
	}
	if statusCode > 500 {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
			"status":      statusCode,
		}).Error("unexpected response status")
	}
	err := json.Unmarshal(body, &dataResponse)
	if err != nil {
		return dataResponse, err
	}
	return dataResponse, nil
}

func (r *Rest[T]) SendDataJson(header map[string]string, req model.Request, argsParam ...any) (dto.Response[T], error) {
	var dataResponse dto.Response[T]
	var agent *fiber.Agent

	countArgs := strings.Count(r.URL, "%s")
	if countArgs > 0 {
		if countArgs == len(argsParam) {
			agent = fiber.Post(fmt.Sprintf(r.URL, argsParam...))
		} else {
			return dataResponse, fmt.Errorf("url have argumen for params url, url have %d args", countArgs)
		}
	} else {
		agent = fiber.Post(r.URL)
	}

	agent = agent.Set("Content-Type", string(ContentTypeJson))
	if r.ApiKey != "" {
		agent = agent.Set("X-API-KEY", r.ApiKey)
	}

	for k, v := range header {
		agent = agent.Set(k, v)
	}

	reqJson, _ := json.Marshal(req)
	agent = agent.Body(reqJson)
	return r.SendAgent(agent)
}

func (r *Rest[T]) SendDataForm(header map[string]string, req model.Request, argsParam ...any) (dto.Response[T], error) {
	var dataResponse dto.Response[T]
	var agent *fiber.Agent

	countArgs := strings.Count(r.URL, "%s")
	if countArgs > 0 {
		if countArgs == len(argsParam) {
			agent = fiber.Post(fmt.Sprintf(r.URL, argsParam...))
		} else {
			return dataResponse, fmt.Errorf("url have argumen for params url, url have %d args", countArgs)
		}
	} else {
		agent = fiber.Post(r.URL)
	}

	if r.ApiKey != "" {
		agent = agent.Set("X-API-KEY", r.ApiKey)
	}

	for k, v := range header {
		agent = agent.Set(k, v)
	}

	args := fiber.AcquireArgs()
	defer fiber.ReleaseArgs(args)
	results := utils.StructToMapString(req)
	for k, v := range results {
		args.Set(k, v)
	}
	agent = agent.Form(args)
	return r.SendAgent(agent)
}

func (r *Rest[T]) SendDataMultiForm(header map[string]string, req model.RequestMultiForm) (dto.Response[T], error) {
	agent := fiber.Post(r.URL)
	if r.ApiKey != "" {
		agent = agent.Set("X-API-KEY", r.ApiKey)
	}

	for k, v := range header {
		agent = agent.Set(k, v)
	}
	multiPartFile := req.GetFile()
	if multiPartFile != nil {
		fmt.Println(multiPartFile.Filename)
	}

	file, err := multiPartFile.Open()
	if err != nil {
		r.Log.WithFields(logrus.Fields{
			"service_url":    r.URL,
			"message_custom": "open file failed",
		}).Error(err)
		var zero dto.Response[T]
		return zero, err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
		}).Error("read file failed")
		var zero dto.Response[T]
		return zero, err
	}
	results := utils.StructToMapString(req)
	ff1 := &fiber.FormFile{
		Fieldname: "file",
		Name:      req.GetFileName(),
		Content:   fileBytes,
	}

	args := fiber.AcquireArgs()
	defer fiber.ReleaseArgs(args)

	for k, v := range results {
		args.Set(k, v)
	}
	agent = agent.
		FileData(ff1).
		MultipartForm(args)
	return r.SendAgent(agent)
}

func (r *Rest[T]) SendAgent(agent *fiber.Agent) (dto.Response[T], error) {
	var dataResponse dto.Response[T]
	var raw dto.Response[json.RawMessage]
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
			"message":     "send data failed",
		}).Error(errs)
		return dataResponse, errs[0]
	}
	if statusCode >= 500 {
		r.Log.WithFields(logrus.Fields{
			"service_url": r.URL,
			"status":      statusCode,
		}).Error("unexpected response status")
	}

	err := json.Unmarshal(body, &raw)
	if err != nil {
		r.Log.WithError(err).WithField("body", string(body)).Error("failed to unmarshal raw response")
		return dataResponse, err
	}

	dataResponse.Error = raw.Error
	dataResponse.Message = raw.Message
	dataResponse.Status = raw.Status

	if string(raw.Data) != `""` && len(raw.Data) > 0 {
		err := json.Unmarshal(raw.Data, &dataResponse.Data)
		if err != nil {
			r.Log.WithError(err).WithField("data", string(raw.Data)).Error("failed to unmarshal typed data")
			return dataResponse, err
		}
	}
	return dataResponse, nil
}
