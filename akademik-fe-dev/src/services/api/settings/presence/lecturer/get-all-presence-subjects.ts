"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

interface PresenceSubjectParam extends QueryParam {
  academic_periode_id: string;
}

export const getAllPresenceSubject = async (
  queryParam: PresenceSubjectParam
): Promise<ApiResponse<PaginationData<SubjectsPresence>>> => {
  try {
    const params = new URLSearchParams();

    if (queryParam.search) {
      params.append("search", queryParam.search);
    }

    if (queryParam.page !== undefined) {
      params.append("page", String(queryParam.page));
    }

    if (queryParam.limit !== undefined) {
      params.append("limit", String(queryParam.limit));
    }

    if (queryParam.sort_by) {
      params.append("sort_by", queryParam.sort_by);
    }

    if (queryParam.sort_direction) {
      params.append("sort_direction", queryParam.sort_direction);
    }

    if (queryParam.academic_periode_id) {
      params.append("academic_period_id", queryParam.academic_periode_id);
    }

    const response = await fetchApi(
      `/lecturer/academic/presence?${params.toString()}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error?.message ?? "Internal Server Error");
  }
};
