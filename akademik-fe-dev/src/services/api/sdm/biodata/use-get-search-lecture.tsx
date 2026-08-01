"use client";

import { useQuery } from "@tanstack/react-query";
import { getSearchLecturer, LecturerParams } from "./get-search-lecture";

export const useSearchLecturer = (queryParam: LecturerParams) => {
  const query = useQuery({
    queryKey: [
      "search-lecturer",
      queryParam.page,
      queryParam.limit,
      queryParam.search,
    ],
    queryFn: async () => {
      const data = await getSearchLecturer({ ...queryParam });

      return data;
    },
    enabled: !!queryParam.search,
    staleTime: 0,
  });

  return query;
};
