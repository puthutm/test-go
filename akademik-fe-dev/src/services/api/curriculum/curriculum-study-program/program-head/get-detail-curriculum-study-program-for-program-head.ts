"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getCurriculumStudyProgramByIdForProgramHead = async ({
  curriculumStudyProgramId,
}: {
  curriculumStudyProgramId: string;
}): Promise<ApiResponse<CurriculumStudyProgramTable>> => {
  try {
    const response = await fetchApi(
      `/program-head/curriculum/academic-period/study-program-curriculums/${curriculumStudyProgramId}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
