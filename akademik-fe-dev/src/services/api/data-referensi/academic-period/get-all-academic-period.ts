"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getAcademicPeriods = async (
  queryParam: QueryParamDataRefensi
): Promise<ApiResponse<PaginationDataReferensi<AcademicPeriod[]>>> => {
  const params = new URLSearchParams();

  if (queryParam.filter) {
    params.append("filter", queryParam.filter);
  }

  if (queryParam.page !== undefined) {
    params.append("page", String(queryParam.page));
  }

  if (queryParam.page_size !== undefined) {
    params.append("page_size", String(queryParam.page_size));
  }

  if (queryParam.sort_by) {
    params.append("sort_by", queryParam.sort_by);
  }

  if (queryParam.sort_direction) {
    params.append("sort_direction", queryParam.sort_direction);
  }

  try {
    const response = await fetchApiDatareferensi(
      `/pmb/academic-periods?${params.toString()}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
