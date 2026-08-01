"use client";

import { useQuery } from "@tanstack/react-query";

import { getCurriculumStudyProgramForAcademic } from "./get-curriculum-study-program-for-academic";
import { AKADEMIK } from "@/lib/constants/role";

export const useGetCurriculumStudyProgramBySemesterIdForAcademic = ({
  semesterNumberId,
  studyProgramId,
  curriculumYearId,
  role,
}: {
  semesterNumberId: string;
  studyProgramId: string;
  curriculumYearId: string;
  role: string;
}) => {
  return useQuery({
    queryKey: [
      "curriculum-study-program-for-academic",
      semesterNumberId,
      studyProgramId,
      curriculumYearId,
      role,
    ],
    queryFn: async () => {
      return await getCurriculumStudyProgramForAcademic({
        curriculumYearId,
        semesterNumberId,
        studyProgramId,
      });
    },
    enabled:
      !!semesterNumberId &&
      !!studyProgramId &&
      !!curriculumYearId &&
      role === AKADEMIK,
  });
};
