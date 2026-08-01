package utils

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"unsia.ac.id/akademic_be/pkg/icems-tools/dto/pagination"
)

var (
	ErrInvalidSortDirection = errors.New("sort must be 'asc' or 'desc'")
	ErrInvalidPageNumber    = errors.New("page number must be >= 1")
	ErrInvalidPageSize      = errors.New("page size must be >= 1")
)

func ValidateAndPrepareRequest(pageable pagination.PageableRequestInterface) (pagination.PageableRequestInterface, error) {
	if pageable.GetDefaultSortBy() == "" {
		pageable.SetDefaultSortBy("created_at")
	}

	if pageable.GetDefaultSort() == "" {
		pageable.SetDefaultSort("asc")
	} else {
		pageable.SetDefaultSort(strings.ToLower(pageable.GetDefaultSort()))
		if pageable.GetDefaultSort() != "asc" && pageable.GetDefaultSort() != "desc" {
			return pageable, ErrInvalidSortDirection
		}
	}

	if pageable.GetDefaultPage() <= 0 {
		pageable.SetDefaultPage(1)
	}

	if pageable.GetDefaultLimit() <= 0 {
		pageable.SetDefaultLimit(15)
	} else if pageable.GetDefaultLimit() > 100 {
		pageable.SetDefaultLimit(100)
	}

	if len(pageable.GetDefaultSearch()) > 255 {
		pageable.SetDefaultSearch(pageable.GetDefaultSearch()[:255])
	}

	return pageable, nil
}

func TotalPage(totalData int64, limit int) int {
	if totalData == 0 {
		return 0
	}
	return int(math.Ceil(float64(totalData) / float64(limit)))
}

func GenerateOffset(page, limit int) int {
	return (page - 1) * limit
}

func FormatPaginationInfo(pageable pagination.PageableRequestInterface) string {
	return fmt.Sprintf("Page: %d, Limit: %d, SortBy: %s, Sort: %s", pageable.GetDefaultPage(), pageable.GetDefaultLimit(), pageable.GetDefaultSortBy(), pageable.GetDefaultSort())
}
