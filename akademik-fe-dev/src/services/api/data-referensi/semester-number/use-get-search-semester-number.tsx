"use client";

import { useQuery } from "@tanstack/react-query";

import { getSearchSemesterNumber } from "./get-search-semester-number";

export const useGetSearchSemesterNumber = (
  queryParams: QueryParamDataRefensi
) => {
  return useQuery({
    queryKey: [
      "search-semester-number",
      queryParams.page,
      queryParams.filter,
      queryParams.page_size,
      queryParams.sort_by,
      queryParams.sort_direction,
    ],
    queryFn: async () => await getSearchSemesterNumber(queryParams),
    enabled: false,
  });
};
