package pageable

// TODO: PageableRequest for class participant
type PageableRequestClassParticipant struct {
	PageableRequest
	ClassID string `query:"class_id"`
}

func (p *PageableRequestClassParticipant) SetDefaultSort(v string) {
	p.Sort = v
}

func (p *PageableRequestClassParticipant) SetDefaultSortBy(v string) {
	p.SortBy = v
}

func (p *PageableRequestClassParticipant) SetDefaultPage(v int) {
	p.Page = v
}

func (p *PageableRequestClassParticipant) SetDefaultLimit(v int) {
	p.Limit = v
}

func (p *PageableRequestClassParticipant) SetDefaultSearch(v string) {
	p.Search = v
}

func (p *PageableRequestClassParticipant) GetDefaultSort() string {
	return p.Sort
}

func (p *PageableRequestClassParticipant) GetDefaultSortBy() string {
	return p.SortBy
}

func (p *PageableRequestClassParticipant) GetDefaultPage() int {
	return p.Page
}

func (p *PageableRequestClassParticipant) GetDefaultLimit() int {
	return p.Limit
}

func (p *PageableRequestClassParticipant) GetDefaultSearch() string {
	return p.Search
}
