package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto"
	servicemodel "unsia.ac.id/akademic_be/internal/service/model"
	"unsia.ac.id/akademic_be/pkg/auth"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstBiodataOfUserController struct {
	log                     *logrus.Logger
	mstBiodataOfUserService *servicemodel.MstBiodataOfUserService
}

func NewMstBiodataOfUserController(
	log *logrus.Logger,
	mstBiodataOfUserService *servicemodel.MstBiodataOfUserService,
) *MstBiodataOfUserController {
	return &MstBiodataOfUserController{
		log:                     log,
		mstBiodataOfUserService: mstBiodataOfUserService,
	}
}

/* Create */
/* Read */
func (c *MstBiodataOfUserController) GetByUserID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	UserID := user.ID

	sksLimit, err := c.mstBiodataOfUserService.GetByUserID(ctx.Context(), UserID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                UserID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": fmt.Sprintf("Failed to get data By User ID : %s", UserID),
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, sksLimit))
}

/* Update */
/* Delete */
