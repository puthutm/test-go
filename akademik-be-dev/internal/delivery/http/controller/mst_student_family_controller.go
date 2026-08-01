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

type MstStudentFamilyController struct {
	log                     *logrus.Logger
	mstStudentFamilyService *servicemodel.MstStudentFamilyService
	mstStudentBioService    *servicemodel.MstStudentBioService
	validate                *validator.Validate
}

func NewMstStudentFamilyController(
	log *logrus.Logger,
	mstStudentFamilyService *servicemodel.MstStudentFamilyService,
	mstStudentBioService *servicemodel.MstStudentBioService,
	validate *validator.Validate,
) *MstStudentFamilyController {
	return &MstStudentFamilyController{
		log:                     log,
		mstStudentFamilyService: mstStudentFamilyService,
		mstStudentBioService:    mstStudentBioService,
		validate:                validate,
	}
}

/* Create */
/* Read */
func (c *MstStudentFamilyController) GetByStudentID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	parentType := ctx.Params("parent_type")

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

	family, err := c.mstStudentFamilyService.GetByStudentID(ctx.Context(), StudentID, parentType)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to get data student family",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, family))
}

/* Update */
func (c *MstStudentFamilyController) UpdateByStudentID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	var req dto.MstStudentFamilyRequest

	parentType := ctx.Params("parent_type")
	req.Type = parentType

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

	oldFamily, err := c.mstStudentFamilyService.GetByStudentID(ctx.Context(), StudentID, parentType)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to get data old student family",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	if oldFamily.ID == uuid.Nil {
		ID := utils.GenerateUUID()
		req.ID = ID.String()
	} else {
		req.ID = oldFamily.ID.String()
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

	err = c.mstStudentFamilyService.UpdateByStudentID(ctx.Context(), req)
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

	updatedFamily, err := c.mstStudentFamilyService.GetByStudentID(ctx.Context(), StudentID, parentType)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to get data updated student family",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessUpdate, updatedFamily))
}

/* Delete */
