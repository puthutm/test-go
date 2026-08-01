package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	servicemodel "unsia.ac.id/akademic_be/internal/service/model"
	"unsia.ac.id/akademic_be/pkg/auth"
	"unsia.ac.id/akademic_be/pkg/utils"
	"unsia.ac.id/akademic_be/pkg/validation"
)

type MstClassParticipantController struct {
	log                        *logrus.Logger
	mstClassParticipantService *servicemodel.MstClassParticipantService
	validate                   *validator.Validate
}

func NewMstClassParticipantController(
	log *logrus.Logger,
	mstClassParticipantService *servicemodel.MstClassParticipantService,
	validate *validator.Validate,
) *MstClassParticipantController {
	return &MstClassParticipantController{
		log:                        log,
		mstClassParticipantService: mstClassParticipantService,
		validate:                   validate,
	}
}

// TODO: Create
func (c *MstClassParticipantController) Create(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	classID := ctx.Params("class_id")
	var req dto.MstClassParticipantRequest
	classIDNew, err := utils.StringToUuid(classID)
	if err != nil {
		c.log.Warnf("Failed to parse string to uuid: %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	req.ClassID = classIDNew
	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}
	err = c.mstClassParticipantService.Create(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Create MstClassParticipantController",
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

func (c *MstClassParticipantController) CreateByProgramHead(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	classID := ctx.Params("class_id")
	var req dto.MstClassParticipantRequest
	classIDNew, err := utils.StringToUuid(classID)
	if err != nil {
		c.log.Warnf("Failed to parse string to uuid: %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	req.ClassID = classIDNew
	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}
	err = c.mstClassParticipantService.CreateByProgramHead(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Create MstClassParticipantController",
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

// TODO: Update

// TODO: Delete
func (c *MstClassParticipantController) DeleteByID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	ID := ctx.Params("participant_id")
	err := c.mstClassParticipantService.DeleteByID(ctx.Context(), ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"ID":                  ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Delete By ID MstClassParticipantController",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusOK, msg.SuccessDelete, ""),
	)
}

func (c *MstClassParticipantController) DeleteByIDAndProgramHead(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	ID := ctx.Params("participant_id")
	err := c.mstClassParticipantService.DeleteByIDAndProgramHead(ctx.Context(), ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"ID":                  ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Delete By ID MstClassParticipantController",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusOK, msg.SuccessDelete, ""),
	)
}

// TODO: Restore

// TODO: Read
func (c *MstClassParticipantController) GetByID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	ID := ctx.Params("participant_id")
	res, err := c.mstClassParticipantService.GetByID(ctx.Context(), ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"ID":                  ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get By ID MstClassParticipantController",
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

func (c *MstClassParticipantController) GetAllWithCount(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	classID := ctx.Params("class_id")

	var pageble pageable.PageableRequestClassParticipant
	pageble.ClassID = classID
	if err := ctx.QueryParser(&pageble); err != nil {
		c.log.Warnf("Failed to parse query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	_, err := utils.ValidateAndPrepareRequest(&pageble)
	if err != nil {
		c.log.Warnf("Failed to query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	res, err := c.mstClassParticipantService.GetAllWithCount(ctx.Context(), pageble)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get all MstClassParticipantController",
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

func (c *MstClassParticipantController) GetAllWithCountByClassIDForLecturer(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	classID := ctx.Params("class_id")

	var pageble pageable.PageableRequestClassParticipant
	pageble.ClassID = classID
	if err := ctx.QueryParser(&pageble); err != nil {
		c.log.Warnf("Failed to parse query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	_, err := utils.ValidateAndPrepareRequest(&pageble)
	if err != nil {
		c.log.Warnf("Failed to query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	res, err := c.mstClassParticipantService.GetAllWithCountByClassIDForLecturer(ctx.Context(), pageble)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get all MstClassParticipantController",
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
