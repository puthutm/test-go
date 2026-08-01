"use client";

import { useQuery } from "@tanstack/react-query";

import { getSearchFieldStudyConcentration } from "./get-search-field-study-concentration";

export const useSearchFielStudyConcentration = () => {
  return useQuery({
    queryKey: ["search-field-study-concentration"],
    queryFn: async () => {
      const data = await getSearchFieldStudyConcentration({
        page: 1,
        filter: "",
        page_size: 500,
      });

      return data;
    },
    enabled: false,
  });
};
