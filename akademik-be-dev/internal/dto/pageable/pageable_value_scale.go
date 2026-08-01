package pageable

// TODO: PageableRequest for value scale
type PageableRequestValueScale struct {
	PageableRequest
	StudyProgramID string `query:"study_program_id"`
	GradeID        string `query:"grade_id"`
}

func (p *PageableRequestValueScale) SetDefaultSort(v string) {
	p.Sort = v
}

func (p *PageableRequestValueScale) SetDefaultSortBy(v string) {
	p.SortBy = v
}

func (p *PageableRequestValueScale) SetDefaultPage(v int) {
	p.Page = v
}

func (p *PageableRequestValueScale) SetDefaultLimit(v int) {
	p.Limit = v
}

func (p *PageableRequestValueScale) SetDefaultSearch(v string) {
	p.Search = v
}

func (p *PageableRequestValueScale) GetDefaultSort() string {
	return p.Sort
}

func (p *PageableRequestValueScale) GetDefaultSortBy() string {
	return p.SortBy
}

func (p *PageableRequestValueScale) GetDefaultPage() int {
	return p.Page
}

func (p *PageableRequestValueScale) GetDefaultLimit() int {
	return p.Limit
}

func (p *PageableRequestValueScale) GetDefaultSearch() string {
	return p.Search
}


