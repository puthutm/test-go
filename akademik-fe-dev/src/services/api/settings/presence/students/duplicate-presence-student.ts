"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";
import { FormDuplicatePresenceSchemaType } from "@/lib/validations/settings/presence-student/duplicate-presence-student-validatoin";

export const duplicatePresenceStudent = async (
  payload: FormDuplicatePresenceSchemaType
) => {
  const reqBody = {
    academic_periode_id: payload.academic_periode_id.value,
    academic_periode_id_target: payload.academic_periode_id_target.value,
    study_program_id: payload.study_program_id.value,
  };

  try {
    const response = await fetchApi(
      "/academic/setting/presence/student/duplicate",
      {
        method: "POST",
        body: JSON.stringify(reqBody),
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    revalidatePath("/settings/presence-student");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error?.message : "Something went wrong"
    );
  }
};
