"use server";

import { fetchApi } from "@/lib/utils/fetch-server";
import { FormPresenceSchemaType } from "@/lib/validations/settings/presence-student/create-presence-student-validation";
import { revalidatePath } from "next/cache";

export const createPresenceStudent = async (
  payload: FormPresenceSchemaType
) => {
  const reqBody = {
    ...payload,
    academic_periode_id: payload.academic_periode_id.value,
    study_program_id: payload.study_program_id.value,
  };
  try {
    const response = await fetchApi(`/academic/setting/presence/student`, {
      method: "POST",
      body: JSON.stringify(reqBody),
      headers: {
        "Content-Type": "application/json",
      },
    });

    revalidatePath("/settings/presence-student");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error?.message : "Something went wrong"
    );
  }
};
