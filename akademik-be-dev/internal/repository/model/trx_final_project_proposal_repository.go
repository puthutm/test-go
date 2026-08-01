package repositorymodel

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/internal/repository/cached"
	"unsia.ac.id/akademic_be/pkg/auth"
)

type TrxFinalProjectProposalRepository struct {
	log *logrus.Logger
	repository.Repository[model.TrxFinalProjectProposal]
	cacheRepository cached.CacheRepository
}

func NewTrxFinalProjectProposalRepository(
	log *logrus.Logger,
	cacheRepository cached.CacheRepository,
) *TrxFinalProjectProposalRepository {
	return &TrxFinalProjectProposalRepository{
		log:             log,
		cacheRepository: cacheRepository,
	}
}

/* Create */
func (r *TrxFinalProjectProposalRepository) Create(db *gorm.DB, ctx context.Context, data *model.TrxFinalProjectProposal) error {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	data.StudentID, _ = uuid.Parse(user.ID)
	return db.WithContext(ctx).Create(data).Error
}

func (r *TrxFinalProjectProposalRepository) AsignAcademicSupervisor(db *gorm.DB, ctx context.Context, req dto.TrxFinalProjectProposalAssignAcademicSupervisorRequest) error {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	item := map[string]interface{}{
		"id":                        req.ID,
		"final_project_proposal_id": req.FinalProjectProposalID,
		"user_id":                   user.ID,
		"lecturer_id":               req.LecturerID,
		"assign_date":               time.Now().UnixMilli(),
	}
	return db.WithContext(ctx).Table("trx_final_project_proposal_lecturers").Create(item).Error
}

/* Read */
func (r *TrxFinalProjectProposalRepository) GetAllByStudent(db *gorm.DB, ctx context.Context) (T []model.TrxFinalProjectProposal, err error) {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	err = db.WithContext(ctx).Where("student_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", user.ID).Find(&T).Error
	return T, err
}

func (r *TrxFinalProjectProposalRepository) GetByID(db *gorm.DB, ID string, data *model.TrxFinalProjectProposal) error {
	return db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).First(data).Error
}

func (r *TrxFinalProjectProposalRepository) GetByIDGroupByStudent(db *gorm.DB, ID string) (T []model.TrxFinalProjectProposal, err error) {
	err = db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).Find(&T).Error
	return T, err
}

func (r *TrxFinalProjectProposalRepository) GetProposalStudentByUser(db *gorm.DB, ID string) (T []model.TrxFinalProjectProposal, err error) {
	err = db.Where("student_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).Find(&T).Error
	return T, err
}

func (r *TrxFinalProjectProposalRepository) GetByStudentIDandStudyProgramID(db *gorm.DB, StudentID string, StudyProgramID string, data *model.TrxFinalProjectProposal) error {
	return db.Where("student_id = ? AND study_program_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", StudentID, StudyProgramID).First(data).Error
}

func (r *TrxFinalProjectProposalRepository) GetAllWithCountProgramHeadIDORLecturer(db *gorm.DB, ctx context.Context, pg pageable.PageableRequestFinalProjectProposal) (T []model.TrxFinalProjectProposal, count int64, err error) {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	query := db.WithContext(ctx).Model(&model.TrxFinalProjectProposal{}).
		Where("student_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", user.ID)

	if pg.Status != nil && *pg.Status != "" {
		query = query.Where("status = ?", *pg.Status)
	}

	search := pg.GetDefaultSearch()
	if search != "" {
		query = query.Where("title_id ILIKE ? OR topic ILIKE ?", "%"+search+"%", "%"+search+"%")
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

	err = query.Offset(offset).Limit(limit).Find(&T).Error
	return T, count, err
}

/* GetMentor */
func (r *TrxFinalProjectProposalRepository) GetMentorLecturerByID(db *gorm.DB, ID string) (T []model.TrxFinalProjectProposalMentorLecturer, err error) {
	err = db.Table("trx_final_project_proposal_lecturers").
		Where("final_project_proposal_id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).
		Scan(&T).Error
	return T, err
}

/* Update */
func (r *TrxFinalProjectProposalRepository) UpdateStatusByID(db *gorm.DB, ctx context.Context, req dto.TrxFinalProjectProposalUpdateStatusRequest) error {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	return db.WithContext(ctx).Model(&model.TrxFinalProjectProposal{}).
		Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"status":            req.Status,
			"confirmation_date": time.Now().UnixMilli(),
			"feedback":          req.Feedback,
			"updated_by":        user.ID,
		}).Error
}

/* Delete */
