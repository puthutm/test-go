"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getBiodataStudent = async (): Promise<
  ApiResponse<BiodataStudent>
> => {
  try {
    const response = await fetchApi("/student/biodata/biodatas");

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
