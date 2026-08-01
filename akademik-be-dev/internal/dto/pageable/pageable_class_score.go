package pageable

type PageableRequestClassScore struct {
	PageableRequest
	ClassId string `query:"-"`
}

func (p *PageableRequestClassScore) SetDefaultSort(v string) {
	p.Sort = v
}

func (p *PageableRequestClassScore) SetDefaultSortBy(v string) {
	p.SortBy = v
}

func (p *PageableRequestClassScore) SetDefaultPage(v int) {
	p.Page = v
}

func (p *PageableRequestClassScore) SetDefaultLimit(v int) {
	p.Limit = v
}

func (p *PageableRequestClassScore) SetDefaultSearch(v string) {
	p.Search = v
}

func (p *PageableRequestClassScore) GetDefaultSort() string {
	return p.Sort
}

func (p *PageableRequestClassScore) GetDefaultSortBy() string {
	return p.SortBy
}

func (p *PageableRequestClassScore) GetDefaultPage() int {
	return p.Page
}

func (p *PageableRequestClassScore) GetDefaultLimit() int {
	return p.Limit
}

func (p *PageableRequestClassScore) GetDefaultSearch() string {
	return p.Search
}

func (p *PageableRequestClassScore) SetDefaultClassId(v string) {
	p.ClassId = v
}

func (p *PageableRequestClassScore) GetDefaultClassId() string {
	return p.ClassId
}
