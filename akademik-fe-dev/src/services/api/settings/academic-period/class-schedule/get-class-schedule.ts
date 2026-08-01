"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getClassSchedule = async ({
  classId,
  queryParam,
}: {
  classId: string;
  queryParam: QueryParam;
}): Promise<ApiResponse<PaginationData<ClassSchedule>>> => {
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

    const response = await fetchApi(
      `/academic/setting/academic-period/classes/${classId}/schedules-as-of-date?${params.toString()}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
