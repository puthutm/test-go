"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getSubjectsTrash = async (
  queryParam: QueryParam
): Promise<ApiResponse<PaginationData<Subject[]>>> => {
  const params = new URLSearchParams();

  if (queryParam.search) {
    params.append("search", queryParam.search);
  }

  if (queryParam.page) {
    params.append("page", String(queryParam.page));
  }

  if (queryParam.limit) {
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
      `/academic/setting/subjects/trash?${params.toString()}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
