"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";
import { FormClassAcademicPeriodDetailType } from "@/lib/validations/settings/academic-period/form-class";

export const updateClass = async (
  classId: string,
  payload: FormClassAcademicPeriodDetailType
) => {
  const reqBody = {
    code: payload.code,
    name: payload.name,
    academic_periode_id: payload.academic_periode_id?.value,
    subject_id: payload.subject_id?.value,
    capacity: payload.capacity,
    number_of_meeting: payload.number_of_meeting,
  };

  try {
    const response = await fetchApi(
      `/academic/setting/academic-period/classes/${classId}`,
      {
        method: "PUT",
        body: JSON.stringify(reqBody),
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    revalidatePath(
      `/settings/academic-period/[academicPeriodId]/classes`,
      "page"
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error?.message : "Something went wrong"
    );
  }
};
