"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getCurriculumYearById = async (
  id: string
): Promise<ApiResponse<CurriculumYear>> => {
  try {
    const response = await fetchApiDatareferensi(
      `/academic/curriculum-years/${id}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
