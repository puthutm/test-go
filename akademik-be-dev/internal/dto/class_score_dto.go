package dto

type ClassScoreResponse struct {
	StudentID        string   `json:"student_id" gorm:"column:student_id"`
	Nim              string   `json:"nim" gorm:"column:nim"`
	StudentName      string   `json:"student_name" gorm:"column:student_name"`
	PresenceScore    float64  `json:"presence_score" gorm:"column:presence_score"`
	TaskScore        float64  `json:"task_score" gorm:"column:task_score"`
	UtsScore         float64  `json:"uts_score" gorm:"column:uts_score"`
	UasScore         float64  `json:"uas_score" gorm:"column:uas_score"`
	FinalScore       float64  `json:"final_score" gorm:"column:final_score"`
	QualityValue     *float64 `json:"quality_value" gorm:"column:quality_value"`
	GradeID          *string  `json:"grade_id" gorm:"column:grade_id"`
	GradeCode        *string  `json:"grade_code" gorm:"column:grade_code"`
	GradeName        *string  `json:"grade_name" gorm:"column:grade_name"`
	GradeDescription *string  `json:"grade_description" gorm:"column:grade_description"`
	IsPassed         bool     `json:"is_passed" gorm:"column:is_passed"`
	PassNote         string   `json:"pass_note" gorm:"column:pass_note"`
	LimitGradeID     *string  `json:"limit_grade_id" gorm:"column:limit_grade_id"`
	LimitGradeCode   *string  `json:"limit_grade_code" gorm:"column:limit_grade_code"`
	LimitGradeLower  *float64 `json:"limit_grade_lower_limit" gorm:"column:limit_grade_lower_limit"`
}

type ClassScoreSummaryResponse struct {
	TotalStudents     int      `json:"total_students" gorm:"column:total_students"`
	TotalPassed       int      `json:"total_passed" gorm:"column:total_passed"`
	TotalNotPassed    int      `json:"total_not_passed" gorm:"column:total_not_passed"`
	AverageFinalScore float64  `json:"average_final_score" gorm:"column:average_final_score"`
	LimitGradeCode    *string  `json:"limit_grade_code" gorm:"column:limit_grade_code"`
	LimitGradeLower   *float64 `json:"limit_grade_lower_limit" gorm:"column:limit_grade_lower_limit"`
}

type ClassScorePageableResponse struct {
	Data     []ClassScoreResponse       `json:"data"`
	Summary  *ClassScoreSummaryResponse `json:"summary"`
	Metadata PageableMetadata           `json:"metadata"`
}

type ClassScoreCheckSaveButtonResponse struct {
	StatusLock bool `json:"status_lock"`
}

type PageableMetadata struct {
	TotalData int64 `json:"total_data"`
	TotalPage int   `json:"total_page"`
	Page      int   `json:"page"`
	Size      int   `json:"size"`
}

type UpdateStatusLockRequest struct {
	AcademicPeriodID string `json:"-"`
	ClassID          string `json:"-"`
	CreatedBy        string `json:"-"`
	StatusLocked     bool   `json:"status_locked"`
	CreatedAt        int64  `json:"-"`
}
