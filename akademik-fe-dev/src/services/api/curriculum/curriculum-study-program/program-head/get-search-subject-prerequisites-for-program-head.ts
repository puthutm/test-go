"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getSearchSubjectPrerequisitesForProgramHead = async ({
  curriculumYearId,
  semesterNumberId,
}: SubjectPrerequisitesParams): Promise<
  ApiResponse<SubjectPrerequisites[]>
> => {
  try {
    const params = new URLSearchParams();

    params.append("curriculum_year_id", curriculumYearId);
    params.append("semester_number_id", semesterNumberId);

    const response = await fetchApi(
      `/program-head/curriculum/academic-period/study-program-curriculums/search/subject?${params.toString()}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
