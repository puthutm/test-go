"use client";

import { useQuery } from "@tanstack/react-query";

import { getCurriculumStudyProgramByIdForProgramHead } from "./get-detail-curriculum-study-program-for-program-head";
import { KAPRODI } from "@/lib/constants/role";

export const useGetDetailCurriculumStudyProgramForProgramHead = ({
  curriculumStudyProgramId,
  role,
}: {
  curriculumStudyProgramId: string;
  role: string;
}) => {
  return useQuery({
    queryKey: [
      "detail-curriculum-study-program-program-head",
      curriculumStudyProgramId,
    ],
    queryFn: async () =>
      await getCurriculumStudyProgramByIdForProgramHead({
        curriculumStudyProgramId,
      }),
    enabled: !!curriculumStudyProgramId && role === KAPRODI,
  });
};
