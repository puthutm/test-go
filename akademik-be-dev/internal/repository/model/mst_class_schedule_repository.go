package repositorymodel

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	msg "unsia.ac.id/akademic_be/internal/config/message"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	"unsia.ac.id/akademic_be/internal/service/command"
	"unsia.ac.id/akademic_be/pkg/auth"
)

type MstClassScheduleRepository struct {
	repository.Repository[model.MstClassSchedule]
	log   *logrus.Logger
	cache cached.CacheRepository
}

func NewMstClassScheduleRepository(
	log *logrus.Logger,
	cache cached.CacheRepository,
) *MstClassScheduleRepository {
	return &MstClassScheduleRepository{
		log:   log,
		cache: cache,
	}
}

/* Create */
func (r *MstClassScheduleRepository) Create(
	db *gorm.DB,
	req *model.MstClassSchedule,
) error {
	err := db.Create(req).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "create class schedule",
		}).Error(msg.ErrCreate.Error())
		return err
	}
	return nil
}

/* Read */
func (r *MstClassScheduleRepository) GetByClassID(
	db *gorm.DB,
	ClassID string,
) (datas []model.MstClassSchedule, err error) {
	err = db.Where("class_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ClassID).Find(&datas).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get by id class schedule",
		}).Error(msg.ErrCreate.Error())
		return datas, err
	}
	return datas, err
}

func (r *MstClassScheduleRepository) GetByID(db *gorm.DB, ID string, data *model.MstClassSchedule) error {
	return db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).First(data).Error
}

func (r *MstClassScheduleRepository) GetByIDForPresence(db *gorm.DB, ID string, data *model.MstClassScheduleForClassSessionPresence) error {
	return db.Table("mst_class_schedules").Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).Scan(data).Error
}

func (r *MstClassScheduleRepository) GetByDayTime(
	db *gorm.DB,
	req command.MstClassScheduleGetByCommand,
	data *model.MstClassSchedule,
) error {
	err := db.Where("class_id = ? AND day_name = ? AND start_time = ? AND end_time = ? AND (deleted_at IS NULL OR deleted_at = 0)",
		req.ClassID, req.DayName, req.StartTime, req.EndTime).First(data).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get by day time class schedule",
		}).Error(msg.ErrRead.Error())
		return err
	}
	return nil
}

func (r *MstClassScheduleRepository) GetAllWithCountByLecturerID(
	db *gorm.DB, ctx context.Context, pg pageable.PageableRequestClassScheduleLecturer,
) (T []model.MstClass, count int64, err error) {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)

	query := db.WithContext(ctx).Table("mst_classes").
		Joins("JOIN mst_class_lecturers ON mst_class_lecturers.class_id = mst_classes.id").
		Where("mst_class_lecturers.lecturer_user_id = ? AND (mst_classes.deleted_at IS NULL OR mst_classes.deleted_at = 0)", user.ID)

	search := pg.GetDefaultSearch()
	if search != "" {
		query = query.Where("mst_classes.name ILIKE ? OR mst_classes.code ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
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

	err = query.Offset(offset).Limit(limit).Scan(&T).Error
	return T, count, err
}

func (r *MstClassScheduleRepository) GetByLecturerIDandActiveAcademicPeriod(db *gorm.DB, ctx context.Context) (T []model.MstClassSchedule, err error) {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	err = db.WithContext(ctx).Table("mst_class_schedules").
		Joins("JOIN mst_class_lecturers ON mst_class_lecturers.class_id = mst_class_schedules.class_id").
		Where("mst_class_lecturers.lecturer_user_id = ? AND (mst_class_schedules.deleted_at IS NULL OR mst_class_schedules.deleted_at = 0)", user.ID).
		Scan(&T).Error
	return T, err
}

/* Update */
func (r *MstClassScheduleRepository) UpdateByDayTime(
	db *gorm.DB,
	req command.MstClassScheduleUpdateCommand,
) error {
	err := db.Model(&model.MstClassSchedule{}).
		Where("class_id = ? AND day_name = ? AND start_time = ? AND end_time = ?",
			req.ClassID, req.DayNameOld, req.StartTimeOld, req.EndTimeOld).
		Updates(map[string]interface{}{
			"day_name":        req.DayName,
			"start_time":      req.StartTime,
			"end_time":        req.EndTime,
			"date":            req.Date,
			"type_of_meeting": req.TypeOfMeeting,
		}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update class schedule",
		}).Error(msg.ErrUpdate.Error())
		return err
	}
	return nil
}

func (r *MstClassScheduleRepository) UpdateByIDForLecturer(
	db *gorm.DB,
	data *model.MstClassSchedule,
) error {
	err := db.Model(&model.MstClassSchedule{}).
		Where("id = ?", data.ID).
		Updates(map[string]interface{}{
			"material_attachment_file_path": data.MaterialAttachmentFilePath,
			"attendance_document_file_path":  data.AttendanceDocumentFilePath,
			"journal_document_file_path":     data.JournalDocumentFilePath,
			"material_plan":                 data.MaterialPlan,
			"material_realization":          data.MaterialRealization,
		}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "update class schedule",
		}).Error(msg.ErrUpdate.Error())
		return err
	}
	return nil
}

/* Delete */
func (r *MstClassScheduleRepository) DeleteByScheduleTemplate(
	db *gorm.DB,
	scheduleTemplateID string,
) error {
	err := db.Where("schedule_template_id = ?", scheduleTemplateID).Delete(&model.MstClassSchedule{}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "delete class schedule by schedule_template_id",
		}).Error(msg.ErrDelete.Error())
		return err
	}
	return nil
}

func (r *MstClassScheduleRepository) DeleteByDayTime(
	db *gorm.DB,
	req command.MstClassScheduleGetByCommand,
) error {
	err := db.Where("class_id = ? AND day_name = ? AND start_time = ? AND end_time = ?",
		req.ClassID, req.DayName, req.StartTime, req.EndTime).Delete(&model.MstClassSchedule{}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "delete class schedule by date time",
		}).Error(msg.ErrDelete.Error())
		return err
	}
	return nil
}

func (r *MstClassScheduleRepository) DeleteByID(
	db *gorm.DB,
	ID string,
) error {
	err := db.Where("id = ?", ID).Delete(&model.MstClassSchedule{}).Error
	if err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "delete class schedule",
		}).Error(msg.ErrDelete.Error())
		return err
	}
	return nil
}

func (r *MstClassScheduleRepository) GetByClassAsDate(
	db *gorm.DB,
	deleted bool,
	pg pageable.PageableRequestClassParticipant,
) (T []model.MstClassSchedule, count int64, err error) {
	query := db.Model(&model.MstClassSchedule{}).Where("class_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", pg.ClassID)

	if err := query.Count(&count).Error; err != nil {
		r.log.WithFields(logrus.Fields{
			"repository": "get class schedule as of date",
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

func (r *MstClassScheduleRepository) GetByScheduleTemplate(
	db *gorm.DB,
	scheduleTemplateID string,
) (T []model.MstClassSchedule, err error) {
	err = db.Where("schedule_template_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", scheduleTemplateID).Find(&T).Error
	return T, err
}
