"use client";

import { useQuery } from "@tanstack/react-query";

import { getCurriculumStudyProgramBySemesterNumberIdForProgramHead } from "./get-curriculum-study-program-by-semester-number-id-for-program-head";
import { KAPRODI } from "@/lib/constants/role";

export const useGetCurriculumStudyProgramBySemesterIdForProgramHead = ({
  curriculumYearId,
  semesterNumberId,
  role,
}: {
  semesterNumberId: string;
  curriculumYearId: string;
  role: string;
}) => {
  return useQuery({
    queryKey: [
      "curriculum-study-program-for-program-head",
      semesterNumberId,
      curriculumYearId,
    ],
    queryFn: async () => {
      return await getCurriculumStudyProgramBySemesterNumberIdForProgramHead({
        curriculumYearId,
        semesterNumberId,
      });
    },
    enabled: !!curriculumYearId && !!semesterNumberId && role === KAPRODI,
  });
};
