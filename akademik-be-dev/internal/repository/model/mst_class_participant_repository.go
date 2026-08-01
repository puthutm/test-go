package repositorymodel

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
)

type MstClassParticipantRepository struct {
	repository.Repository[model.MstClassParticipant]
	log   *logrus.Logger
	cache cached.CacheRepository
}

func NewMstClassParticipantRepository(log *logrus.Logger, cache cached.CacheRepository) *MstClassParticipantRepository {
	return &MstClassParticipantRepository{
		log:   log,
		cache: cache,
	}
}

/* Create */
func (r *MstClassParticipantRepository) Create(db *gorm.DB, data *model.MstClassParticipant) error {
	err := db.Create(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "create class participant",
			"student_id": data.StudentID,
		}).Error(msg.ErrCreate.Error())
		return err
	}
	return nil
}

/* Read */
func (r *MstClassParticipantRepository) GetAllWithCount(
	db *gorm.DB,
	deleted bool,
	pg pageable.PageableRequestClassParticipant,
) (T []model.MstClassParticipant, count int64, err error) {
	query := db.Model(&model.MstClassParticipant{}).Where("class_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", pg.ClassID)

	search := pg.GetDefaultSearch()
	if search != "" {
		query = query.Where("student_id IN (SELECT id FROM mst_student_bios WHERE name ILIKE ? OR nik ILIKE ?)", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get all class participant",
		}).Error(msg.ErrMultipleRead.Error())
		return T, count, err
	}

	page := pg.GetDefaultPage()
	limit := pg.GetDefaultLimit()
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

func (r *MstClassParticipantRepository) GetByID(db *gorm.DB, ID string, data *model.MstClassParticipant) error {
	return db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).First(data).Error
}

func (r *MstClassParticipantRepository) GetAllWithCountByClassIDForLecturer(
	db *gorm.DB,
	deleted bool,
	pg pageable.PageableRequestClassParticipant,
) (T []model.MstClassParticipant, count int64, err error) {
	return r.GetAllWithCount(db, deleted, pg)
}

/* Delete */
func (r *MstClassParticipantRepository) DeleteByID(db *gorm.DB, ID string) error {
	return db.Where("id = ?", ID).Delete(&model.MstClassParticipant{}).Error
}
