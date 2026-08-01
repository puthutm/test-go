package pageable

// TODO: PageableRequest for subjec
type PageableRequestSubject struct {
	PageableRequest
	StudyProgramID   string `query:"study_program_id"`
	SubjectTypeID    string `query:"subject_type_id"`
	SubjectGroupID   string `query:"subject_group_id"`
	CurriculumYearID string `query:"curriculum_year_id"`
}

func (p *PageableRequestSubject) SetDefaultSort(v string) {
	p.Sort = v
}

func (p *PageableRequestSubject) SetDefaultSortBy(v string) {
	p.SortBy = v
}

func (p *PageableRequestSubject) SetDefaultPage(v int) {
	p.Page = v
}

func (p *PageableRequestSubject) SetDefaultLimit(v int) {
	p.Limit = v
}

func (p *PageableRequestSubject) SetDefaultSearch(v string) {
	p.Search = v
}

func (p *PageableRequestSubject) SetDefaultStudyProgramID(v string) {
	p.StudyProgramID = v
}

func (p *PageableRequestSubject) SetDefaultSubjectTypeID(v string) {
	p.SubjectTypeID = v
}

func (p *PageableRequestSubject) SetDefaultSubjectGroupID(v string) {
	p.SubjectGroupID = v
}

func (p *PageableRequestSubject) SetDefaultCurriculumYearID(v string) {
	p.CurriculumYearID = v
}

func (p *PageableRequestSubject) GetDefaultSort() string {
	return p.Sort
}

func (p *PageableRequestSubject) GetDefaultSortBy() string {
	return p.SortBy
}

func (p *PageableRequestSubject) GetDefaultPage() int {
	return p.Page
}

func (p *PageableRequestSubject) GetDefaultLimit() int {
	return p.Limit
}

func (p *PageableRequestSubject) GetDefaultSearch() string {
	return p.Search
}

func (p *PageableRequestSubject) GetDefaultStudyProgramID() string {
	return p.StudyProgramID
}

func (p *PageableRequestSubject) GetDefaultSubjectTypeID() string {
	return p.SubjectTypeID
}

func (p *PageableRequestSubject) GetDefaultSubjectGroupID() string {
	return p.SubjectGroupID
}

func (p *PageableRequestSubject) GetDefaultCurriculumYearID() string {
	return p.CurriculumYearID
}
