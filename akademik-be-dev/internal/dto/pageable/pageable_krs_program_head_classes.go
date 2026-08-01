package pageable

// PageableKrsProgramHeadClassesRequest - Pageable request for KRS program head classes
type PageableKrsProgramHeadClassesRequest struct {
	PageableRequest
	AcademicPeriodeID *string `query:"academic_periode_id"`
}

func (p *PageableKrsProgramHeadClassesRequest) SetDefaultSort(v string) {
	p.Sort = v
}

func (p *PageableKrsProgramHeadClassesRequest) SetDefaultSortBy(v string) {
	p.SortBy = v
}

func (p *PageableKrsProgramHeadClassesRequest) SetDefaultPage(v int) {
	p.Page = v
}

func (p *PageableKrsProgramHeadClassesRequest) SetDefaultLimit(v int) {
	p.Limit = v
}

func (p *PageableKrsProgramHeadClassesRequest) SetDefaultSearch(v string) {
	p.Search = v
}

func (p *PageableKrsProgramHeadClassesRequest) SetDefaultAcademicPeriodeId(v string) {
	if v == "" {
		p.AcademicPeriodeID = nil
	} else {
		p.AcademicPeriodeID = &v
	}
}

func (p *PageableKrsProgramHeadClassesRequest) GetDefaultSort() string {
	return p.Sort
}

func (p *PageableKrsProgramHeadClassesRequest) GetDefaultSortBy() string {
	return p.SortBy
}

func (p *PageableKrsProgramHeadClassesRequest) GetDefaultPage() int {
	return p.Page
}

func (p *PageableKrsProgramHeadClassesRequest) GetDefaultLimit() int {
	return p.Limit
}

func (p *PageableKrsProgramHeadClassesRequest) GetDefaultSearch() string {
	return p.Search
}

func (p *PageableKrsProgramHeadClassesRequest) GetDefaultAcademicPeriodeId() *string {
	if p.AcademicPeriodeID == nil || *p.AcademicPeriodeID == "" {
		return nil
	}
	return p.AcademicPeriodeID
}
