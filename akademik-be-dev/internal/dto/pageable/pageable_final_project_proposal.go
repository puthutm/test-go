package pageable

// TODO: PageableRequest for subjec
type PageableRequestFinalProjectProposal struct {
	PageableRequest
	StudyProgramId    *string `query:"study_program_id"`
	Status            *string `query:"status"`
	AcademicPeriodeId *string `query:"academic_period_id"`
}

func (p *PageableRequestFinalProjectProposal) SetDefaultSort(v string) {
	p.Sort = v
}

func (p *PageableRequestFinalProjectProposal) SetDefaultSortBy(v string) {
	p.SortBy = v
}

func (p *PageableRequestFinalProjectProposal) SetDefaultPage(v int) {
	p.Page = v
}

func (p *PageableRequestFinalProjectProposal) SetDefaultLimit(v int) {
	p.Limit = v
}

func (p *PageableRequestFinalProjectProposal) SetDefaultSearch(v string) {
	p.Search = v
}

func (p *PageableRequestFinalProjectProposal) GetDefaultSort() string {
	return p.Sort
}

func (p *PageableRequestFinalProjectProposal) GetDefaultSortBy() string {
	return p.SortBy
}

func (p *PageableRequestFinalProjectProposal) GetDefaultPage() int {
	return p.Page
}

func (p *PageableRequestFinalProjectProposal) GetDefaultLimit() int {
	return p.Limit
}

func (p *PageableRequestFinalProjectProposal) GetDefaultSearch() string {
	return p.Search
}
