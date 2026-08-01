package dto

import "github.com/google/uuid"

// TrxKrsLecturerStudentResponse - Response DTO untuk mahasiswa KRS
type TrxKrsLecturerStudentResponse struct {
	KrsHeaderID         uuid.UUID `json:"krs_header_id"`
	AcademicPeriodeID   uuid.UUID `json:"academic_periode_id"`
	AcademicPeriodeName string    `json:"academic_periode_name"`
	StudentName         string    `json:"student_name"`
	StudentNIM          string    `json:"student_nim"`
	TotalSKS            int       `json:"total_sks"`
	StudentID           uuid.UUID `json:"student_id"`
	CreatedAt           int64     `json:"created_at"`
}

// TrxKrsLecturerStudentDetailResponse - Response DTO untuk detail KRS mahasiswa
type TrxKrsLecturerStudentDetailResponse struct {
	StudentID           uuid.UUID                   `json:"student_id"`
	StudentNIM          string                      `json:"student_nim"`
	StudentName         string                      `json:"student_name"`
	StudyProgramName    string                      `json:"study_program_name"`
	AcademicPeriodeName string                      `json:"academic_periode_name"`
	TotalSKSTaken       int                         `json:"total_sks_taken"`
	KRSItems            []TrxKrsLecturerItemResponse `json:"krs_items"`
}

// TrxKrsLecturerItemResponse - Response DTO untuk item KRS
type TrxKrsLecturerItemResponse struct {
	KrsItemID     uuid.UUID `json:"krs_item_id"`
	SubjectNameID *string   `json:"subject_name_id"`
	SubjectNameEN *string   `json:"subject_name_en"`
	ClassCode     string    `json:"class_code"`
	ClassName     string    `json:"class_name"`
	TotalSKS      int       `json:"total_sks"`
}

// TrxKrsLecturerStudentItemUpdateStatusRequest - Request DTO untuk update status KRS item
type TrxKrsLecturerStudentItemUpdateStatusRequest struct {
	KrsItemID    string  `json:"-" validate:"required,uuid"` // from path parameter
	ItemStatus   string  `json:"item_status" validate:"required,oneof=approved rejected"`
	RejectReason *string `json:"reject_reason" validate:"omitempty,stringMax=500"`
}

// TrxKrsLecturerStudentItemUpdateResponse - Response DTO untuk update status KRS item
type TrxKrsLecturerStudentItemUpdateResponse struct {
	KrsItemID    uuid.UUID `json:"krs_item_id"`
	ItemStatus   string    `json:"item_status"`
	RejectReason *string   `json:"reject_reason,omitempty"`
	UpdatedAt    int64     `json:"updated_at"`
	UpdatedBy    uuid.UUID `json:"updated_by"`
}

// TrxKrsPickClassGetRequest - Request DTO untuk GetPickClassesByUserID
type TrxKrsPickClassGetRequest struct {
	AcademicPeriodeID *string `query:"academic_periode_id"` // Optional: jika kosong, gunakan periode aktif
}

// TrxKrsAcademicPeriodResponse - Response DTO untuk periode akademik
type TrxKrsAcademicPeriodResponse struct {
	ID                 uuid.UUID `json:"id"`
	Code               *string   `json:"code"`
	Fullname           *string   `json:"fullname"`
	Shortname          *string   `json:"shortname"`
	IsActive           *bool     `json:"is_active"`
	StartDateOfCollege *int64    `json:"start_date_of_college"`
	EndDateOfCollege   *int64    `json:"end_date_of_college"`
}

// TrxKrsClassForPickResponse - Response DTO untuk kelas yang tersedia diambil
type TrxKrsClassForPickResponse struct {
	ClassID     uuid.UUID `json:"class_id"`
	SubjectID   uuid.UUID `json:"subject_id"`
	SubjectCode string    `json:"subject_code"`
	SubjectName string    `json:"subject_name"`
	Schedule    string    `json:"schedule"`
	ClassCode   string    `json:"class_code"`
	ClassName   string    `json:"class_name"`
	SKS         int       `json:"sks"`
	Capacity    int       `json:"capacity"`
	UsedQuota   int       `json:"used_quota"`
	QuotaText   string    `json:"quota_text"`
	ButtonState string    `json:"button_state"` // "taken" | "full" | "take"
}

// TrxKrsPickClassResponse - Nested response structure berisi academic_periods dan classes
type TrxKrsPickClassResponse struct {
	AcademicPeriods []TrxKrsAcademicPeriodResponse `json:"academic_periods"`
	Classes          []TrxKrsClassForPickResponse   `json:"classes"`
}

// TrxKrsProgramHeadClassResponse - Response DTO untuk kelas KRS program head
type TrxKrsProgramHeadClassResponse struct {
	AcademicPeriodeName string  `json:"academic_periode_name"`
	ClassCode           string  `json:"class_code"`
	ClassName           string  `json:"class_name"`
	LecturerName        string  `json:"lecturer_name"`
	Schedule            string  `json:"schedule"`
	ClassQuota          int     `json:"class_quota"`
	Filled              int     `json:"filled"`
	Remaining           int     `json:"remaining"`
	ClassID             uuid.UUID `json:"class_id"`
	CreatedAt           int64   `json:"created_at"`
}

// TrxKrsSavedGetRequest - Request DTO untuk GetSavedByUserID
type TrxKrsSavedGetRequest struct {
	AcademicPeriodeID *string `query:"academic_periode_id"` // Optional: jika kosong, gunakan periode aktif
}

// TrxKrsSavedItemResponse - Response DTO untuk kelas yang sudah disimpan dalam KRS
type TrxKrsSavedItemResponse struct {
	KrsItemID     uuid.UUID `json:"krs_item_id"`
	KrsID         uuid.UUID `json:"krs_id"`
	ClassID       uuid.UUID `json:"class_id"`
	SubjectID     uuid.UUID `json:"subject_id"`
	SubjectCode   string    `json:"subject_code"`
	SubjectName   string    `json:"subject_name"`
	Schedule      string    `json:"schedule"`
	LecturerNames string    `json:"lecturer_names"`
	ClassName     string    `json:"class_name"`
	SKS           int       `json:"sks"`
	ItemStatus    string    `json:"item_status"` // "waiting" | "approved" | "rejected"
}

// TrxKrsTakeClassRequest - Request untuk mengambil kelas
type TrxKrsTakeClassRequest struct {
	ClassID uuid.UUID `json:"class_id" validate:"required,uuid"`
}

// TrxKrsTakeClassResponse - Response setelah mengambil kelas
type TrxKrsTakeClassResponse struct {
	KrsID                 uuid.UUID  `json:"krs_id"`
	KrsItemID             uuid.UUID  `json:"krs_item_id"`
	AcademicPeriodeID     uuid.UUID  `json:"academic_periode_id"`
	PrevAcademicPeriodeID *uuid.UUID `json:"prev_academic_periode_id,omitempty"`
	PrevIPS               *float64   `json:"prev_ips,omitempty"`
	MaxSKSAllowed         int        `json:"max_sks_allowed"`
	TotalSKSSelected      int        `json:"total_sks_selected"`
	RemainingSKS          int        `json:"remaining_sks"`
	ClassCapacity         int        `json:"class_capacity"`
	UsedQuotaAfter        int        `json:"used_quota_after"`
}

// TrxKrsMaxSksInfoResponse - Response DTO untuk informasi maksimal SKS mahasiswa
type TrxKrsMaxSksInfoResponse struct {
	StudentID               uuid.UUID  `json:"student_id"`
	AcademicPeriodeID       uuid.UUID  `json:"academic_periode_id"`
	AcademicPeriodeIDBefore *uuid.UUID `json:"academic_periode_id_before,omitempty"`
	IpsBefore               float64    `json:"ips_before"`
	SksLimitID              uuid.UUID  `json:"sks_limit_id"`
	IpsMin                  float64    `json:"ips_min"`
	IpsMax                  float64    `json:"ips_max"`
	MaxSks                  int        `json:"max_sks"`
}