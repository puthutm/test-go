package repositorymodel

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
)

type MstStudentFamilyRepository struct {
	repository.Repository[model.MstStudentFamily]
	log *logrus.Logger
}

func NewMstStudentFamilyRepository(log *logrus.Logger) *MstStudentFamilyRepository {
	return &MstStudentFamilyRepository{
		log: log,
	}
}

/* Create */
/* Read */
func (r *MstStudentFamilyRepository) GetByStudentID(db *gorm.DB, StudentID string, parentType string, data *model.MstStudentFamily) error {
	return db.Where("student_id = ? AND status_kinship = ? AND (deleted_at IS NULL OR deleted_at = 0)", StudentID, parentType).First(data).Error
}

/* Update */
func (r *MstStudentFamilyRepository) UpdateByStudentID(db *gorm.DB, req dto.MstStudentFamilyRequest) error {
	return db.Model(&model.MstStudentFamily{}).
		Where("student_id = ? AND status_kinship = ?", req.StudentID, req.Type).
		Updates(map[string]interface{}{
			"name":                 req.Name,
			"nik":                  req.NIK,
			"educational_level_id": req.EducationLevelID,
			"birth_place_id":       req.BirthPlaceID,
			"birth_date":           req.BirthDate,
			"job_id":               req.JobID,
			"income":               req.Income,
			"email":                req.Email,
			"phone":                req.Phone,
			"phone2":               req.Phone2,
			"address":              req.Address,
			"life_status":          req.LifeStatus,
			"updated_at":           time.Now().UnixMilli(),
		}).Error
}

/* Delete */
