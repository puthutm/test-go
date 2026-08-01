"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";
import { GradeCompositionFormType } from "@/lib/validations/settings/grade-composition/grade-composition";

export const createGradeComposition = async (
  payload: GradeCompositionFormType
) => {
  const reqBody = {
    academic_periode_id: payload.academic_periode_id.value,
    value_element_id: payload.value_element_id.value,
    percentage: Number(payload.percentage),
    is_passing_requirement: payload.is_passing_requirement,
  };

  try {
    const response = await fetchApi("/academic/setting/value-compositions", {
      method: "POST",
      body: JSON.stringify(reqBody),
      headers: {
        "Content-Type": "application/json",
      },
    });

    revalidatePath("/settings/grade-composition", "page");

    return response;
  } catch (error: any) {
    throw new Error(error);
  }
};
