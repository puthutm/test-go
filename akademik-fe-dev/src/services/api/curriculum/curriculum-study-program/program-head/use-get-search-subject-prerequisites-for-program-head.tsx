"use client";

import { useQuery } from "@tanstack/react-query";

import { getSearchSubjectPrerequisitesForProgramHead } from "./get-search-subject-prerequisites-for-program-head";

export const useGetSearchSubjectPrerequisitesForProgramHead = ({
  curriculumYearId,
  semesterNumberId,
}: SubjectPrerequisitesParams) => {
  return useQuery({
    queryKey: [
      "search-subject-prerequisites-program-head",
      curriculumYearId,
      semesterNumberId,
    ],
    queryFn: async () =>
      await getSearchSubjectPrerequisitesForProgramHead({
        curriculumYearId,
        semesterNumberId,
      }),
    enabled: false,
  });
};
