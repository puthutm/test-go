"use client";

import { useQuery } from "@tanstack/react-query";
import { getSearchStudentByStudyProgram } from "./get-search-student-by-study-program";

export const useGetSearchStudentByStudyProgram = ({
  studyProgramId,
  search,
  limit = 200,
}: {
  studyProgramId: string;
  search?: string;
  limit?: number;
}) => {
  return useQuery({
    queryKey: ["get-search-student", studyProgramId, search, limit],
    queryFn: async () =>
      await getSearchStudentByStudyProgram({
        page: 1,
        study_program_id: studyProgramId,
        search: search,
      }),
    enabled: !!search,
  });
};
