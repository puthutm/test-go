package controller

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"unsia.ac.id/akademic_be/internal/config"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/service"
	servicemodel "unsia.ac.id/akademic_be/internal/service/model"
	"unsia.ac.id/akademic_be/pkg/auth"
	"unsia.ac.id/akademic_be/pkg/utils"
	"unsia.ac.id/akademic_be/pkg/validation"
)

type MstStudentEducationController struct {
	log                        *logrus.Logger
	mstStudentEducationService *servicemodel.MstStudentEducationService
	mstStudentBioService       *servicemodel.MstStudentBioService
	validate                   *validator.Validate
	storage                    *service.StorageService
	cnf                        *config.Config
}

func NewMstStudentEducationController(
	log *logrus.Logger,
	mstStudentEducationService *servicemodel.MstStudentEducationService,
	mstStudentBioService *servicemodel.MstStudentBioService,
	validate *validator.Validate,
	storage *service.StorageService,
	cnf *config.Config,
) *MstStudentEducationController {
	return &MstStudentEducationController{
		log:                        log,
		mstStudentEducationService: mstStudentEducationService,
		mstStudentBioService:       mstStudentBioService,
		validate:                   validate,
		storage:                    storage,
		cnf:                        cnf,
	}
}

/* Create */
/* Read */
func (c *MstStudentEducationController) GetByStudentID(ctx *fiber.Ctx) error {
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

	education, err := c.mstStudentEducationService.GetByStudentID(ctx.Context(), StudentID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to get data student education",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	CertificateFileURL := education.CertificateFilepath
	education.CertificateFileURL = CertificateFileURL

	TranscriptsFileURL := education.TranscriptsFilepath
	education.TranscriptsFileURL = TranscriptsFileURL
	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, education))
}

/* Update */
func (c *MstStudentEducationController) UpdateByID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	var req dto.MstStudentEducationRequest

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

	oldEducation, err := c.mstStudentEducationService.GetByStudentID(ctx.Context(), StudentID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to get data old student education",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	if oldEducation.ID == uuid.Nil {
		ID := utils.GenerateUUID()
		req.ID = ID.String()
	} else {
		req.ID = oldEducation.ID.String()
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

	certificateFile, err := ctx.FormFile("certificate_file")
	if err != nil && err.Error() != "there is no uploaded file associated with the given key" {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": fmt.Sprintf("Error load file : %+v", err),
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, fmt.Sprintf("Error load file : %+v", err)))
	}

	subfolder := "student-" + req.StudentID + "/education"
	validateExtensions := []string{".pdf"}
	maxSize := int64(1)

	if certificateFile != nil {
		path, err := c.storage.UploadFileV3(
			ctx.Context(), certificateFile, true, subfolder, "", validateExtensions, maxSize,
		)
		if err != nil {
			utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
				"code-error":          codeError,
				"user":                user.ID,
				"location":            utils.ErrorLocation(),
				"message-error-debug": fmt.Sprintf("Failed to upload file : %+v", err),
			})

			statusCode := fiber.StatusInternalServerError
			if fiberErr, ok := err.(*fiber.Error); ok {
				statusCode = fiberErr.Code
			}

			return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, fmt.Sprintf("Failed to upload file certificate : %+v. Support file extension .pdf", err)))

		}
		req.CertificateFilepath = &path
	} else {
		req.CertificateFilepath = oldEducation.CertificateFilepath
	}

	transcriptsFile, err := ctx.FormFile("transcripts_file")
	if err != nil && err.Error() != "there is no uploaded file associated with the given key" {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": fmt.Sprintf("Error load file : %+v", err),
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, fmt.Sprintf("Error load file : %+v", err)))
	}
	if transcriptsFile != nil {
		path, err := c.storage.UploadFileV3(
			ctx.Context(), certificateFile, true, subfolder, "", validateExtensions, maxSize,
		)
		if err != nil {
			utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
				"code-error":          codeError,
				"user":                user.ID,
				"location":            utils.ErrorLocation(),
				"message-error-debug": fmt.Sprintf("Failed to upload file : %+v", err),
			})

			statusCode := fiber.StatusInternalServerError
			if fiberErr, ok := err.(*fiber.Error); ok {
				statusCode = fiberErr.Code
			}

			return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, fmt.Sprintf("Failed to upload file transcripts : %+v. Support file extension .pdf", err)))

		}
		req.TranscriptsFilepath = &path
	} else {
		req.TranscriptsFilepath = oldEducation.TranscriptsFilepath
	}

	err = c.mstStudentEducationService.UpdateByID(ctx.Context(), req)
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

	updatedEducation, err := c.mstStudentEducationService.GetByStudentID(ctx.Context(), StudentID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to get data updated student education",
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessUpdate, updatedEducation))
}

/* Delete */
