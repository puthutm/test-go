"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";
import { FormClassForProgramHeadType } from "@/lib/validations/curriculum/form-class-detail-schema";

export const updateClassForProgramHead = async (
  classId: string,
  payload: FormClassForProgramHeadType
) => {
  const reqBody = {
    code: payload.code,
    name: payload.name,
    academic_periode_id: payload.academic_period_id,
    subject_id: payload.subject_id?.value,
    capacity: payload.capacity,
    number_of_meeting: payload.number_of_meeting,
  };

  try {
    const response = await fetchApi(
      `/program-head/curriculum/academic-period/classes/${classId}`,
      {
        method: "PUT",
        body: JSON.stringify(reqBody),
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    revalidatePath(
      `/settings/academic-period/[academicPeriodId}/classes`,
      "page"
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error?.message : "Something went wrong"
    );
  }
};
