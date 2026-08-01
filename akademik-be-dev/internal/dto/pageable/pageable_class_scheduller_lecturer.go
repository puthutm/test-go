package pageable

type PageableRequestClassScheduleLecturer struct {
	PageableRequest
	StudyProgramId   *string `query:"study_program_id"`
	CurriculumYearId *string `query:"curriculum_year_id"`
	ClassId          *string `query:"class_id"`
}

func (p *PageableRequestClassScheduleLecturer) SetDefaultSort(v string) {
	p.Sort = v
}

func (p *PageableRequestClassScheduleLecturer) SetDefaultSortBy(v string) {
	p.SortBy = v
}

func (p *PageableRequestClassScheduleLecturer) SetDefaultPage(v int) {
	p.Page = v
}

func (p *PageableRequestClassScheduleLecturer) SetDefaultLimit(v int) {
	p.Limit = v
}

func (p *PageableRequestClassScheduleLecturer) SetDefaultSearch(v string) {
	p.Search = v
}

func (p *PageableRequestClassScheduleLecturer) SetDefaultStudyProgramId(v string) {
	if v == "" {
		p.StudyProgramId = nil
	} else {
		p.StudyProgramId = &v
	}
}

func (p *PageableRequestClassScheduleLecturer) SetDefaultCurriculumYearId(v string) {
	p.CurriculumYearId = &v
}

func (p *PageableRequestClassScheduleLecturer) SetDefaultClassId(v string) {
	p.ClassId = &v
}

func (p *PageableRequestClassScheduleLecturer) GetDefaultSort() string {
	return p.Sort
}

func (p *PageableRequestClassScheduleLecturer) GetDefaultSortBy() string {
	return p.SortBy
}

func (p *PageableRequestClassScheduleLecturer) GetDefaultPage() int {
	return p.Page
}

func (p *PageableRequestClassScheduleLecturer) GetDefaultLimit() int {
	return p.Limit
}

func (p *PageableRequestClassScheduleLecturer) GetDefaultSearch() string {
	return p.Search
}

func (p *PageableRequestClassScheduleLecturer) GetDefaultStudyProgramId() *string {
	if p.StudyProgramId == nil || *p.StudyProgramId == "" {
		return nil
	}
	return p.StudyProgramId
}

func (p *PageableRequestClassScheduleLecturer) GetDefaultCurriculumYearId() *string {
	if p.CurriculumYearId == nil || *p.CurriculumYearId == "" {
		return nil
	}
	return p.CurriculumYearId
}

func (p *PageableRequestClassScheduleLecturer) GetDefaultClassId() *string {
	if p.ClassId == nil || *p.ClassId == "" {
		return nil
	}
	return p.ClassId
}
