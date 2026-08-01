package pageable

type PageableRequestClass struct {
	PageableRequest
	AcademicPeriodeId string  `query:"academic_periode_id"`
	StudyProgramId    *string `query:"study_program_id"`
}

func (p *PageableRequestClass) SetDefaultSort(v string) {
	p.Sort = v
}

func (p *PageableRequestClass) SetDefaultSortBy(v string) {
	p.SortBy = v
}

func (p *PageableRequestClass) SetDefaultPage(v int) {
	p.Page = v
}

func (p *PageableRequestClass) SetDefaultLimit(v int) {
	p.Limit = v
}

func (p *PageableRequestClass) SetDefaultSearch(v string) {
	p.Search = v
}

func (p *PageableRequestClass) SetDefaultAcademicPeriodeId(v string) {
	p.AcademicPeriodeId = v
}

func (p *PageableRequestClass) SetDefaultStudyProgramId(v string) {
	if v == "" {
		p.StudyProgramId = nil
	} else {
		p.StudyProgramId = &v
	}
}

func (p *PageableRequestClass) GetDefaultSort() string {
	return p.Sort
}

func (p *PageableRequestClass) GetDefaultSortBy() string {
	return p.SortBy
}

func (p *PageableRequestClass) GetDefaultPage() int {
	return p.Page
}

func (p *PageableRequestClass) GetDefaultLimit() int {
	return p.Limit
}

func (p *PageableRequestClass) GetDefaultSearch() string {
	return p.Search
}

func (p *PageableRequestClass) GetDefaultAcademicPeriodeId() string {
	return p.AcademicPeriodeId
}

func (p *PageableRequestClass) GetDefaultStudyProgramId() *string {
	if p.StudyProgramId == nil || *p.StudyProgramId == "" {
		return nil
	}
	return p.StudyProgramId
}

type PageableClassGetByUserAndAcademicPeriodAndSubject struct {
	PageableRequest
	AcademicPeriodID *string `query:"-" validate:"required"`
	SubjectID        *string `query:"-" validate:"required"`
}
