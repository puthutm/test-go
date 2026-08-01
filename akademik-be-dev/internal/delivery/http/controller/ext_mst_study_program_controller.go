package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto"
	servicemodel "unsia.ac.id/akademic_be/internal/service/model"
	"unsia.ac.id/akademic_be/pkg/auth"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstStudyProgramController struct {
	log                    *logrus.Logger
	MstStudyProgramService *servicemodel.MstStudyProgramService
}

func NewMstStudyProgramController(
	log *logrus.Logger,
	MstStudyProgramService *servicemodel.MstStudyProgramService,
) *MstStudyProgramController {
	return &MstStudyProgramController{
		log:                    log,
		MstStudyProgramService: MstStudyProgramService,
	}
}

/* Create */
/* Read */
func (c *MstStudyProgramController) GetByLecturerIDandActiveAcademicPeriod(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	distributionOfStudyPrograms, err := c.MstStudyProgramService.GetByLecturerIDandActiveAcademicPeriod(ctx.Context())
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get distribution of study programs",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, distributionOfStudyPrograms),
	)
}

/* Update */
/* Delete */
