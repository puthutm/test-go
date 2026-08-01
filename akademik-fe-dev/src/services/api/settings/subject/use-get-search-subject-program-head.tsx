"use client";

import { useQuery } from "@tanstack/react-query";
import {
  getSearchSubjectForProgramHead,
  SubjectProgramHeadParams,
} from "./get-search-subject-program-head";

export const useGetSearchSubjectProgramHead = (
  queryParams: SubjectProgramHeadParams
) => {
  return useQuery({
    queryKey: [
      "search-subject-for-program-head",
      queryParams.curriculum_year_id,
      queryParams.search,
    ],
    queryFn: async () =>
      await getSearchSubjectForProgramHead({
        ...queryParams,
      }),
    enabled: false,
  });
};
