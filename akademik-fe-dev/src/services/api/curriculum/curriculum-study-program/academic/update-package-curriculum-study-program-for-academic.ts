"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const updatePackageCurriculumStudyProgramForAcademic = async ({
  isPackage,
  semesterNumberId,
  studyProgramId,
  curriculumYearId,
}: {
  semesterNumberId: string;
  studyProgramId: string;
  isPackage: boolean;
  curriculumYearId: string;
}) => {
  try {
    const reqBody = {
      semester_number_id: semesterNumberId,
      study_program_id: studyProgramId,
      is_package: isPackage,
      curriculum_year_id: curriculumYearId,
    };
    const response = await fetchApi(
      `/academic/curriculum/study-program-curriculums/package`,
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
