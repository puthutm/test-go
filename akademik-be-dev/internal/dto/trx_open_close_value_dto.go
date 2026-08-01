package dto

type TrxOpenCloseValueRequest struct {
	AcademicPeriodeID string `json:"-"`
	ClassID           string `json:"-"`
	StatusLocked      bool   `json:"status_locked"`
}
