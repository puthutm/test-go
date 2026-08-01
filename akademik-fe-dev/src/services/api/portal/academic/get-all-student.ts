import { fetchApi } from "@/lib/utils/fetch-server";

export const getAllStudentsForAcademic = async (
  queryParam: QueryParam
): Promise<ApiResponse<PaginationData<PortalStudentAcademic>>> => {
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
      `/academic/portal/students?${params.toString()}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error
        ? error.message
        : "An error occurred while fetching data"
    );
  }
};
