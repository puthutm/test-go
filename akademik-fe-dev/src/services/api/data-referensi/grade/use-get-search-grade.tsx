"use client";

import { useQuery } from "@tanstack/react-query";

import { getSearchGrade } from "./get-search-grade";

export const useGetSearchGrade = (queryParams: QueryParamDataRefensi) => {
  return useQuery({
    queryKey: [
      "search-grade",
      queryParams.page,
      queryParams.filter,
      queryParams.page_size,
      queryParams.sort_by,
      queryParams.sort_direction,
    ],
    queryFn: async () => await getSearchGrade(queryParams),
    enabled: false,
  });
};
