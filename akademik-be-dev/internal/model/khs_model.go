package model

type KhsSubject struct {
	AcademicPeriodeId   string  `gorm:"column:academic_periode_id"`
	AcademicPeriodeName string  `gorm:"column:academic_periode_name"`
	SubjectCode         string  `gorm:"column:subject_code"`
	SubjectName         string  `gorm:"column:subject_name"`
	TotalSks            int     `gorm:"column:total_sks"`
	FinalScore          float64 `gorm:"column:final_score"`
	GradeCode           string  `gorm:"column:grade_code"`
	QualityValue        float64 `gorm:"column:quality_value"`
	Bobot               float64 `gorm:"column:bobot"`
	IsPassed            int     `gorm:"column:is_passed"`
}

type KhsSemester struct {
	AcademicPeriodeId   string  `gorm:"column:academic_periode_id"`
	AcademicPeriodeName string  `gorm:"column:academic_periode_name"`
	TotalBobot          float64 `gorm:"column:total_bobot"`
	TotalSks            int     `gorm:"column:total_sks"`
	Ips                 float64 `gorm:"column:ips"`
}

type KhsNotPassed struct {
	AcademicPeriodeId   string  `gorm:"column:academic_periode_id"`
	AcademicPeriodeName string  `gorm:"column:academic_periode_name"`
	SubjectCode         string  `gorm:"column:subject_code"`
	SubjectName         string  `gorm:"column:subject_name"`
	FinalScore          float64 `gorm:"column:final_score"`
	GradeCode           string  `gorm:"column:grade_code"`
}
