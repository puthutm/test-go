"use client";

import { useQuery } from "@tanstack/react-query";
import { getSearchFieldOfStudy } from "./get-search-field-of-study";

export const useSearchFieldOfStudy = () => {
  const query = useQuery({
    queryKey: ["search-field-of-study"],
    queryFn: async () => {
      const data = await getSearchFieldOfStudy({
        page: 1,
        filter: "",
        page_size: 100,
      });

      return data;
    },
  });

  return query;
};
