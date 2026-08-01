"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export interface SubjectParams {
  curriculum_year_id: string;
  study_program_id: string;
}

type SubjectType = Pick<Subject, "id" | "code" | "name_id">;

export const getSearchSubject = async (
  queryParams: SubjectParams
): Promise<ApiResponse<SubjectType[]>> => {
  const params = new URLSearchParams();

  params.append("curriculum_year_id", queryParams.curriculum_year_id);
  params.append("study_program_id", queryParams.study_program_id);

  try {
    const response = await fetchApi(
      `/academic/setting/subjects/search?${params.toString()}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
