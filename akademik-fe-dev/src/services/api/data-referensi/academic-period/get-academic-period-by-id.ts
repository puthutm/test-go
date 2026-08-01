"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getAcademicPeriodById = async (
  id: string
): Promise<ApiResponse<AcademicPeriod>> => {
  try {
    const response = await fetchApiDatareferensi(`/pmb/academic-periods/${id}`);

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
