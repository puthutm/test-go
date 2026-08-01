package pageable

// TODO: PageableRequest for value composition
type PageableRequestValueComposition struct {
	PageableRequest
	ValueElementID    string `query:"value_element_id"`
	AcademicPeriodeID string `query:"academic_periode_id"`
}

func (p *PageableRequestValueComposition) SetDefaultSort(v string) {
	p.Sort = v
}

func (p *PageableRequestValueComposition) SetDefaultSortBy(v string) {
	p.SortBy = v
}

func (p *PageableRequestValueComposition) SetDefaultPage(v int) {
	p.Page = v
}

func (p *PageableRequestValueComposition) SetDefaultLimit(v int) {
	p.Limit = v
}

func (p *PageableRequestValueComposition) SetDefaultSearch(v string) {
	p.Search = v
}

func (p *PageableRequestValueComposition) GetDefaultSort() string {
	return p.Sort
}

func (p *PageableRequestValueComposition) GetDefaultSortBy() string {
	return p.SortBy
}

func (p *PageableRequestValueComposition) GetDefaultPage() int {
	return p.Page
}

func (p *PageableRequestValueComposition) GetDefaultLimit() int {
	return p.Limit
}

func (p *PageableRequestValueComposition) GetDefaultSearch() string {
	return p.Search
}
