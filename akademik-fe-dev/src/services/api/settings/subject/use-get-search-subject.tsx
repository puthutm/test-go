"use client";

import { useQuery } from "@tanstack/react-query";
import { getSearchSubject, SubjectParams } from "./get-search-subject";

export const useGetSearchSubject = (queryParams: SubjectParams) => {
  return useQuery({
    queryKey: [
      "search-subject",
      queryParams.curriculum_year_id,
      queryParams.study_program_id,
    ],
    queryFn: async () =>
      await getSearchSubject({
        ...queryParams,
      }),
    enabled: !!queryParams.curriculum_year_id && !!queryParams.study_program_id,
  });
};
