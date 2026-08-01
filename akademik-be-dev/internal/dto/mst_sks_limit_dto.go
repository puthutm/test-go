package dto

import (
	"github.com/google/uuid"
)

/* Request */
type MstSKSLimitRequest struct {
	ID       string `json:"-"`
	IPSMin   string `json:"ips_min"  validate:"required,numeric"`
	IPSMax   string `json:"ips_max"  validate:"required,numeric"`
	SKSLimit string `json:"sks_limit"  validate:"required,numeric"`
}

/* Response */
type MstSKSLimitResponse struct {
	ID        uuid.UUID `json:"id"`
	IPSMin    float64   `json:"ips_min" gorm:"column:ips_min"`
	IPSMax    float64   `json:"ips_max" gorm:"column:ips_max"`
	SKSLimit  int       `json:"sks_limit" gorm:"column:sks_limit"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt *int64    `json:"updated_at"`
	DeletedAt *int64    `json:"deleted_at"`
}
