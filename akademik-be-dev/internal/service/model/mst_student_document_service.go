package servicemodel

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
	"unsia.ac.id/akademic_be/pkg/utils"
)

type MstStudentDocumentService struct {
	log                          *logrus.Logger
	db                           *gorm.DB
	cache                        cached.CacheRepository
	mstStudentDocumentRepository *repositorymodel.MstStudentDocumentRepository
}

func NewMstStudentDocumentService(
	log *logrus.Logger,
	db *gorm.DB,
	cache cached.CacheRepository,
	mstStudentDocumentRepository *repositorymodel.MstStudentDocumentRepository,
) *MstStudentDocumentService {
	return &MstStudentDocumentService{
		log:   log,
		db:    db,
		cache: cache,
	}
}

// handleCreate
func (s *MstStudentDocumentService) handleCreateOnlyUser(
	ctx context.Context,
	req dto.MstStudentDocumentRequestOnlyUser,
) error {
	tx := s.db.WithContext(ctx)

	data := &model.MstStudentDocument{
		ID: utils.GenerateUUID(),
	}

	err := s.mstStudentDocumentRepository.Create(
		tx,
		data,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "handle create student document only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

func (s *MstStudentDocumentService) handleCreate(
	ctx context.Context,
	req dto.MstStudentDocumentRequest,
) error {
	return nil
}

func (s *MstStudentDocumentService) handleUpdateOnlyUser(
	ctx context.Context,
	req dto.MstStudentDocumentRequestOnlyUser,
) error {
	tx := s.db.WithContext(ctx)

	data := &model.MstStudentDocument{
		ID: utils.GenerateUUID(),
	}

	err := s.mstStudentDocumentRepository.Update(
		tx,
		data,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "handle update  student document only user",
			"request": req,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

func (s *MstStudentDocumentService) handleUpdate(
	ctx context.Context,
	req dto.MstStudentDocumentRequest,
) error {
	return nil
}

// Create
func (s *MstStudentDocumentService) Create(ctx context.Context, req any) error {
	switch r := req.(type) {
	case dto.MstStudentDocumentRequest:
		return s.handleCreate(ctx, r)
	case dto.MstStudentDocumentRequestOnlyUser:
		return s.handleCreateOnlyUser(ctx, r)
	default:
		return msg.ErrFailedType
	}
}

// Update

// Delete
func (s *MstStudentDocumentService) Delete(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	err := s.mstStudentDocumentRepository.DeleteByID(
		tx,
		ID,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "delete student document",
			"id":      ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

// Restore
func (s *MstStudentDocumentService) Restore(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	err := s.mstStudentDocumentRepository.RestoreByID(
		tx,
		ID,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "restore student document",
			"id":      ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

// GetByID
func (s *MstStudentDocumentService) GetByID(ctx context.Context, ID string) error {
	tx := s.db.WithContext(ctx)

	data := new(model.MstStudentDocument)

	err := s.mstStudentDocumentRepository.GetByID(
		tx,
		ID,
		data,
	)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get by id student document",
			"id":      ID,
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}

// GetAllWithCount
func (s *MstStudentDocumentService) GetAllWithCount(ctx context.Context, pageble pageable.PageableRequest) error {
	tx := s.db.WithContext(ctx)

	res, count, err := s.mstStudentDocumentRepository.GetAllWithCount(
		tx,
		true,
		pageble,
	)
	fmt.Println(res, count)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"service": "get all student document",
			"all":     "all data",
		}).Error(err)
		return utils.ErrorSpToErrorFiber(err)
	}

	return nil
}
