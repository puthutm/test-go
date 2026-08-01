package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto"
	servicemodel "unsia.ac.id/akademic_be/internal/service/model"
	"unsia.ac.id/akademic_be/pkg/auth"
	"unsia.ac.id/akademic_be/pkg/utils"
	"unsia.ac.id/akademic_be/pkg/validation"
)

type MstStudentBioController struct {
	log                  *logrus.Logger
	mstStudentBioService *servicemodel.MstStudentBioService
	validate             *validator.Validate
}

func NewMstStudentBioController(
	log *logrus.Logger,
	mstStudentBioService *servicemodel.MstStudentBioService,
	validate *validator.Validate,
) *MstStudentBioController {
	return &MstStudentBioController{
		log:                  log,
		mstStudentBioService: mstStudentBioService,
		validate:             validate,
	}
}

// TODO: Update student bio
func (c *MstStudentBioController) UpdateBioGeneralOnlyUser(ctx *fiber.Ctx) error {
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	codeError := ctx.Locals("code-error").(string)

	var req dto.MstStudentBioUpdateOnlyUser
	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	idString, _ := utils.StringToUuid(userClaims.ID)
	req.UserID = idString
	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}

	err := c.mstStudentBioService.Update(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Update MstStudentBio General",
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

func (c *MstStudentBioController) UpdateBioCompletenessOnlyUser(ctx *fiber.Ctx) error {
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	codeError := ctx.Locals("code-error").(string)

	var req dto.MstStudentBioUpdateCompletenessOnlyUser
	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	idString, _ := utils.StringToUuid(userClaims.ID)
	req.UserID = idString
	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}

	filePath, err := ctx.FormFile("signature_path_file")
	if err != nil && err.Error() != "there is no uploaded file associated with the given key" {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": msg.ErrGetFile.Error(),
		})
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, msg.ErrGetFile.Error()))
	}
	req.SignaturePathFile = filePath

	err = c.mstStudentBioService.Update(ctx.Context(), req)
	if err != nil {
		c.log.Warnf("Failed to udpdate bio completeness: %v", err)
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Update MstStudentBio completeness",
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

func (c *MstStudentBioController) UpdateBioInformationOnlyUser(ctx *fiber.Ctx) error {
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	codeError := ctx.Locals("code-error").(string)

	var req dto.MstStudentBioUpdateInformationOnlyUser
	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	idString, _ := utils.StringToUuid(userClaims.ID)
	req.UserID = idString
	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}

	err := c.mstStudentBioService.Update(ctx.Context(), req)
	if err != nil {
		c.log.Warnf("Failed to udpdate bio information: %v", err)
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Update MstStudentBio information",
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

func (c *MstStudentBioController) UpdateBioDocumentOnlyUser(ctx *fiber.Ctx) error {
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	codeError := ctx.Locals("code-error").(string)

	var req dto.MstStudentBioUpdateDocumentOnlyUser
	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	idString, _ := utils.StringToUuid(userClaims.ID)
	req.UserID = idString
	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}

	bpjsEmploymentFilePath, err := ctx.FormFile("bpjs_employment_filepath")
	if err != nil && err.Error() != "there is no uploaded file associated with the given key" {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": msg.ErrGetFile.Error(),
		})
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, msg.ErrGetFile.Error()))
	}
	bpjsHealthcareFilePath, err := ctx.FormFile("bpjs_healthcare_filepath")
	if err != nil && err.Error() != "there is no uploaded file associated with the given key" {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": msg.ErrGetFile.Error(),
		})
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, msg.ErrGetFile.Error()))
	}
	npwpFilePath, err := ctx.FormFile("npwp_filepath")
	if err != nil && err.Error() != "there is no uploaded file associated with the given key" {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": msg.ErrGetFile.Error(),
		})
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, msg.ErrGetFile.Error()))
	}

	req.BPJSEmploymentFilePath = bpjsEmploymentFilePath
	req.BPJSHealthcareFilePath = bpjsHealthcareFilePath
	req.NPWPFilePath = npwpFilePath

	err = c.mstStudentBioService.Update(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Update MstStudentBio document",
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

func (c *MstStudentBioController) UpdateBioBankAccountOnlyUser(ctx *fiber.Ctx) error {
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	codeError := ctx.Locals("code-error").(string)

	var req dto.MstStudentBioUpdateBankAccountOnlyUser
	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	idString, _ := utils.StringToUuid(userClaims.ID)
	req.UserID = idString
	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}

	filePath, err := ctx.FormFile("account_file_path")
	if err != nil && err.Error() != "there is no uploaded file associated with the given key" {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": msg.ErrGetFile.Error(),
		})
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, msg.ErrGetFile.Error()))
	}
	req.AccountFilePath = filePath

	err = c.mstStudentBioService.Update(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Update MstStudentBio Bank Account",
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

// Delete
// Restore

// TODO:Get By ID
func (c *MstStudentBioController) GetBioGeneralOnlyUser(ctx *fiber.Ctx) error {
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	codeError := ctx.Locals("code-error").(string)

	res, err := c.mstStudentBioService.GetGeneralByID(ctx.Context(), "", userClaims.ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get MstStudentBio General",
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

func (c *MstStudentBioController) GetBioCompletenessOnlyUser(ctx *fiber.Ctx) error {
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	codeError := ctx.Locals("code-error").(string)

	res, err := c.mstStudentBioService.GetCompletenesByID(ctx.Context(), "", userClaims.ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get MstStudentBio completeness",
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

func (c *MstStudentBioController) GetBioInformationOnlyUser(ctx *fiber.Ctx) error {
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	codeError := ctx.Locals("code-error").(string)

	res, err := c.mstStudentBioService.GetinformationByID(ctx.Context(), "", userClaims.ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get MstStudentBio Information",
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

func (c *MstStudentBioController) GetBioDocumentOnlyUser(ctx *fiber.Ctx) error {
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	codeError := ctx.Locals("code-error").(string)

	res, err := c.mstStudentBioService.GetDocumentByID(ctx.Context(), "", userClaims.ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get MstStudentBio Document",
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

func (c *MstStudentBioController) GetBioBankAccountOnlyUser(ctx *fiber.Ctx) error {
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	codeError := ctx.Locals("code-error").(string)

	res, err := c.mstStudentBioService.GetBankAccountByID(ctx.Context(), "", userClaims.ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get MstStudentBio Bank Account",
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
