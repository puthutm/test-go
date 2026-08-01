package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	servicemodel "unsia.ac.id/akademic_be/internal/service/model"
	"unsia.ac.id/akademic_be/pkg/auth"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstLecturerController struct {
	log                *logrus.Logger
	mstLecturerService *servicemodel.MstLecturerService
}

func NewMstLecturerController(
	log *logrus.Logger,
	mstLecturerService *servicemodel.MstLecturerService,
) *MstLecturerController {
	return &MstLecturerController{
		log:                log,
		mstLecturerService: mstLecturerService,
	}
}

/* Read */
func (c *MstLecturerController) GetAllWithCountByProgramHeadID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	var pageble pageable.PageableRequest
	if err := ctx.QueryParser(&pageble); err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": fmt.Sprintf("Failed to parse query : %+v", err),
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	validRequest, err := utils.ValidateAndPrepareRequest(&pageble)
	if err != nil {
		return err
	}

	studentStudyPrograms, err := c.mstLecturerService.GetAllWithCountByProgramHeadID(ctx.Context(), validRequest)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to get data Student Study Program",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, studentStudyPrograms))
}
