"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";
import { AssessmentWeightFormType } from "@/lib/validations/settings/assessment-weight/form-assessment-weight";

interface IFormAssessmentWeight {
  attitude_behavior_percentage: number;
  task_percentage: number;
  uts_percentage: number;
  uas_percentage: number;
}

export const updateAssessmentWeight = async (
  payload: AssessmentWeightFormType
) => {
  const reqBody: IFormAssessmentWeight = {
    attitude_behavior_percentage: Number(payload.attitude_behavior_percentage),
    task_percentage: Number(payload.task_percentage),
    uts_percentage: Number(payload.uts_percentage),
    uas_percentage: Number(payload.uas_percentage),
  };

  try {
    const response = await fetchApi("/academic/setting/assessment-weight", {
      method: "POST",
      body: JSON.stringify(reqBody),
      headers: {
        "Content-Type": "application/json",
      },
    });

    revalidatePath("/settings/assessment-weight", "page");

    return response;
  } catch (error: any) {
    throw new Error(
      error instanceof Error ? error?.message : "Something went wrong"
    );
  }
};
