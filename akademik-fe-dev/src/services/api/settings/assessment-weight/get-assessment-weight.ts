"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getAssessmentWeight = async () => {
  try {
    const response = await fetchApi("/academic/setting/assessment-weight");

    return response;
  } catch (error: any) {
    throw new Error(error.message || "Internal Server Error");
  }
};
