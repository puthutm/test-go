"use client";

import { useQuery } from "@tanstack/react-query";

import { getSearchAcademicPeriods } from "./get-search-academic-period";

export const useGetSearchAcademicPeriod = () => {
  return useQuery({
    queryKey: ["search-academic-periods"],
    queryFn: async () =>
      await getSearchAcademicPeriods({
        page: 1,
        page_size: 1000,
      }),
  });
};
