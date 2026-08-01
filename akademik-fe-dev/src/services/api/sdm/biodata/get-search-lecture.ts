"use server";

import { fetchApiSdm } from "@/lib/utils/fetch-server";

interface BiodataOfLecture {
  id: string;
  user_id: string;
  name_of_user: string;
}

export interface LecturerParams extends QueryParam {
  study_program_id?: string;
}

export const getSearchLecturer = async (
  queryParam: LecturerParams
): Promise<ApiResponse<PaginationData<BiodataOfLecture>>> => {
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

  if (queryParam.study_program_id) {
    params.append("study_program_id", queryParam.study_program_id);
  }

  try {
    const response = await fetchApiSdm(
      `/biographical/biographi/admin/search?${params.toString()}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
