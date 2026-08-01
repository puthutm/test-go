package modeldatareferensi

type AcademicPeriodDetailWithSessionReponse struct {
	ID                     string `json:"id"`
	Code                   string `json:"code"`
	AcademicYearID         string `json:"academic_year_id"`
	AcademicYear           string `json:"academic_year"`
	SemesterID             string `json:"semester_id"`
	Semester               string `json:"semester"`
	Fullname               string `json:"fullname"`
	Shortname              string `json:"shortname"`
	StartDateOfCollege     string `json:"start_date_of_college"`
	EndDateOfCollege       string `json:"end_date_of_college"`
	StartDateOfUts         string `json:"start_date_of_uts"`
	EndDateOfUts           string `json:"end_date_of_uts"`
	StartDateOfUas         string `json:"start_date_of_uas"`
	EndDateOfUas           string `json:"end_date_of_uas"`
	NumberOfLectureMeeting string `json:"number_of_lecture_meeting"`
	IsActive               bool   `json:"is_active"`
	UtsSession             int    `json:"uts_session"`
	UasSession             int    `json:"uas_session"`
	CreatedAt              int64  `json:"created_at"`
	UpdatedAt              int64  `json:"updated_at"`
}

func (p *AcademicPeriodDetailWithSessionReponse) GetID() string {
	return p.ID
}
