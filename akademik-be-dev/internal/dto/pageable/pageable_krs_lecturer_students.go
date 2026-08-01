package pageable

// PageableKrsLecturerStudentsRequest - Pageable request for KRS lecturer students
type PageableKrsLecturerStudentsRequest struct {
	PageableRequest
	AcademicPeriodeID *string `query:"academic_periode_id"`
}

func (p *PageableKrsLecturerStudentsRequest) SetDefaultSort(v string) {
	p.Sort = v
}

func (p *PageableKrsLecturerStudentsRequest) SetDefaultSortBy(v string) {
	p.SortBy = v
}

func (p *PageableKrsLecturerStudentsRequest) SetDefaultPage(v int) {
	p.Page = v
}

func (p *PageableKrsLecturerStudentsRequest) SetDefaultLimit(v int) {
	p.Limit = v
}

func (p *PageableKrsLecturerStudentsRequest) SetDefaultSearch(v string) {
	p.Search = v
}

func (p *PageableKrsLecturerStudentsRequest) SetDefaultAcademicPeriodeId(v string) {
	if v == "" {
		p.AcademicPeriodeID = nil
	} else {
		p.AcademicPeriodeID = &v
	}
}

func (p *PageableKrsLecturerStudentsRequest) GetDefaultSort() string {
	return p.Sort
}

func (p *PageableKrsLecturerStudentsRequest) GetDefaultSortBy() string {
	return p.SortBy
}

func (p *PageableKrsLecturerStudentsRequest) GetDefaultPage() int {
	return p.Page
}

func (p *PageableKrsLecturerStudentsRequest) GetDefaultLimit() int {
	return p.Limit
}

func (p *PageableKrsLecturerStudentsRequest) GetDefaultSearch() string {
	return p.Search
}

func (p *PageableKrsLecturerStudentsRequest) GetDefaultAcademicPeriodeId() *string {
	if p.AcademicPeriodeID == nil || *p.AcademicPeriodeID == "" {
		return nil
	}
	return p.AcademicPeriodeID
}
