"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";
import { CurriculumStudyProgramForAcademicSchemaType } from "@/lib/validations/curriculum/form-curriculum-study-program";

export const updateCurriculumStudyProgramForAcademic = async ({
  id,
  payload,
}: {
  id: string;
  payload: CurriculumStudyProgramForAcademicSchemaType;
}) => {
  try {
    const reqBody = {
      curriculum_year_id: payload.curriculum_year_id,
      limit_grade_id: payload.limit_grade_id,
      semester_number_id: payload.semester_number_id,
      study_program_id: payload.study_program_id,
      subject_id: payload.subject_id.value,
      is_mandatory: payload.is_mandatory,
      subject_prerequisites:
        payload.subject_prerequisites?.map((data) => data.value) ?? null,
      field_study_concentration_id:
        payload.field_study_concentration_id?.value ?? null,
    };

    const response = await fetchApi(
      `/academic/curriculum/study-program-curriculums/${id}`,
      {
        method: "PUT",
        body: JSON.stringify(reqBody),
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    revalidatePath("/curriculum/curriculum-study-program", "page");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Internal Server Error"
    );
  }
};
