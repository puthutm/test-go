"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getCurriculumStudyProgramForAcademic = async ({
  curriculumYearId,
  semesterNumberId,
  studyProgramId,
}: {
  semesterNumberId: string;
  studyProgramId: string;
  curriculumYearId: string;
}) => {
  try {
    const params = new URLSearchParams();

    params.append("semester_number_id", semesterNumberId);
    params.append("study_program_id", studyProgramId);
    params.append("curriculum_year_id", curriculumYearId);
    const response = await fetchApi(
      `/academic/curriculum/study-program-curriculums?${params.toString()}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
