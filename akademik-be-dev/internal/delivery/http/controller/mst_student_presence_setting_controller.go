// Package controller
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

type MstStudentPresenceSettingController struct {
	log      *logrus.Logger
	validate *validator.Validate

	//
	mstStudentPresenceSettingService *servicemodel.MstStudentPresenceSettingService
}

func NewMstStudentPresenceSettingController(
	log *logrus.Logger,
	validate *validator.Validate,

	//
	mstStudentPresenceSettingService *servicemodel.MstStudentPresenceSettingService,
) *MstStudentPresenceSettingController {
	return &MstStudentPresenceSettingController{
		log:      log,
		validate: validate,
		//
		mstStudentPresenceSettingService: mstStudentPresenceSettingService,
	}
}

func (c *MstStudentPresenceSettingController) Create(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	var req dto.MstStudentPresenceSettingCreateRequest

	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}

	err := c.mstStudentPresenceSettingService.Create(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Create MstStudentPresenceSettingController",
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

func (c *MstStudentPresenceSettingController) Duplicate(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	var req dto.MstStudentPresenceSettingDuplicateRequest

	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}

	err := c.mstStudentPresenceSettingService.Duplicate(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Duplicate MstStudentPresenceSettingController",
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

func (c *MstStudentPresenceSettingController) GetAllWithCount(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	var pageble pageable.PageableStudentSettingGets
	if err := ctx.QueryParser(&pageble); err != nil {
		c.log.Warnf("Failed to parse query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	_, err := utils.ValidateAndPrepareRequest(&pageble)
	if err != nil {
		c.log.Warnf("Failed to query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	res, err := c.mstStudentPresenceSettingService.GetAllWithCount(ctx.Context(), pageble)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get all MstStudentPresenceSettingController",
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
