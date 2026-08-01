package dto

type KhsSubjectResponse struct {
	SubjectCode  string  `json:"subject_code"`
	SubjectName  string  `json:"subject_name"`
	TotalSks     int     `json:"total_sks"`
	FinalScore   float64 `json:"final_score"`
	GradeCode    string  `json:"grade_code"`
	QualityValue float64 `json:"quality_value"`
	Weight       float64 `json:"weight"`
	IsPassed     bool    `json:"is_passed"`
}

type KhsSemesterResponse struct {
	AcademicPeriodeId   string               `json:"academic_periode_id"`
	AcademicPeriodeName string               `json:"academic_periode_name"`
	TotalWeight         float64              `json:"total_weight"`
	TotalSks            int                  `json:"total_sks"`
	Ips                 float64              `json:"ips"`
	Subjects            []KhsSubjectResponse `json:"subjects"`
}

type KhsNotPassedResponse struct {
	AcademicPeriodeId   string  `json:"academic_periode_id"`
	AcademicPeriodeName string  `json:"academic_periode_name"`
	SubjectCode         string  `json:"subject_code"`
	SubjectName         string  `json:"subject_name"`
	FinalScore          float64 `json:"final_score"`
	GradeCode           string  `json:"grade_code"`
}

type KhsDataResponse struct {
	Semesters []KhsSemesterResponse  `json:"semesters"`
	NotPassed []KhsNotPassedResponse `json:"not_passed"`
}
