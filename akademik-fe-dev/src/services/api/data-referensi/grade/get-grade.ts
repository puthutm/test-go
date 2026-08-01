"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getGrade = async (): Promise<ApiResponse<GradeOptions[] | undefined>> => {
  try {
    const response = await fetchApiDatareferensi(
      `/academic/grades/search?page_size=1000`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
