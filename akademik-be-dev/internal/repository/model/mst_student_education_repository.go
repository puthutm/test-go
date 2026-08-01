package repositorymodel

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
)

type MstStudentEducationRepository struct {
	repository.Repository[model.MstStudentEducation]
	log *logrus.Logger
}

func NewMstStudentEducationRepository(log *logrus.Logger) *MstStudentEducationRepository {
	return &MstStudentEducationRepository{
		log: log,
	}
}

/* Create */
/* Read */
func (r *MstStudentEducationRepository) GetByStudentID(db *gorm.DB, StudentID string, data *model.MstStudentEducation) error {
	return db.Where("student_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", StudentID).First(data).Error
}

/* Update */
func (r *MstStudentEducationRepository) UpdateByID(db *gorm.DB, req dto.MstStudentEducationRequest) error {
	return db.Model(&model.MstStudentEducation{}).
		Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"student_id":           req.StudentID,
			"institution_name":     req.InstitutionName,
			"school_major":         req.SchoolMajor,
			"nisn":                 req.NISN,
			"national_exam_score":  req.NationalExamScore,
			"certificate_number":   req.CertificateNumber,
			"certificate_filepath": req.CertificateFilepath,
			"transcripts_filepath": req.TranscriptsFilepath,
			"updated_at":           time.Now().UnixMilli(),
		}).Error
}

/* Delete */
