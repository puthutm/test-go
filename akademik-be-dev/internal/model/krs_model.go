package model

import (
	"time"

	"github.com/google/uuid"
)

// TrxKrsLecturerStudent - Result set dari SP sp_krs_lecturer_students_get
// Berisi data mahasiswa yang mengajukan KRS untuk kelas yang diajar oleh dosen
type TrxKrsLecturerStudent struct {
	RowNum              int       `gorm:"column:RowNum" json:"-"`
	KrsHeaderID         uuid.UUID `gorm:"column:krs_header_id" json:"krs_header_id"`
	AcademicPeriodeID   uuid.UUID `gorm:"column:academic_periode_id" json:"academic_periode_id"`
	AcademicPeriodeName string    `gorm:"column:academic_periode_name" json:"academic_periode_name"`
	StudentName         string    `gorm:"column:student_name" json:"student_name"`
	StudentNIM          string    `gorm:"column:student_nim" json:"student_nim"`
	TotalSKS            int       `gorm:"column:total_sks" json:"total_sks"`
	StudentID           uuid.UUID `gorm:"column:student_id" json:"student_id"`
	CreatedAt           int64     `gorm:"column:created_at" json:"created_at"`
}

// TrxKrsLecturerStudentDetail - Result set 1: Student detail dari SP sp_krs_lecturer_students_get_by_krs_header_id
type TrxKrsLecturerStudentDetail struct {
	StudentID           uuid.UUID `gorm:"column:student_id" json:"student_id"`
	StudentNIM          string    `gorm:"column:student_nim" json:"student_nim"`
	StudentName         string    `gorm:"column:student_name" json:"student_name"`
	StudyProgramName    string    `gorm:"column:sutudy_program_name" json:"study_program_name"`
	AcademicPeriodeName string    `gorm:"column:academic_periode_name" json:"academic_periode_name"`
}

// TrxKrsLecturerStudentTotalSKS - Result set 2: Total SKS taken
type TrxKrsLecturerStudentTotalSKS struct {
	TotalSKSTaken int `gorm:"column:total_sks_taken" json:"total_sks_taken"`
}

// TrxKrsLecturerStudentItem - Result set 3: KRS items (subjects with waiting status)
type TrxKrsLecturerStudentItem struct {
	KrsItemID     uuid.UUID `gorm:"column:krs_item_id" json:"krs_item_id"`
	SubjectNameID *string   `gorm:"column:subject_name_id" json:"subject_name_id"`
	SubjectNameEN *string   `gorm:"column:subject_name_en" json:"subject_name_en"`
	ClassCode     string    `gorm:"column:class_code" json:"class_code"`
	ClassName     string    `gorm:"column:class_name" json:"class_name"`
	TotalSKS      int       `gorm:"column:total_sks" json:"total_sks"`
}

// TrxKrsLecturerStudentItemUpdate - Model untuk update KRS item status
type TrxKrsLecturerStudentItemUpdate struct {
	KrsItemID    uuid.UUID `gorm:"column:krs_item_id" json:"krs_item_id"`
	ItemStatus   string    `gorm:"column:item_status" json:"item_status"`
	RejectReason *string   `gorm:"column:reject_reason" json:"reject_reason,omitempty"`
	UpdatedAt    int64     `gorm:"column:updated_at" json:"updated_at"`
	UpdatedBy    uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
}

// TrxKrsAcademicPeriod - Result set 1 dari SP sp_krs_pick_classes_get_by_user_id
// Berisi daftar periode akademik untuk dropdown/filter
type TrxKrsAcademicPeriod struct {
	ID                 uuid.UUID  `gorm:"column:id" json:"id"`
	Code               *string    `gorm:"column:code" json:"code"`
	Fullname           *string    `gorm:"column:fullname" json:"fullname"`
	Shortname          *string    `gorm:"column:shortname" json:"shortname"`
	IsActive           *bool      `gorm:"column:is_active" json:"is_active"`
	StartDateOfCollege *time.Time `gorm:"column:start_date_of_college" json:"start_date_of_college"`
	EndDateOfCollege   *time.Time `gorm:"column:end_date_of_college" json:"end_date_of_college"`
}

// TrxKrsClassForPick - Result set 2 dari SP sp_krs_pick_classes_get_by_user_id
// Berisi daftar kelas yang tersedia dengan info kuota, jadwal, dan status button
type TrxKrsClassForPick struct {
	ClassID     uuid.UUID `gorm:"column:class_id" json:"class_id"`
	SubjectID   uuid.UUID `gorm:"column:subject_id" json:"subject_id"`
	SubjectCode string    `gorm:"column:subject_code" json:"subject_code"`
	SubjectName string    `gorm:"column:subject_name" json:"subject_name"`
	Schedule    string    `gorm:"column:schedule" json:"schedule"`
	ClassCode   string    `gorm:"column:class_code" json:"class_code"`
	ClassName   string    `gorm:"column:class_name" json:"class_name"`
	SKS         int       `gorm:"column:sks" json:"sks"`
	Capacity    int       `gorm:"column:capacity" json:"capacity"`
	UsedQuota   int       `gorm:"column:used_quota" json:"used_quota"`
	QuotaText   string    `gorm:"column:quota_text" json:"quota_text"`
	ButtonState string    `gorm:"column:button_state" json:"button_state"` // "taken" | "full" | "take"
}

// TrxKrsProgramHeadClass - Result set dari SP sp_krs_student_classes_get_by_kaprodi_user_id
// Berisi data kelas yang ada dalam program studi ketua program studi untuk approval KRS
type TrxKrsProgramHeadClass struct {
	No                  int       `gorm:"column:no" json:"-"`
	AcademicPeriodeName string    `gorm:"column:academic_periode_name" json:"academic_periode_name"`
	ClassCode           string    `gorm:"column:class_code" json:"class_code"`
	ClassName           string    `gorm:"column:class_name" json:"class_name"`
	LecturerName        string    `gorm:"column:lecturer_name" json:"lecturer_name"`
	Schedule            string    `gorm:"column:schedule" json:"schedule"`
	ClassQuota          int       `gorm:"column:class_quota" json:"class_quota"`
	Filled              int       `gorm:"column:filled" json:"filled"`
	Remaining           int       `gorm:"column:remaining" json:"remaining"`
	ClassID             uuid.UUID `gorm:"column:class_id" json:"class_id"`
	CreatedAt           int64     `gorm:"column:created_at" json:"created_at"`
}

// TrxKrsSavedItem - Result set dari SP sp_krs_saved_get_by_user_id
// Berisi kelas yang sudah disimpan dalam KRS mahasiswa
type TrxKrsSavedItem struct {
	KrsItemID     uuid.UUID `gorm:"column:krs_item_id" json:"krs_item_id"`
	KrsID         uuid.UUID `gorm:"column:krs_id" json:"krs_id"`
	ClassID       uuid.UUID `gorm:"column:class_id" json:"class_id"`
	SubjectID     uuid.UUID `gorm:"column:subject_id" json:"subject_id"`
	SubjectCode   string    `gorm:"column:subject_code" json:"subject_code"`
	SubjectName   string    `gorm:"column:subject_name" json:"subject_name"`
	Schedule      string    `gorm:"column:schedule" json:"schedule"`
	LecturerNames string    `gorm:"column:lecturer_names" json:"lecturer_names"`
	ClassName     string    `gorm:"column:class_name" json:"class_name"`
	SKS           int       `gorm:"column:sks" json:"sks"`
	ItemStatus    string    `gorm:"column:item_status" json:"item_status"` // "waiting" | "approved" | "rejected"
}

// TrxKrsTakeClassResult - Result set dari SP sp_krs_take_class_by_user_id
type TrxKrsTakeClassResult struct {
	KrsID                 uuid.UUID  `gorm:"column:krs_id" json:"krs_id"`
	KrsItemID             uuid.UUID  `gorm:"column:krs_item_id" json:"krs_item_id"`
	AcademicPeriodeID     uuid.UUID  `gorm:"column:academic_periode_id" json:"academic_periode_id"`
	PrevAcademicPeriodeID *uuid.UUID `gorm:"column:prev_academic_periode_id" json:"prev_academic_periode_id,omitempty"`
	PrevIPS               *float64   `gorm:"column:prev_ips" json:"prev_ips,omitempty"`
	MaxSKSAllowed         int        `gorm:"column:max_sks_allowed" json:"max_sks_allowed"`
	TotalSKSSelected      int        `gorm:"column:total_sks_selected" json:"total_sks_selected"`
	RemainingSKS          int        `gorm:"column:remaining_sks" json:"remaining_sks"`
	ClassCapacity         int        `gorm:"column:class_capacity" json:"class_capacity"`
	UsedQuotaAfter        int        `gorm:"column:used_quota_after" json:"used_quota_after"`
}

// TrxKrsMaxSksInfo - Result set dari SP sp_krs_max_sks_get_by_user_id
type TrxKrsMaxSksInfo struct {
	StudentID               uuid.UUID  `gorm:"column:student_id"`
	AcademicPeriodeID       uuid.UUID  `gorm:"column:academic_periode_id"`
	AcademicPeriodeIDBefore *uuid.UUID `gorm:"column:academic_periode_id_before"`
	IpsBefore               float64    `gorm:"column:ips_before"`
	SksLimitID              uuid.UUID  `gorm:"column:sks_limit_id"`
	IpsMin                  float64    `gorm:"column:ips_min"`
	IpsMax                  float64    `gorm:"column:ips_max"`
	MaxSks                  int        `gorm:"column:max_sks"`
}
