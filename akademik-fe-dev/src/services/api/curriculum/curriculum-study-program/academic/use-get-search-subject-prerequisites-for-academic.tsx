"use client";

import { useQuery } from "@tanstack/react-query";

import { getSearchSubjectPrerequisitesForAcademic } from "./get-search-subject-prerequisites-for-academic";

export const useGetSearchSubjectPrerequisitesForAcademic = ({
  curriculumYearId,
  semesterNumberId,
  studyProgramId,
}: SubjectPrerequisitesParams) => {
  return useQuery({
    queryKey: [
      "search-subject-prerequisites",
      curriculumYearId,
      semesterNumberId,
      studyProgramId,
    ],
    queryFn: async () =>
      await getSearchSubjectPrerequisitesForAcademic({
        curriculumYearId,
        semesterNumberId,
        studyProgramId,
      }),
    enabled: false,
  });
};
