"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getCurriculumStudyProgramBySemesterNumberIdForProgramHead =
  async ({
    curriculumYearId,
    semesterNumberId,
  }: {
    semesterNumberId: string;
    curriculumYearId: string;
  }): Promise<ApiResponse<CurriculumStudyProgram[]>> => {
    try {
      const params = new URLSearchParams();

      params.append("semester_number_id", semesterNumberId);
      params.append("curriculum_year_id", curriculumYearId);
      const response = await fetchApi(
        `/program-head/curriculum/academic-period/study-program-curriculums?${params.toString()}`
      );

      return response;
    } catch (error) {
      throw new Error(
        error instanceof Error ? error.message : "Something went wrong"
      );
    }
  };
