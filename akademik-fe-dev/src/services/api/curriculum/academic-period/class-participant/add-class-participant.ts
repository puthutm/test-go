"use server";

import { fetchApi } from "@/lib/utils/fetch-server";
import { FormClassParticipantSchemaType } from "@/lib/validations/settings/academic-period/form-class-participant";
import { revalidatePath } from "next/cache";

export const addClassParticipantForProgramHead = async (
  classId: string,
  payload: FormClassParticipantSchemaType
) => {
  try {
    const response = await fetchApi(
      `/program-head/curriculum/academic-period/classes/${classId}/participants`,
      {
        method: "POST",
        body: JSON.stringify(payload),
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    revalidatePath(
      "/curriculum/academic-period/[academicPeriodId]/classes/[classId]/edit",
      "page"
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
