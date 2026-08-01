"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";
import { FormSubjectSchemaType } from "@/lib/validations/settings/subject/form-subject";

export const editSubject = async (
  subjectId: string,
  payload: FormSubjectSchemaType
) => {
  const reqBody = {
    curriculum_year_id: payload.curriculum_year_id.value,
    study_program_id: payload.study_program_id.value,
    course_type_id: payload.course_type_id.value,
    course_group_id: payload.course_group_id.value,
    code: payload.code,
    name_id: payload.name_id,
    name_en: payload.name_en,
    face_to_face_sks: payload.face_to_face_sks,
    practicum_sks: payload.practicum_sks ?? 0,
    field_practice_sks: payload.field_practice_sks ?? 0,
    simulation_sks: payload?.simulation_sks ?? 0,
    field_of_studies_id: payload.field_of_studies_id.value,
    supporting_lecturer_id: payload.supporting_lecturer_id.map(
      (data) => data.value
    ),
    developer_rps_lecturer_id: payload.developer_rps_lecturer_id.map(
      (data) => data.value
    ),
    subject_coordinator_lecturer_id:
      payload.subject_coordinator_lecturer_id.map((data) => data.value),
    is_mku: payload.is_mku,
    is_sap: payload.is_sap,
    is_silabus: payload.is_silabus,
    is_teaching_material: payload.is_teaching_material,
    is_diktat: payload.is_diktat,
  };

  try {
    const response = await fetchApi(`/academic/setting/subjects/${subjectId}`, {
      method: "PUT",
      body: JSON.stringify(reqBody),
      headers: {
        "Content-Type": "application/json",
      },
    });

    revalidatePath("/settings/subject", "page");

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR CREATE SUBJECT");
    throw new Error(error);
  }
};
