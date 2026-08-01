"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";
import { GradeScaleFormType } from "@/lib/validations/settings/grade-scale/form-grade-scale";

export const createGradeScale = async (payload: GradeScaleFormType) => {
    const reqBody:IFormGradeScale = {
      study_program_id:payload.study_program_id.value,
      grade_id:payload.grade_id?.value,
      description:payload.description,
      weight_value:Number(payload.weight_value),
      lower_value:Number(payload.lower_value),
      upper_value:Number(payload.upper_value)
    };


  try {
    const response = await fetchApi("/academic/setting/value-scales", {
      method: "POST",
      body: JSON.stringify(reqBody),
      headers: {
        "Content-Type": "application/json",
      },
    });

    revalidatePath("/settings/grade-scale", "page");

    return response;
  } catch (error: any) {
    throw new Error(error);
  }
};
