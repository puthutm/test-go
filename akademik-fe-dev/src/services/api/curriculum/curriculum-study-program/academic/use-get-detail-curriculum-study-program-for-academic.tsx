"use client";

import { useQuery } from "@tanstack/react-query";
import { getCurriculumStudyProgramById } from "./get-detail-curriculum-study-program-for-academic";
import { AKADEMIK } from "@/lib/constants/role";

export const useGetDetailCurriculumStudyProgramForAcademic = ({
  curriculumStudyProgramId,
  role,
}: {
  curriculumStudyProgramId: string;
  role: string;
}) => {
  return useQuery({
    queryKey: ["detail-curriculum-study-program", curriculumStudyProgramId],
    queryFn: async () =>
      await getCurriculumStudyProgramById({ curriculumStudyProgramId }),
    enabled: !!curriculumStudyProgramId && role === AKADEMIK,
  });
};
