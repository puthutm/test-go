package repositorymodel

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
)

type MstStudentDocumentRepository struct {
	log *logrus.Logger
	repository.Repository[model.MstStudentDocument]
}

func NewMstStudentDocumentRepository(
	log *logrus.Logger,
) *MstStudentDocumentRepository {
	return &MstStudentDocumentRepository{
		log: log,
	}
}

// Create
func (r *MstStudentDocumentRepository) Create(db *gorm.DB, data *model.MstStudentDocument) error {
	err := db.Create(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "create student_document",
			"student_id": data.StudentID,
		}).Error(msg.ErrCreate.Error())
		return err
	}
	return nil
}

// Update
func (r *MstStudentDocumentRepository) Update(db *gorm.DB, data *model.MstStudentDocument) error {
	err := db.Model(&model.MstStudentDocument{}).Where("id = ?", data.ID).Updates(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update student_document",
			"id":         data.ID,
		}).Error(msg.ErrUpdate.Error())
		return err
	}
	return nil
}

// Delete
func (r *MstStudentDocumentRepository) DeleteByID(db *gorm.DB, ID string) error {
	err := db.Where("id = ?", ID).Delete(&model.MstStudentDocument{}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update student_document",
			"id":         ID,
		}).Error(msg.ErrDelete.Error())
		return err
	}
	return nil
}

// Restore
func (r *MstStudentDocumentRepository) RestoreByID(db *gorm.DB, ID string) error {
	err := db.Model(&model.MstStudentDocument{}).Where("id = ?", ID).Update("deleted_at", nil).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update student_document",
			"id":         ID,
		}).Error(msg.ErrRestore.Error())
		return err
	}
	return nil
}

// GetByID
func (r *MstStudentDocumentRepository) GetByID(db *gorm.DB, ID string, data *model.MstStudentDocument) error {
	err := db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).First(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update student_document",
			"id":         ID,
		}).Error(msg.ErrRestore.Error())
		return err
	}
	return nil
}

// GetAllWithCount
func (r *MstStudentDocumentRepository) GetAllWithCount(
	db *gorm.DB,
	deleted bool,
	pageble pageable.PageableRequest,
) (T []model.MstStudentDocument, count int64, err error) {
	query := db.Model(&model.MstStudentDocument{})
	if deleted {
		query = query.Where("deleted_at IS NOT NULL AND deleted_at > 0")
	} else {
		query = query.Where("deleted_at IS NULL OR deleted_at = 0")
	}

	if err := query.Count(&count).Error; err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get all student_document",
		}).Error(msg.ErrMultipleRead.Error())
		return T, count, err
	}

	page := pageble.Page
	limit := pageble.Limit
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	err = query.Offset(offset).Limit(limit).Find(&T).Error
	return T, count, err
}
