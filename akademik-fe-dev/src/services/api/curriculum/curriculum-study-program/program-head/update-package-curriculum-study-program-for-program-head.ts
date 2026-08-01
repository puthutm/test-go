"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const updatePackageCurriculumStudyProgramForProgramHead = async ({
  isPackage,
  semesterNumberId,
  curriculumYearId,
}: {
  semesterNumberId: string;
  isPackage: boolean;
  curriculumYearId: string;
}) => {
  try {
    const reqBody = {
      semester_number_id: semesterNumberId,
      is_package: isPackage,
      curriculum_year_id: curriculumYearId,
    };
    const response = await fetchApi(
      `/program-head/curriculum/academic-period/study-program-curriculums/package`,
      {
        method: "PUT",
        body: JSON.stringify(reqBody),
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
