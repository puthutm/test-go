package repositorymodel

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
)

type MstStudentDomicileRepository struct {
	repository.Repository[model.MstStudentDomicile]
	log *logrus.Logger
}

func NewMstStudentDomicileRepository(log *logrus.Logger) *MstStudentDomicileRepository {
	return &MstStudentDomicileRepository{
		log: log,
	}
}

/* Create */
/* Read */
func (r *MstStudentDomicileRepository) GetByStudentID(db *gorm.DB, StudentID string, data *model.MstStudentDomicile) error {
	return db.Where("student_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", StudentID).First(data).Error
}

/* Update */
func (r *MstStudentDomicileRepository) UpdateByStudentID(db *gorm.DB, req dto.MstStudentDomicileRequest) error {
	return db.Model(&model.MstStudentDomicile{}).
		Where("student_id = ?", req.StudentID).
		Updates(map[string]interface{}{
			"domicile_id": req.ID,
			"country_id":  req.CountryID,
			"province_id": req.ProvinceID,
			"city_id":     req.CityID,
			"district_id": req.DistrictID,
			"village_id":  req.VillageID,
			"rt":          req.RT,
			"rw":          req.RW,
			"address":     req.Address,
			"postal_code": req.PostalCode,
			"distance":    req.Distance,
			"updated_at":  time.Now().UnixMilli(),
		}).Error
}

/* Delete */
