interface QueryParam {
  page: number;
  limit?: number | null;
  search?: string | null;
  sort_by?: string | null;
  sort_direction?: string | null;
}

interface QueryParamDataRefensi {
  page: number;
  page_size?: number | null;
  filter?: string | null;
  sort_by?: string | null;
  sort_direction?: string | null;
}
