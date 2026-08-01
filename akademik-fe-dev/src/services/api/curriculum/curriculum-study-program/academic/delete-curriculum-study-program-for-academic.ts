"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const deleteCurriculumStudyProgramForAcademic = async (
  curriculumStudyProgramId: string
) => {
  try {
    const response = await fetchApi(
      `/academic/curriculum/study-program-curriculums/${curriculumStudyProgramId}`,
      {
        method: "DELETE",
      }
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
