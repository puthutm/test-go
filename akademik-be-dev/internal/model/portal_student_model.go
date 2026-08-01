package model

type PortalStudentBatch struct {
	BatchDetailID string `gorm:"column:batch_detail_id"`
	BatchName     string `gorm:"column:batch_name"`
}

type PortalStudentList struct {
	RowNo                        int     `gorm:"column:row_no"`
	StudentID                    string  `gorm:"column:student_id"`
	StudentStudyProgramID        string  `gorm:"column:student_study_program_id"`
	Nim                          string  `gorm:"column:nim"`
	StudentName                  string  `gorm:"column:student_name"`
	StudentStatus                string  `gorm:"column:student_status"`
	StudyProgramID               string  `gorm:"column:study_program_id"`
	StudyProgramCode             string  `gorm:"column:study_program_code"`
	StudyProgramName             string  `gorm:"column:study_program_name"`
	LecturerPABiodataID          *string `gorm:"column:lecturer_pa_biodata_id"`
	LecturerPAUserID             *string `gorm:"column:lecturer_pa_user_id"`
	LecturerPAName               *string `gorm:"column:lecturer_pa_name"`
	EntryAcademicPeriodID        *string `gorm:"column:entry_academic_period_id"`
	EntryAcademicPeriodName      *string `gorm:"column:entry_academic_period_name"`
	CurrentAcademicPeriodID      string  `gorm:"column:current_academic_period_id"`
	CurrentAcademicPeriodName    string  `gorm:"column:current_academic_period_name"`
	CurrentAcademicPeriodShortname string `gorm:"column:current_academic_period_shortname"`
	EntryPeriodRank              *int    `gorm:"column:entry_period_rank"`
	CurrentPeriodRank            int     `gorm:"column:current_period_rank"`
	CurrentSemester              *int    `gorm:"column:current_semester"`
	SemesterLabel                *string `gorm:"column:semester_label"`
	TotalSksTaken                int     `gorm:"column:total_sks_taken"`
	Ipk                          float64 `gorm:"column:ipk"`
	CreatedAt                    int64   `gorm:"column:created_at"`
}
