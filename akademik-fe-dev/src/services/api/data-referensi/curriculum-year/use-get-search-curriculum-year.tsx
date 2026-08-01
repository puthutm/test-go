"use client";

import { useQuery } from "@tanstack/react-query";

import { getSearchCurriculumYears } from "./get-search-curriculum-year";

export const useSearchCurriculumYear = () => {
  const query = useQuery({
    queryKey: ["search-curriculum-year"],
    queryFn: async () => {
      const data = await getSearchCurriculumYears({
        page: 1,
        filter: "",
        page_size: 100,
      });

      return data;
    },
  });

  return query;
};
