package repositorymodel

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
)

type MstClassScheduleTemplateRepository struct {
	repository.Repository[model.MstClassScheduleTemplate]
	log   *logrus.Logger
	cache cached.CacheRepository
}

func NewMstClassScheduleTemplateRepository(
	log *logrus.Logger,
	cache cached.CacheRepository,
) *MstClassScheduleTemplateRepository {
	return &MstClassScheduleTemplateRepository{
		log:   log,
		cache: cache,
	}
}

/* Create */

func (r *MstClassScheduleTemplateRepository) Create(db *gorm.DB, req *model.MstClassScheduleTemplate) error {
	err := db.Create(req).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "create class schedule template",
			"id":         req.ID,
		}).Error(err.Error())
		return err
	}

	return nil
}

/* Read */
func (r *MstClassScheduleTemplateRepository) GetByID(db *gorm.DB, id string, data *model.MstClassScheduleTemplate) error {
	err := db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", id).First(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get class schedule template by id",
			"id":         id,
		}).Error(msg.ErrRead.Error())
		return err
	}

	return nil
}

func (r *MstClassScheduleTemplateRepository) GetByClassID(db *gorm.DB, classID string, templteData *model.MstClassScheduleTemplate) error {
	err := db.Where("class_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", classID).First(templteData).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get class schedule templates by class_id",
			"class_id":   classID,
		}).Error(err.Error())
		return err
	}
	return nil
}

/* Update */
func (r *MstClassScheduleTemplateRepository) UpdateByID(db *gorm.DB, req *model.MstClassScheduleTemplate) error {
	err := db.Model(&model.MstClassScheduleTemplate{}).
		Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"class_id":        req.ClassID,
			"day_name":        req.DayName,
			"start_time":      req.StartTime,
			"end_time":        req.EndTime,
			"type_of_meeting": req.TypeOfMeeting,
			"updated_at":      req.UpdatedAt,
			"updated_by":      req.UpdatedBy,
		}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update class schedule template by id",
			"id":         req.ID,
		}).Error(msg.ErrUpdate.Error())
		return err
	}

	return nil
}

/* Delete */
func (r *MstClassScheduleTemplateRepository) DeleteByID(db *gorm.DB, id string) error {
	err := db.Where("id = ?", id).Delete(&model.MstClassScheduleTemplate{}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "delete class schedule template by id",
			"id":         id,
		}).Error(msg.ErrDelete.Error())
		return err
	}

	return nil
}
