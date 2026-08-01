"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getTrashGradeComposition = async (
  queryParam: IQueryParamsGradeComposition
): Promise<ApiResponse<PaginationData<IGradeComposition>>> => {
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
  if (queryParam.value_element_id !== undefined) {
    params.append("value_element_id", String(queryParam.value_element_id));
  }

  if (queryParam.sort_by) {
    params.append("sort_by", queryParam.sort_by);
  }
  try {
    const xhr = await fetchApi(
      `/academic/setting/value-compositions/trash?${params.toString()}`,
      {
        method: "GET",
      }
    );
    return xhr;
  } catch (err: any) {
    throw new Error(err.message || "gagal get data sampah komposisi nilai");
  }
};