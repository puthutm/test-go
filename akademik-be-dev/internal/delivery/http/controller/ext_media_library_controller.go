package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gitlab.unsia.ac.id/icems/storage-library-be/folders"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto"
	servicemodel "unsia.ac.id/akademic_be/internal/service/model"
	"unsia.ac.id/akademic_be/pkg/auth"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstMediaLibraryController struct {
	log                    *logrus.Logger
	mstMediaLibraryService *servicemodel.MstMediaLibraryService
}

func NewMstMediaLibraryController(
	log *logrus.Logger,
	mstMediaLibraryService *servicemodel.MstMediaLibraryService,
) *MstMediaLibraryController {
	return &MstMediaLibraryController{
		log:                    log,
		mstMediaLibraryService: mstMediaLibraryService,
	}
}

// Create
func (c *MstMediaLibraryController) Create(ctx *fiber.Ctx) error {
	return nil
}

func (c *MstMediaLibraryController) CreateFolderRoot(ctx *fiber.Ctx) error {
	return nil
}

func (c *MstMediaLibraryController) CreateFolderByParent(ctx *fiber.Ctx) error {
	return nil
}

// GET
func (c *MstMediaLibraryController) GetFileByFolderAndSubject(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	var req folders.MstFolderRequest_GetFileByFolderAndSubject

	if err := ctx.QueryParser(&req); err != nil {
		c.log.Warnf("Failed to parse query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	finalProjectProposals, err := c.mstMediaLibraryService.GetFileByFolderAndSubject(
		ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"req":                 req,
			"message-error-debug": "Failed to get file by folder and subject",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, finalProjectProposals))
}

func (c *MstMediaLibraryController) GetFolderByParent(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	parentID := ctx.Params("parent_id")

	finalProjectProposals, err := c.mstMediaLibraryService.GetFolderByParent(
		ctx.Context(), parentID,
	)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"id":                  parentID,
			"message-error-debug": "Failed to get folder by parentID",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, finalProjectProposals))
}

func (c *MstMediaLibraryController) GetFolderRoot(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	finalProjectProposals, err := c.mstMediaLibraryService.GetFolderRoot(
		ctx.Context(),
	)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to get folder root",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, finalProjectProposals))
}
