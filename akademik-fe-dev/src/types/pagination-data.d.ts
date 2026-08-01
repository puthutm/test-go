interface PaginationData<T> {
  metadata: {
    page: number;
    size: number;
    total_data: number;
    total_page: number;
  };
  data: T[];
}

interface PaginationDataReferensi<T> {
  metadata: {
    page: number;
    sub_total: number;
    per_page: number;
    total: number;
    total_pages: number;
  };
  data: T[];
}
