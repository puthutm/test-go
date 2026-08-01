"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getAllPresenceClassBySubjectId = async ({
  academicPeriodId,
  subjectId,
  queryParam,
}: {
  academicPeriodId: string;
  subjectId: string;
  queryParam: QueryParam;
}): Promise<ApiResponse<PaginationData<ClassPresence>> | undefined> => {
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

    const response = await fetchApi(
      `/lecturer/academic/presence/academic-periods/${academicPeriodId}/subjects/${subjectId}/class?${params.toString()}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error?.message ?? "Internal Server Error");
  }
};
