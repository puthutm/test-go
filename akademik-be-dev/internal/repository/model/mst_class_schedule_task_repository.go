package repositorymodel

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
)

type MstClassScheduleTaskRepository struct {
	repository.Repository[model.MstClassScheduleTask]
	log             *logrus.Logger
	cacheRepository cached.CacheRepository
}

func NewMstClassScheduleTaskRepository(log *logrus.Logger, cacheRepository cached.CacheRepository) *MstClassScheduleTaskRepository {
	return &MstClassScheduleTaskRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

/* Create */
func (r *MstClassScheduleTaskRepository) Create(db *gorm.DB, req dto.MstClassScheduleTaskRequest) error {
	record := map[string]interface{}{
		"id":                                 req.ID,
		"schedule_id":                        req.ScheduleID,
		"title":                              req.Title,
		"description":                        req.Description,
		"is_gradeable":                       req.IsGradeable,
		"is_use_deadline":                    req.IsUseDeadline,
		"deadline_of_assignment_submission": req.DeadlineOfAssignmentSubmission,
		"time_to_open":                       req.TimeToOpen,
		"retake":                             req.Retake,
		"created_at":                         time.Now().UnixMilli(),
	}
	return db.Table("mst_class_schedule_tasks").Create(record).Error
}

/* Read */
func (r *MstClassScheduleTaskRepository) GetAll(db *gorm.DB, ClassID string) (T []model.MstClassScheduleTask, err error) {
	err = db.Table("mst_class_schedule_tasks").
		Joins("JOIN mst_class_schedules ON mst_class_schedules.id = mst_class_schedule_tasks.schedule_id").
		Where("mst_class_schedules.class_id = ? AND (mst_class_schedule_tasks.deleted_at IS NULL OR mst_class_schedule_tasks.deleted_at = 0)", ClassID).
		Scan(&T).Error
	return T, err
}

func (r *MstClassScheduleTaskRepository) GetByID(db *gorm.DB, ID string, data *model.MstClassScheduleTask) error {
	return db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).First(data).Error
}

/* Update */
/* Delete */
