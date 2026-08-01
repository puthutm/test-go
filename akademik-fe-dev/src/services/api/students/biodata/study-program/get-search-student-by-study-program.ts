"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

interface QueryParamStudent extends QueryParam {
  study_program_id: string;
}

interface Student {
  student_id: string;
  student_nim: string;
  student_name: string;
}

export const getSearchStudentByStudyProgram = async (
  queryParam: QueryParamStudent
): Promise<ApiResponse<PaginationData<Student>>> => {
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

  if (queryParam.study_program_id) {
    params.append("study_program_id", queryParam.study_program_id);
  }

  try {
    const response = await fetchApi(
      `/student/biodata/study-program?${params.toString()}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
