package repositorymodel

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/dto/pageable"
	"unsia.ac.id/akademic_be/internal/model"
	"unsia.ac.id/akademic_be/internal/repository"
	"unsia.ac.id/akademic_be/pkg/auth"
)

type MstSKSLimitRepository struct {
	repository.Repository[model.MstSKSLimit]
	log *logrus.Logger
}

func NewMstSKSLimitRepository(log *logrus.Logger) *MstSKSLimitRepository {
	return &MstSKSLimitRepository{
		log: log,
	}
}

/* Create */
func (r *MstSKSLimitRepository) Create(db *gorm.DB, ctx context.Context, req dto.MstSKSLimitRequest) error {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)

	item := map[string]interface{}{
		"id":         req.ID,
		"ips_min":    req.IPSMin,
		"ips_max":    req.IPSMax,
		"sks_limit":  req.SKSLimit,
		"created_at": time.Now().UnixMilli(),
		"created_by": user.ID,
	}
	return db.WithContext(ctx).Table("mst_sks_limits").Create(item).Error
}

/* Read */
func (r *MstSKSLimitRepository) GetAllWithCount(db *gorm.DB, deleted bool, pg pageable.PageableRequestInterface) (T []model.MstSKSLimit, count int64, err error) {
	query := db.Model(&model.MstSKSLimit{})
	if deleted {
		query = query.Where("deleted_at IS NOT NULL AND deleted_at > 0")
	} else {
		query = query.Where("deleted_at IS NULL OR deleted_at = 0")
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

func (r *MstSKSLimitRepository) GetByID(db *gorm.DB, ID string, data *model.MstSKSLimit) error {
	return db.Where("id = ? AND (deleted_at IS NULL OR deleted_at = 0)", ID).First(data).Error
}

/* Update */
func (r *MstSKSLimitRepository) UpdateByID(db *gorm.DB, ctx context.Context, req dto.MstSKSLimitRequest) error {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)

	return db.WithContext(ctx).Table("mst_sks_limits").
		Where("id = ?", req.ID).
		Updates(map[string]interface{}{
			"ips_min":    req.IPSMin,
			"ips_max":    req.IPSMax,
			"sks_limit":  req.SKSLimit,
			"updated_at": time.Now().UnixMilli(),
			"updated_by": user.ID,
		}).Error
}

func (r *MstSKSLimitRepository) RestoreByID(db *gorm.DB, ID string) error {
	return db.Table("mst_sks_limits").Where("id = ?", ID).Update("deleted_at", nil).Error
}

/* Delete */
func (r *MstSKSLimitRepository) DeleteByID(db *gorm.DB, ctx context.Context, ID string) error {
	user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)
	return db.WithContext(ctx).Table("mst_sks_limits").
		Where("id = ?", ID).
		Updates(map[string]interface{}{
			"deleted_at": time.Now().UnixMilli(),
			"deleted_by": user.ID,
		}).Error
}
