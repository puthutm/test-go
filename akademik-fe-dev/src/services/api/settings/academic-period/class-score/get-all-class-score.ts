"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

interface Args {
  academicPeriodId: string;
  classId: string;
  queryParam: QueryParam;
}

export const getAllClassScore = async ({
  academicPeriodId,
  classId,
  queryParam,
}: Args): Promise<ApiResponse<PaginationData<ClassScore>> | undefined> => {
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

  try {
    const response = await fetchApi(
      `/academic/setting/academic-period/${academicPeriodId}/classes/${classId}/class-score?${params.toString()}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
