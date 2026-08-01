"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getCompletenessStudent = async (): Promise<
  ApiResponse<CompletenessStudent> | undefined
> => {
  try {
    const response = await fetchApi("/student/biodata/completeness");

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
