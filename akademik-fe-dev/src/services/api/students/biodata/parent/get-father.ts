"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getFather = async (): Promise<
  ApiResponse<ParentStudent> | undefined
> => {
  try {
    const response = await fetchApi("/student/biodata/parents/father");

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
  }
};
