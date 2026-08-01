"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getSearchSemesterNumber = async (
  queryParams: QueryParamDataRefensi
): Promise<ApiResponse<SemesterNumberOptions[]>> => {
  const params = new URLSearchParams();

  params.append("page", queryParams.page.toString());

  if (queryParams.filter) params.append("filter", queryParams.filter);
  if (queryParams.page_size)
    params.append("page_size", queryParams.page_size.toString());
  if (queryParams.sort_by) params.append("sort_by", queryParams.sort_by);
  if (queryParams.sort_direction)
    params.append("sort_direction", queryParams.sort_direction);

  try {
    const response = await fetchApiDatareferensi(
      `/pmb/semester-numbers/search?${params.toString()}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
