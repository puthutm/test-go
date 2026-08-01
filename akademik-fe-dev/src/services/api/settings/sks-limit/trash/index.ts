
"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getTrashSksLimit = async (
  queryParam: IQueryParamsSksLimits
): Promise<ApiResponse<PaginationData<ISksLimit>>> => {
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
  try {
    const xhr = await fetchApi(
      `/academic/setting/sks-limits/trash?${params.toString()}`,
      {
        method: "GET",
      }
    );
    return xhr;
  } catch (err: any) {
    throw new Error(err.message || "gagal get data sampah batas sks");
  }
};
