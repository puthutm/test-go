package dto

type PortalStudentBatchResponse struct {
	BatchDetailID string `json:"batch_detail_id"`
	BatchName     string `json:"batch_name"`
}

type PortalStudentCreateRequest struct {
	ID            string `json:"-"`
	BatchDetailID string `json:"batch_detail_id" validate:"required,uuid"`
	NIK           string `json:"nik" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	Password      string `json:"password"`
	Phone         string `json:"phone" validate:"required"`
}

type PortalStudentBulkFailedRow struct {
	Row   int    `json:"row"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Error string `json:"error"`
}

type PortalStudentBulkResponse struct {
	TotalData  int                                `json:"total_data"`
	Success    int                                `json:"success"`
	Failed     int                                `json:"failed"`
	FailedRows []PortalStudentBulkFailedRow       `json:"failed_rows"`
}

type PortalStudentListResponse struct {
	StudentID                    string   `json:"student_id"`
	StudentStudyProgramID        string   `json:"student_study_program_id"`
	Nim                          string   `json:"nim"`
	StudentName                  string   `json:"student_name"`
	StudentStatus                string   `json:"student_status"`
	StudyProgramID               string   `json:"study_program_id"`
	StudyProgramCode             string   `json:"study_program_code"`
	StudyProgramName             string   `json:"study_program_name"`
	LecturerPABiodataID          *string  `json:"lecturer_pa_biodata_id"`
	LecturerPAUserID             *string  `json:"lecturer_pa_user_id"`
	LecturerPAName               *string  `json:"lecturer_pa_name"`
	EntryAcademicPeriodID        *string  `json:"entry_academic_period_id"`
	EntryAcademicPeriodName      *string  `json:"entry_academic_period_name"`
	CurrentAcademicPeriodID      string   `json:"current_academic_period_id"`
	CurrentAcademicPeriodName    string   `json:"current_academic_period_name"`
	CurrentAcademicPeriodShortname string  `json:"current_academic_period_shortname"`
	EntryPeriodRank              *int     `json:"entry_period_rank"`
	CurrentPeriodRank            int      `json:"current_period_rank"`
	CurrentSemester              *int     `json:"current_semester"`
	SemesterLabel                *string  `json:"semester_label"`
	TotalSksTaken                int      `json:"total_sks_taken"`
	Ipk                          float64  `json:"ipk"`
	CreatedAt                    int64    `json:"created_at"`
}
