"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getOriginalEducationStudent = async (): Promise<
  ApiResponse<OriginalEducationStudent>
> => {
  try {
    const response = await fetchApi("/student/biodata/original-educations");

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
