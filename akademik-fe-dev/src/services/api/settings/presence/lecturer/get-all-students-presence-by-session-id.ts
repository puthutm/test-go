"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

interface PresenceStudentBySessionIdParam extends QueryParam {
  status?: string;
}

export const getAllStudentPresenceBySessionId = async (
  sessionId: string,
  queryParam: PresenceStudentBySessionIdParam
): Promise<ApiResponse<PaginationData<StudentPresenceSession>> | undefined> => {
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

    if (queryParam.status) {
      params.append("status", queryParam.status);
    }

    const response = await fetchApi(
      `/lecturer/academic/presence/students/sessions/${sessionId}?${params.toString()}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error?.message ?? "Internal Server Error");
  }
};
