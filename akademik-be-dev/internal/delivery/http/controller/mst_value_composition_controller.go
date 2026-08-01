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

type MstValueCompositionController struct {
	log                        *logrus.Logger
	mstValueCompositionService *servicemodel.MstValueCompositionService
	validate                   *validator.Validate
}

func NewMstValueCompositionController(
	log *logrus.Logger,
	mstValueCompositionService *servicemodel.MstValueCompositionService,
	validate *validator.Validate,
) *MstValueCompositionController {
	return &MstValueCompositionController{
		log:                        log,
		mstValueCompositionService: mstValueCompositionService,
		validate:                   validate,
	}
}

// TODO: Create
func (c *MstValueCompositionController) Create(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	var req dto.MstValueCompositionRequest
	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}
	err := c.mstValueCompositionService.Create(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Create MstValueCompositionController",
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
func (c *MstValueCompositionController) UpdateByID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	var req dto.MstValueCompositionUpdate
	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	ID := ctx.Params("value_composition_id")
	idString, _ := utils.StringToUuid(ID)
	req.ID = idString
	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}
	err := c.mstValueCompositionService.UpdateByID(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Update By ID MstValueCompositionController",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusOK, msg.SuccessUpdate, ""),
	)
}

// TODO: Delete
func (c *MstValueCompositionController) DeleteByID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	ID := ctx.Params("value_composition_id")
	err := c.mstValueCompositionService.DeleteByID(ctx.Context(), ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"ID":                  ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Delete By ID MstValueCompositionController",
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
func (c *MstValueCompositionController) RestoreByID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	ID := ctx.Params("value_composition_id")
	err := c.mstValueCompositionService.RestoreByID(ctx.Context(), ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"ID":                  ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Restore By ID MstValueCompositionController",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusOK, msg.SuccessRestore, ""),
	)
}

// TODO: Read
func (c *MstValueCompositionController) GetByID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	ID := ctx.Params("value_composition_id")
	res, err := c.mstValueCompositionService.GetByID(ctx.Context(), ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"ID":                  ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get By ID MstValueCompositionController",
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

func (c *MstValueCompositionController) GetAcademicPeriodsWithCount(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	var pageble pageable.PageableRequest
	if err := ctx.QueryParser(&pageble); err != nil {
		c.log.Warnf("Failed to parse query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	_, err := utils.ValidateAndPrepareRequest(&pageble)
	if err != nil {
		c.log.Warnf("Failed to query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	res, err := c.mstValueCompositionService.GetAcademicPeriodsWithCount(ctx.Context(), pageble)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get academic periods with count MstValueCompositionController",
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

func (c *MstValueCompositionController) Duplicate(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	var req dto.MstValueCompositionDuplicateRequest
	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}
	err := c.mstValueCompositionService.DuplicateByAcademicPeriodID(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Duplicate MstValueCompositionController",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusOK, msg.SuccessCreate, ""),
	)
}

func (c *MstValueCompositionController) GetAllWithCount(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	var pageble pageable.PageableRequestValueComposition
	if err := ctx.QueryParser(&pageble); err != nil {
		c.log.Warnf("Failed to parse query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	_, err := utils.ValidateAndPrepareRequest(&pageble)
	if err != nil {
		c.log.Warnf("Failed to query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	res, err := c.mstValueCompositionService.GetAllWithCount(ctx.Context(), pageble)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get all MstValueCompositionController",
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

func (c *MstValueCompositionController) GetAllTrashWithCount(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	var pageble pageable.PageableRequestValueComposition
	if err := ctx.QueryParser(&pageble); err != nil {
		c.log.Warnf("Failed to parse query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	_, err := utils.ValidateAndPrepareRequest(&pageble)
	if err != nil {
		c.log.Warnf("Failed to query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	res, err := c.mstValueCompositionService.GetAllTrashWithCount(ctx.Context(), pageble)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get all trash MstValueCompositionController",
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
