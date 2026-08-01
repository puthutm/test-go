package repositorymodel

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
)

type PortalStudentRepository interface {
	GetAllBatches(db *gorm.DB) ([]model.PortalStudentBatch, error)
	CreateStudent(db *gorm.DB, req dto.PortalStudentCreateRequest) error
	GetStudentListWithCount(db *gorm.DB, pg pageable.PageableRequestClass) ([]model.PortalStudentList, int64, error)
}

type portalStudentRepository struct {
	log *logrus.Logger
}

func NewPortalStudentRepository(log *logrus.Logger) PortalStudentRepository {
	return &portalStudentRepository{
		log: log,
	}
}

func (r *portalStudentRepository) GetAllBatches(db *gorm.DB) ([]model.PortalStudentBatch, error) {
	var batches []model.PortalStudentBatch
	err := db.Table("mst_portal_batches").Where("deleted_at IS NULL OR deleted_at = 0").Find(&batches).Error
	return batches, err
}

func (r *portalStudentRepository) CreateStudent(db *gorm.DB, req dto.PortalStudentCreateRequest) error {
	record := map[string]interface{}{
		"id":              req.ID,
		"batch_detail_id": req.BatchDetailID,
		"nik":             req.NIK,
		"name":            req.Name,
		"email":           req.Email,
		"password":        req.Password,
		"phone":           req.Phone,
		"created_at":      time.Now().UnixMilli(),
	}
	return db.Table("mst_student_bios").Create(record).Error
}

func (r *portalStudentRepository) GetStudentListWithCount(db *gorm.DB, pg pageable.PageableRequestClass) ([]model.PortalStudentList, int64, error) {
	var students []model.PortalStudentList
	var count int64

	studyProgramId := pg.GetDefaultStudyProgramId()
	search := pg.GetDefaultSearch()

	query := db.Table("mst_student_bios").Where("deleted_at IS NULL OR deleted_at = 0")

	if studyProgramId != nil && *studyProgramId != "" {
		query = query.Where("study_program_id = ?", *studyProgramId)
	}

	if search != "" {
		query = query.Where("name ILIKE ? OR nik ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
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

	if err := query.Offset(offset).Limit(limit).Scan(&students).Error; err != nil {
		return nil, 0, err
	}

	return students, count, nil
}
