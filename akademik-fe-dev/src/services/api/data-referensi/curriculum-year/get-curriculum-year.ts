"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getCurriculumYear = async (
  params: QueryParamDataRefensi
): Promise<ApiResponse<CurriculumYear[] | undefined>> => {
  try {
    const response = await fetchApiDatareferensi(
      `/academic/curriculum-years/search?page=${params.page}&page_size=${params.page_size}&filter=${params.filter}`
    );

    return response;
  } catch (error: any) {
    throw new Error(error);
  }
};
