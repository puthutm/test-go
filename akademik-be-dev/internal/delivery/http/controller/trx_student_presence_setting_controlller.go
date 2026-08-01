// Package controller
package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	icemstutil "unsia.ac.id/akademic_be/pkg/icems-tools/utils"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/delivery/http/middleware"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	servicemodel "unsia.ac.id/akademic_be/internal/service/model"
	"unsia.ac.id/akademic_be/pkg/auth"
	"unsia.ac.id/akademic_be/pkg/utils"
	"unsia.ac.id/akademic_be/pkg/validation"
)

type TrxStudentPresenceSettingController struct {
	log      *logrus.Logger
	validate *validator.Validate

	trxStudentPresenceSettingService *servicemodel.TrxStudentPresenceSettingService
}

func NewTrxStudentPresenceSettingController(
	log *logrus.Logger,
	validate *validator.Validate,

	trxStudentPresenceSettingService *servicemodel.TrxStudentPresenceSettingService,
) *TrxStudentPresenceSettingController {
	return &TrxStudentPresenceSettingController{
		log:                              log,
		validate:                         validate,
		trxStudentPresenceSettingService: trxStudentPresenceSettingService,
	}
}

func (c *TrxStudentPresenceSettingController) CreateOrUpdate(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := middleware.GetUserClaimsCtx(ctx.Context())

	var req dto.TrxStudentPresenceSettingCreateOrUpdateRequest

	academicPeriodID, _ := icemstutil.StringToUuid(ctx.Params("academic_period_id"))
	subjectID, _ := icemstutil.StringToUuid(ctx.Params("subject_id"))

	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	req.AcademicPeriodeID = academicPeriodID
	req.SubjectID = subjectID

	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}

	err := c.trxStudentPresenceSettingService.CreateOrUpdate(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Create Or Update trxStudentPresenceSettingController",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusCreated, msg.SuccessCreate, ""),
	)
}

func (c *TrxStudentPresenceSettingController) CreateOrUpdateStudentPresence(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := middleware.GetUserClaimsCtx(ctx.Context())

	sessionID, _ := icemstutil.StringToUuid(ctx.Params("session_id"))
	var req dto.TrxStudentPresenceSaveParamBySessionRequest

	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	req.SessionID = sessionID

	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}

	err := c.trxStudentPresenceSettingService.CreateOrUpdateStudentPresence(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Create Or Update studentPresence trxStudentPresenceSettingController",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusCreated, msg.SuccessCreate, ""),
	)
}

func (c *TrxStudentPresenceSettingController) CreateOrUpdateStudentPresenceSlice(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := middleware.GetUserClaimsCtx(ctx.Context())

	sessionID, _ := icemstutil.StringToUuid(ctx.Params("session_id"))
	var req dto.TrxStudentPresenceSliceSaveParamBySessionRequest

	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	req.SessionID = sessionID

	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}

	err := c.trxStudentPresenceSettingService.CreateOrUpdateStudentPresenceSlice(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Create Or Update studentPresence slice trxStudentPresenceSettingController",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusCreated, msg.SuccessCreate, ""),
	)
}

func (c *TrxStudentPresenceSettingController) GetComponentForLecturer(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	var req dto.TrxStudentPresenceGetForLecturerRequest
	if err := ctx.QueryParser(&req); err != nil {
		c.log.Warnf("Failed to parse query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	academicPeriodID := ctx.Params("academic_period_id")
	subjectID := ctx.Params("subject_id")

	req.AcademicPeriodeID = academicPeriodID
	req.SubjectID = subjectID

	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}

	res, err := c.trxStudentPresenceSettingService.GetPresenceComponentForLecturer(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get component for lecturer MstStudentPresenceSettingController",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, res),
	)
}

func (c *TrxStudentPresenceSettingController) GetComponentBySession(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	sessionID := ctx.Params("session_id")

	res, err := c.trxStudentPresenceSettingService.GetPresenceComponentBySession(ctx.Context(), sessionID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get component by session MstStudentPresenceSettingController",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, res),
	)
}

func (c *TrxStudentPresenceSettingController) GetSessionPresenceByClassID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	classID := ctx.Params("class_id")
	if classID == "" {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(dto.CreateError(fiber.StatusBadRequest, codeError, "class_id is required"))
	}

	res, err := c.trxStudentPresenceSettingService.GetSessionPresenceByClassID(ctx.Context(), classID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to get session presence by class ID",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).
			JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).
		JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, res))
}

func (c *TrxStudentPresenceSettingController) GetStudentPresenceBySessionWithCount(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	sessionID := ctx.Params("session_id")

	var pageble pageable.PageableStudentPresenceBySession
	if err := ctx.QueryParser(&pageble); err != nil {
		c.log.Warnf("Failed to parse query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	pageble.SessionID = sessionID

	_, err := utils.ValidateAndPrepareRequest(&pageble)
	if err != nil {
		c.log.Warnf("Failed to query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	fails := validation.Validate(c.validate, &pageble)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}
	res, err := c.trxStudentPresenceSettingService.GetStudentPresenceBySessionWithCount(ctx.Context(), pageble)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to GetStudentPresenceBySessionWithCount MstClassLecturerController",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, res),
	)
}
