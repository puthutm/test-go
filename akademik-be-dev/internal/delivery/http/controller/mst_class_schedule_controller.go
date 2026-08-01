package controller

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	utilicems "unsia.ac.id/akademic_be/pkg/icems-tools/utils"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	servicemodel "unsia.ac.id/akademic_be/internal/service/model"
	"unsia.ac.id/akademic_be/pkg/auth"
	"unsia.ac.id/akademic_be/pkg/utils"
	"unsia.ac.id/akademic_be/pkg/validation"
)

type MstClassScheduleController struct {
	log                     *logrus.Logger
	mstClassScheduleService *servicemodel.MstClassScheduleService
	validate                *validator.Validate
}

func NewMstClassScheduleController(
	log *logrus.Logger,
	mstClassScheduleService *servicemodel.MstClassScheduleService,
	validate *validator.Validate,
) *MstClassScheduleController {
	return &MstClassScheduleController{
		log:                     log,
		mstClassScheduleService: mstClassScheduleService,
		validate:                validate,
	}
}

// TODO: Create
func (c *MstClassScheduleController) Create(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	classID := ctx.Params("class_id")
	var req dto.MstClassScheduleRequest
	classIDNew, err := utils.StringToUuid(classID)
	if err != nil {
		c.log.Warnf("Failed to parse string to uuid: %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	req.ClassID = classIDNew
	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}
	err = c.mstClassScheduleService.Create(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Create MstClassScheduleController",
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

func (c *MstClassScheduleController) CreateByProgramHead(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	classID := ctx.Params("class_id")
	var req dto.MstClassScheduleRequest
	classIDNew, err := utils.StringToUuid(classID)
	if err != nil {
		c.log.Warnf("Failed to parse string to uuid: %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	req.ClassID = classIDNew
	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}
	err = c.mstClassScheduleService.CreateByProgramHead(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Create MstClassScheduleController",
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
func (c *MstClassScheduleController) UpdateByDayTime(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	classID := ctx.Params("class_id")
	var req dto.MstClassScheduleUpdateRequest
	classIDNew, err := utils.StringToUuid(classID)
	if err != nil {
		c.log.Warnf("Failed to parse string to uuid: %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	req.ClassID = classIDNew

	if err := ctx.BodyParser(&req); err != nil {
		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateErrorValidation(fails))
	}
	err = c.mstClassScheduleService.UpdateByDayTime(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"request":             req,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to update MstClassScheduleController",
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

func (c *MstClassScheduleController) UpdateByIDForLecturer(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	ID := ctx.Params("class_schedule_id")

	var req dto.MstClassScheduleUpdateForLecturerRequest

	classScheduleID, err := utils.StringToUuid(ID)
	if err != nil {
		c.log.Warnf("Failed to parse string to uuid: %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	err = ctx.BodyParser(&req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": fmt.Sprintf("Failed to parse body : %+v", err),
		})

		c.log.Warnf("Failed to parse body : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	req.ID = classScheduleID

	fails := validation.Validate(c.validate, &req)
	if len(fails) > 0 {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(dto.CreateErrorValidation(fails))
	}

	MaterialAttachmentFile, err := ctx.FormFile("material_attachment_file")
	if err != nil && err.Error() != "there is no uploaded file associated with the given key" {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": msg.ErrGetFile.Error(),
		})
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, msg.ErrGetFile.Error()))
	}
	req.MaterialAttachmentFile = MaterialAttachmentFile

	AttendanceDocumentFile, err := ctx.FormFile("attendance_document_file")
	if err != nil && err.Error() != "there is no uploaded file associated with the given key" {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": msg.ErrGetFile.Error(),
		})
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, msg.ErrGetFile.Error()))
	}
	req.AttendanceDocumentFile = AttendanceDocumentFile

	JournalDocumentFile, err := ctx.FormFile("journal_document_file")
	if err != nil && err.Error() != "there is no uploaded file associated with the given key" {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": msg.ErrGetFile.Error(),
		})
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, msg.ErrGetFile.Error()))
	}
	req.JournalDocumentFile = JournalDocumentFile

	err = c.mstClassScheduleService.UpdateByIDForLecturer(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": fmt.Sprintf("Failed to get update Class By ID : %s", ID),
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessUpdate, ""))
}

// TODO: Delete
func (c *MstClassScheduleController) DeleteByID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	ID := ctx.Params("schedule_id")
	err := c.mstClassScheduleService.DeleteByID(ctx.Context(), ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"ID":                  ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Delete By ID MstClassScheduleController",
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

func (c *MstClassScheduleController) DeleteDayTime(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	var scheduleRequestGet dto.MstClassScheduleGetByRequest
	if err := ctx.QueryParser(&scheduleRequestGet); err != nil {
		c.log.Warnf("Failed to parse query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	err := c.mstClassScheduleService.DeleteDayTime(ctx.Context(), scheduleRequestGet)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"req":                 scheduleRequestGet,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Delete By ID MstClassScheduleController",
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

func (c *MstClassScheduleController) DeleteByScheduleTemplate(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	ID := ctx.Query("schedule_template_id")

	err := c.mstClassScheduleService.DeleteByScheduleTemplate(ctx.Context(), ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Delete By ID MstClassScheduleController",
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

// TODO Generate
func (c *MstClassScheduleController) GenerateByAcademic(ctx *fiber.Ctx) error {
	return c.wrapGenerate(ctx, c.mstClassScheduleService.GenerateByAcademic)
}

func (c *MstClassScheduleController) GenerateByProgramHead(ctx *fiber.Ctx) error {
	return c.wrapGenerate(ctx, c.mstClassScheduleService.GenerateByProgramHead)
}

// TODO: Read
func (c *MstClassScheduleController) GetByDayTime(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	ID := ctx.Params("class_id")
	var scheduleRequestGet dto.MstClassScheduleGetByRequest
	if err := ctx.QueryParser(&scheduleRequestGet); err != nil {
		c.log.Warnf("Failed to parse query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}
	res, err := c.mstClassScheduleService.GetByDayTime(ctx.Context(), scheduleRequestGet)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"class ID":            ID,
			"req":                 scheduleRequestGet,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get By ID MstClassScheduleController",
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

func (c *MstClassScheduleController) GetByID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	ID := ctx.Params("class_schedule_id")

	classSchedule, err := c.mstClassScheduleService.GetByID(ctx.Context(), ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": fmt.Sprintf("Failed to get data Class Schedule By ID : %s", ID),
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, classSchedule))
}

func (c *MstClassScheduleController) GetByIDForPresence(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	ID := ctx.Params("class_schedule_id")

	classSchedule, err := c.mstClassScheduleService.GetByID(ctx.Context(), ID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                user.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": fmt.Sprintf("Failed to get data Class Schedule By ID For presence : %s", ID),
		})

		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}

		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, classSchedule))
}

func (c *MstClassScheduleController) GetByClassID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	classID := ctx.Params("class_id")
	res, err := c.mstClassScheduleService.GetByClassID(ctx.Context(), classID)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get by class id MstClassScheduleController",
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

func (c *MstClassScheduleController) GetByClassAsDate(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
	classID := ctx.Params("class_id")

	var pageble pageable.PageableRequestClassParticipant

	if err := ctx.QueryParser(&pageble); err != nil {
		c.log.Warnf("Failed to parse query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	_, err := utils.ValidateAndPrepareRequest(&pageble)
	if err != nil {
		c.log.Warnf("Failed to query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	pageble.ClassID = classID
	res, err := c.mstClassScheduleService.GetByClassAsDate(ctx.Context(), pageble)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get all MstClassScheduleController",
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

func (c *MstClassScheduleController) GetAllWithCountByLecuturerID(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	var pageble pageable.PageableRequestClassScheduleLecturer
	if err := ctx.QueryParser(&pageble); err != nil {
		c.log.Warnf("Failed to parse query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	_, err := utils.ValidateAndPrepareRequest(&pageble)
	if err != nil {
		c.log.Warnf("Failed to query : %+v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.CreateError(fiber.StatusBadRequest, codeError, err.Error()))
	}

	res, err := c.mstClassScheduleService.GetAllWithCountByLecturerID(ctx.Context(), pageble)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":            codeError,
			"user":                  userClaims.ID,
			"location":              utils.ErrorLocation(),
			"mes/`sage-error-debug": "Failed to Get all Mst Class Schedule By Lectuer ID",
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

func (c *MstClassScheduleController) GetByLecturerIDandActiveAcademicPeriod(ctx *fiber.Ctx) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	classes, err := c.mstClassScheduleService.GetByLecturerIDandActiveAcademicPeriod(ctx.Context())
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to Get academic system distributions",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, classes),
	)
}

// TODO: Helper
// generate
func (c *MstClassScheduleController) wrapGenerate(
	ctx *fiber.Ctx,
	handler func(context.Context, dto.GenerateScheduleRequest) error,
) error {
	codeError := ctx.Locals("code-error").(string)
	userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

	var req dto.GenerateScheduleRequest
	classID := ctx.Params("class_id")
	req.ClassID, _ = utilicems.StringToUuid(classID)

	err := handler(ctx.Context(), req)
	if err != nil {
		utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": "Failed to generate schedule",
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, ""),
	)
}

// statusCode
func errorStatusController(
	ctx *fiber.Ctx,
	log *logrus.Logger,
	err error,
	codeError,
	msgSucces,
	msgErrorForLog string,
	data any,
) error {
	if err != nil {

		userClaims := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
		utils.CreateCaptureAndLogFileError(log, err, map[string]any{
			"code-error":          codeError,
			"user":                userClaims.ID,
			"location":            utils.ErrorLocation(),
			"message-error-debug": msgErrorForLog,
		})
		statusCode := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			statusCode = fiberErr.Code
		}
		return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(
		dto.CreateSuccess(fiber.StatusOK, msgSucces, data),
	)
}
