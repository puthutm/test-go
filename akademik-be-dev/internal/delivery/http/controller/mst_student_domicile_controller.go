package controller

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto"
	servicemodel "unsia.ac.id/akademic_be/internal/service/model"
	"unsia.ac.id/akademic_be/pkg/auth"
	"unsia.ac.id/akademic_be/pkg/utils"
	"unsia.ac.id/akademic_be/pkg/validation"
)

type MstStudentDomicileController struct {
	log                       *logrus.Logger
	mstStudentDomicileService *servicemodel.MstStudentDomicileService
	mstStudentBioService      *servicemodel.MstStudentBioService
	validate                  *validator.Validate
}

func NewMstStudentDomicileController(
	log *logrus.Logger,
	mstStudentDomicileService *servicemodel.MstStudentDomicileService,
	mstStudentBioService *servicemodel.MstStudentBioService,
	validate *validator.Validate,
) *MstStudentDomicileController {
	return &MstStudentDomicileController{
		log:                       log,
		mstStudentDomicileService: mstStudentDomicileService,
		mstStudentBioService:      mstStudentBioService,
		validate:                  validate,
	}
}

/* Create */
/* Read */
func (c *MstStudentDomicileController) GetByStudentID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	studentBiodata, err := c.mstStudentBioService.GetGeneralByID(ctx.Context(), "", user.ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to get data student biodata",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}
	StudentID := studentBiodata.ID.String()

	domicile, err := c.mstStudentDomicileService.GetByStudentID(ctx.Context(), StudentID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to get data student domicile",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, domicile))
}

/* Update */
func (c *MstStudentDomicileController) UpdateByStudentID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	var req dto.MstStudentDomicileRequest

	studentBiodata, err := c.mstStudentBioService.GetGeneralByID(ctx.Context(), "", user.ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to get data student biodata",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	StudentID := studentBiodata.ID.String()
	req.StudentID = StudentID

	oldDomicile, err := c.mstStudentDomicileService.GetByStudentID(ctx.Context(), StudentID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to get data old student domicile",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	if oldDomicile.ID == uuid.Nil {
		ID := utils.GenerateUUID()
		req.ID = ID.String()
	} else {
		req.ID = oldDomicile.ID.String()
	}

	if err := ctx.BodyParser(&req); err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": fmt.Sprintf("Failed to parse body : %+v", err),
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(dto.CreateErrorValidation(fails))
	}

	err = c.mstStudentDomicileService.UpdateByStudentID(ctx.Context(), req)
	if err != nil {
		createMsg := utils.CreateMsgDebuging(err.Error(), StudentID, utils.ErrorLocation())
		utils.PrintMsgDebuging(createMsg)
		c.log.Warnf(createMsg)

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	updatedDomicile, err := c.mstStudentDomicileService.GetByStudentID(ctx.Context(), StudentID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to get data updated student domicile",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessUpdate, updatedDomicile))
}

/* Delete */
