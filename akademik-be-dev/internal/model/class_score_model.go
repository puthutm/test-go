package model

type ClassScore struct {
	RowNo            int      `gorm:"column:row_no"`
	StudentID        string   `gorm:"column:student_id"`
	Nim              string   `gorm:"column:nim"`
	StudentName      string   `gorm:"column:student_name"`
	PresenceScore    float64  `gorm:"column:presence_score"`
	TaskScore        float64  `gorm:"column:task_score"`
	UtsScore         float64  `gorm:"column:uts_score"`
	UasScore         float64  `gorm:"column:uas_score"`
	FinalScore       float64  `gorm:"column:final_score"`
	QualityValue     *float64 `gorm:"column:quality_value"`
	GradeID          *string  `gorm:"column:grade_id"`
	GradeCode        *string  `gorm:"column:grade_code"`
	GradeName        *string  `gorm:"column:grade_name"`
	GradeDescription *string  `gorm:"column:grade_description"`
	IsPassed         bool     `gorm:"column:is_passed"`
	PassNote         string   `gorm:"column:pass_note"`
	LimitGradeID     *string  `gorm:"column:limit_grade_id"`
	LimitGradeCode   *string  `gorm:"column:limit_grade_code"`
	LimitGradeLower  *float64 `gorm:"column:limit_grade_lower_limit"`
}

func (ClassScore) TableName() string {
	return "class_scores"
}

type ClassScoreSummary struct {
	TotalStudents     int      `gorm:"column:total_students"`
	TotalPassed       int      `gorm:"column:total_passed"`
	TotalNotPassed    int      `gorm:"column:total_not_passed"`
	AverageFinalScore float64  `gorm:"column:average_final_score"`
	LimitGradeCode    *string  `gorm:"column:limit_grade_code"`
	LimitGradeLower   *float64 `gorm:"column:limit_grade_lower_limit"`
}

func (ClassScoreSummary) TableName() string {
	return "class_score_summaries"
}
