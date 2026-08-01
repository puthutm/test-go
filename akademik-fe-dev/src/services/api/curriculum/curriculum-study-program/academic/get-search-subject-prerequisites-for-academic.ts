"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getSearchSubjectPrerequisitesForAcademic = async ({
  curriculumYearId,
  semesterNumberId,
  studyProgramId,
}: SubjectPrerequisitesParams): Promise<
  ApiResponse<SubjectPrerequisites[]>
> => {
  try {
    const params = new URLSearchParams();

    params.append("curriculum_year_id", curriculumYearId);
    params.append("semester_number_id", semesterNumberId);
    params.append("study_program_id", studyProgramId as string);

    const response = await fetchApi(
      `/academic/curriculum/study-program-curriculums/search/subject?${params.toString()}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
