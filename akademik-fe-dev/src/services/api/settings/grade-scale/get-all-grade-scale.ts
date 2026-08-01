"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getGradeScale = async (
  queryParam: IQueryParamsGradeScale
): Promise<ApiResponse<PaginationData<IGradeScale>>> => {
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
  if (queryParam.study_program_id !== undefined) {
    params.append("study_program_id", String(queryParam.study_program_id));
  }

  if (queryParam.sort_by) {
    params.append("sort_by", queryParam.sort_by);
  }
  try {
    const xhr = await fetchApi(
      `/academic/setting/value-scales?${params.toString()}`,
      {
        method: "GET",
      }
    );
    return xhr;
  } catch (err: any) {
    throw new Error(err.message || "gagal get data skala nilai");
  }
};
