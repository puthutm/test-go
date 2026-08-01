"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getMother = async (): Promise<
  ApiResponse<ParentStudent> | undefined
> => {
  try {
    const response = await fetchApi("/student/biodata/parents/mother");

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
  }
};
