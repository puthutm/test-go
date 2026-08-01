"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

interface QueryParamClass extends QueryParam {
  academic_periode_id: string;
  study_program_id: string;
}

export const getAllClassByAcademicPeriodIdForAcademic = async (
  queryParam: QueryParamClass
): Promise<ApiResponse<PaginationData<Class[]>>> => {
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

  if (queryParam.academic_periode_id) {
    params.append("academic_periode_id", queryParam.academic_periode_id);
  }

  if (queryParam.study_program_id) {
    params.append("study_program_id", queryParam.study_program_id);
  }

  try {
    const response = await fetchApi(
      `/academic/setting/academic-period/classes?${params.toString()}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
