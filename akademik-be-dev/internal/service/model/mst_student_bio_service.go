package servicemodel

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/converter"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/internal/service"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstStudentBioService struct {
	log                     *logrus.Logger
	db                      *gorm.DB
	cache                   cached.CacheRepository
	mstStudentBioRepository *repositorymodel.MstStudentBioRepository
	storageService          *service.StorageService
}

func NewMstStudentBioService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstStudentBioRepository *repositorymodel.MstStudentBioRepository,

	storageService *service.StorageService,
) *MstStudentBioService {
	return &MstStudentBioService{
		log:                     log,
		db:                      db,
		cache:                   cache,
		mstStudentBioRepository: mstStudentBioRepository,
		storageService:          storageService,
	}
}

// Create

// TODO: Handle Update Biodata
// TODO: Biodata general
func (s *MstStudentBioService) handleUpdateBioGeneral(
	ctx context.Context,
	req dto.MstStudentBioUpdate,
) error {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentBio)

	data = converter.ConvertMstStudentBioUpdateToModelPointer(req)

	err := s.mstStudentBioRepository.Update(
		tx,
		data,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update student update ",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstStudentBioService) handleUpdateBioGeneralOnlyUser(
	ctx context.Context,
	req dto.MstStudentBioUpdateOnlyUser,
) error {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentBio)

	err := s.mstStudentBioRepository.CheckByUserID(tx, req.UserID.String(), data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "check student update only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	req.ID = data.ID

	err = s.mstStudentBioRepository.GetByID(tx, data.ID.String(), data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "check student update only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	converter.ConvertMstStudentBioUpdateOnlyUserToModelPointer(req, data)

	err = s.mstStudentBioRepository.Update(
		tx,
		data,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update student update only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Biodata Completenes
func (s *MstStudentBioService) handleUpdateBioCompletenesOnlyUser(
	ctx context.Context,
	req dto.MstStudentBioUpdateCompletenessOnlyUser,
) error {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentBio)
	err := s.mstStudentBioRepository.CheckByUserID(tx, req.UserID.String(), data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "check student update only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	req.ID = data.ID

	err = s.mstStudentBioRepository.GetCompletenessByID(tx, data.ID.String(), data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "check student update only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	converter.ConvertMstStudentBioUpdateCompletenessOnlyUserToModelPointer(req, data)

	if req.SignaturePathFile != nil {
		pathSignature, err := s.storageService.UploadFileV2(
			ctx, req.SignaturePathFile,
			"student/biodata/completenes",
			[]string{".jpeg", ".jpg", ".png", ".pdf"}, 1,
		)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service":     "update student completenes only user",
				"use service": "storage service upload",
				"request":     req,
			}).Error(err)
			return fiber.NewError(fiber.StatusInternalServerError, msg.ErrUpdate.Error())
		}
		data.SignaturePathFile = &pathSignature
	}

	err = s.mstStudentBioRepository.UpdateCompleteness(
		tx,
		data,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update student completenes only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Biodata Information
func (s *MstStudentBioService) handleUpdateBioInformationOnlyUser(
	ctx context.Context,
	req dto.MstStudentBioUpdateInformationOnlyUser,
) error {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentBio)
	err := s.mstStudentBioRepository.CheckByUserID(tx, req.UserID.String(), data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "check student update only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	req.ID = data.ID

	err = s.mstStudentBioRepository.GetInformationByID(tx, data.ID.String(), data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "check student update only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	converter.ConvertMstStudentBioUpdateInformationOnlyUserToModelPointer(req, data)

	err = s.mstStudentBioRepository.UpdateInformation(
		tx,
		data,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update student information only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Biodata Document
func (s *MstStudentBioService) handleUpdateBioDocumentOnlyUser(
	ctx context.Context,
	req dto.MstStudentBioUpdateDocumentOnlyUser,
) error {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentBio)
	err := s.mstStudentBioRepository.CheckByUserID(tx, req.UserID.String(), data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "check student update only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	req.ID = data.ID

	err = s.mstStudentBioRepository.GetDocumentByID(tx, data.ID.String(), data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "check student update only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	converter.ConvertMstStudentBioUpdateDocumentOnlyUserToModelPointer(req, data)

	if req.BPJSEmploymentFilePath != nil {
		path, err := s.storageService.UploadFileV2(
			ctx, req.BPJSEmploymentFilePath,
			"student/biodata/document",
			[]string{".jpeg", ".jpg", ".png", ".pdf"}, 1,
		)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service":     "update student document only user",
				"use service": "storage service upload",
				"request":     req,
			}).Error(err)
			return fiber.NewError(fiber.StatusInternalServerError, msg.ErrUpdate.Error())
		}
		data.BPJSEmploymentFilepath = &path
	}
	if req.BPJSHealthcareFilePath != nil {
		path, err := s.storageService.UploadFileV2(
			ctx, req.BPJSHealthcareFilePath,
			"student/biodata/document",
			[]string{".jpeg", ".jpg", ".png", ".pdf"}, 1,
		)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service":     "update student document only user",
				"use service": "storage service upload",
				"request":     req,
			}).Error(err)
			return fiber.NewError(fiber.StatusInternalServerError, msg.ErrUpdate.Error())
		}
		data.BPJSHealthcareFilepath = &path
	}
	if req.NPWPFilePath != nil {
		path, err := s.storageService.UploadFileV2(
			ctx, req.NPWPFilePath,
			"student/biodata/document",
			[]string{".jpeg", ".jpg", ".png", ".pdf"}, 1,
		)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service":     "update student document only user",
				"use service": "storage service upload",
				"request":     req,
			}).Error(err)
			return fiber.NewError(fiber.StatusInternalServerError, msg.ErrUpdate.Error())
		}
		data.NPWPFilepath = &path
	}

	err = s.mstStudentBioRepository.UpdateDocuments(
		tx,
		data,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update student document only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

// TODO: Biodata Account Bank
func (s *MstStudentBioService) handleUpdateBioBankAccountOnlyUser(
	ctx context.Context,
	req dto.MstStudentBioUpdateBankAccountOnlyUser,
) error {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentBio)
	err := s.mstStudentBioRepository.CheckByUserID(tx, req.UserID.String(), data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "check student update only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	req.ID = data.ID

	err = s.mstStudentBioRepository.GetBankAccountByID(tx, data.ID.String(), data)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "check student update only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	converter.ConvertMstStudentBioUpdateBankAccountOnlyUserToModelPointer(req, data)

	if req.AccountFilePath != nil {
		path, err := s.storageService.UploadFileV2(
			ctx, req.AccountFilePath,
			"student/biodata/completenes",
			[]string{".jpeg", ".jpg", ".png", ".pdf"}, 1,
		)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service":     "update student bank account only user",
				"use service": "storage service upload",
				"request":     req,
			}).Error(err)
			return fiber.NewError(fiber.StatusInternalServerError, msg.ErrUpdate.Error())
		}
		data.AccountFilepath = &path
	}

	err = s.mstStudentBioRepository.UpdateBankAccount(
		tx,
		data,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "update student bank account only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}
	return nil
}

func (s *MstStudentBioService) Update(ctx context.Context, req any) error {
	switch r := req.(type) {

	// TODO: General
	case dto.MstStudentBioUpdate:
		return s.handleUpdateBioGeneral(ctx, r)
	case dto.MstStudentBioUpdateOnlyUser:
		return s.handleUpdateBioGeneralOnlyUser(ctx, r)

		// TODO: Completeness
	case dto.MstStudentBioUpdateCompletenessOnlyUser:
		return s.handleUpdateBioCompletenesOnlyUser(ctx, r)

		// TODO: Information
	case dto.MstStudentBioUpdateInformationOnlyUser:
		return s.handleUpdateBioInformationOnlyUser(ctx, r)

		// TODO: Document
	case dto.MstStudentBioUpdateDocumentOnlyUser:
		return s.handleUpdateBioDocumentOnlyUser(ctx, r)

		// TODO: Bank Account
	case dto.MstStudentBioUpdateBankAccountOnlyUser:
		return s.handleUpdateBioBankAccountOnlyUser(ctx, r)

		// default
	default:
		return msg.ErrFailedType
	}
}

// Delete
// Restore

// TODO: GetGeneralByID -> userID != "" -> validasi
func (s *MstStudentBioService) GetGeneralByID(ctx context.Context, ID, userID string) (*dto.MstStudentBioGeneralResponse, error) {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentBio)

	if userID != "" {
		err := s.mstStudentBioRepository.CheckByUserID(tx, userID, data)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service": "check student only user",
				"user-id": userID,
			}).Error(err)
			return nil, utils.ErrorSpToErrorFiber(err)
		}

		ID = data.ID.String()

		// if data.ID.String() != ID {
		// 	return nil, fiber.NewError(fiber.StatusForbidden, msg.ErrUnauthorized.Error())
		// }
	}

	err := s.mstStudentBioRepository.GetByID(
		tx,
		ID,
		data,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id student general",
			"id":      ID,
		}).Error(err)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := converter.ConvertMstStudentBioGeneralToResponse(data)

	return res, nil
}

func (s *MstStudentBioService) GetinformationByID(ctx context.Context, ID, userID string) (*dto.MstStudentBioInfomationResponse, error) {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentBio)

	if userID != "" {
		err := s.mstStudentBioRepository.CheckByUserID(tx, userID, data)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service": "check student only user",
				"user-id": userID,
			}).Error(err)
			return nil, utils.ErrorSpToErrorFiber(err)
		}

		ID = data.ID.String()
		// if data.ID.String() != ID {
		// 	return nil, fiber.NewError(fiber.StatusForbidden, msg.ErrUnauthorized.Error())
		// }
	}

	err := s.mstStudentBioRepository.GetInformationByID(
		tx,
		ID,
		data,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id student information",
			"id":      ID,
		}).Error(err)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := converter.ConvertMstStudentBioInformationToResponse(data)

	return res, nil
}

func (s *MstStudentBioService) GetCompletenesByID(ctx context.Context, ID, userID string) (*dto.MstStudentBioCompletenesResponse, error) {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentBio)

	if userID != "" {
		err := s.mstStudentBioRepository.CheckByUserID(tx, userID, data)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service": "check student only user",
				"user-id": userID,
			}).Error(err)
			return nil, utils.ErrorSpToErrorFiber(err)
		}

		ID = data.ID.String()
		// if data.ID.String() != ID {
		// 	return nil, fiber.NewError(fiber.StatusForbidden, msg.ErrUnauthorized.Error())
		// }
	}

	err := s.mstStudentBioRepository.GetCompletenessByID(
		tx,
		ID,
		data,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id student completenes",
			"id":      ID,
		}).Error(err)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := converter.ConvertMstStudentBioCompletenesToResponse(data)

	return res, nil
}

func (s *MstStudentBioService) GetDocumentByID(ctx context.Context, ID, userID string) (*dto.MstStudentBioDocumentResponse, error) {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentBio)

	if userID != "" {
		err := s.mstStudentBioRepository.CheckByUserID(tx, userID, data)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service": "check student only user",
				"user-id": userID,
			}).Error(err)
			return nil, utils.ErrorSpToErrorFiber(err)
		}

		ID = data.ID.String()
		// if data.ID.String() != ID {
		// 	return nil, fiber.NewError(fiber.StatusForbidden, msg.ErrUnauthorized.Error())
		// }
	}

	err := s.mstStudentBioRepository.GetDocumentByID(
		tx,
		ID,
		data,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id student document",
			"id":      ID,
		}).Error(err)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := converter.ConvertMstStudentBioToDocumentResponse(data)

	return res, nil
}

func (s *MstStudentBioService) GetBankAccountByID(ctx context.Context, ID, userID string) (*dto.MstStudentBioBankAccountResponse, error) {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentBio)

	if userID != "" {
		err := s.mstStudentBioRepository.CheckByUserID(tx, userID, data)
		if err != nil {
			s.log.WithFields(logrus.Fields{
				"service": "check student only user",
				"user-id": userID,
			}).Error(err)
			return nil, utils.ErrorSpToErrorFiber(err)
		}

		ID = data.ID.String()

		// if data.ID.String() != ID {
		// 	return nil, fiber.NewError(fiber.StatusForbidden, msg.ErrUnauthorized.Error())
		// }
	}

	err := s.mstStudentBioRepository.GetBankAccountByID(
		tx,
		ID,
		data,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id student bank account",
			"id":      ID,
		}).Error(err)
		return nil, utils.ErrorSpToErrorFiber(err)
	}

	res := converter.ConvertMstStudentBioToBankAccountResponse(data)

	return res, nil
}

// TODO: GetAll
