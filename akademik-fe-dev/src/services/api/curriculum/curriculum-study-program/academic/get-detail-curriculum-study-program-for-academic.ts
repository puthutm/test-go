"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getCurriculumStudyProgramById = async ({
  curriculumStudyProgramId,
}: {
  curriculumStudyProgramId: string;
}): Promise<ApiResponse<CurriculumStudyProgramTable>> => {
  try {
    const response = await fetchApi(
      `/academic/curriculum/study-program-curriculums/${curriculumStudyProgramId}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
