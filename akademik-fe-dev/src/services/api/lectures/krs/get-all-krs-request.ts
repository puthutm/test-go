"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

interface QueryParamsKrs extends QueryParam {
  academic_periode_id?: string;
}

export const getAllKRSRequest = async (
  queryParams: QueryParamsKrs
): Promise<ApiResponse<PaginationData<KRS[]>>> => {
  try {
    const params = new URLSearchParams();

    if (queryParams.search) {
      params.append("search", queryParams.search);
    }

    if (queryParams.page !== undefined) {
      params.append("page", String(queryParams.page));
    }

    if (queryParams.limit !== undefined) {
      params.append("limit", String(queryParams.limit));
    }

    if (queryParams.sort_by) {
      params.append("sort_by", queryParams.sort_by);
    }

    if (queryParams.sort_direction) {
      params.append("sort_direction", queryParams.sort_direction);
    }

    if (queryParams.academic_periode_id) {
      params.append("academic_periode_id", queryParams.academic_periode_id);
    }

    const response = await fetchApi(
      `/lecturer/lectures/krs-requests?${params.toString()}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
