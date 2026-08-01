"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

type SubjectType = Pick<Subject, "id" | "code" | "name_id">;

export type SubjectProgramHeadParams = {
  curriculum_year_id: string;
  search?: string;
};

export const getSearchSubjectForProgramHead = async (
  queryParams: SubjectProgramHeadParams
): Promise<ApiResponse<SubjectType[]>> => {
  const params = new URLSearchParams();

  params.append("curriculum_year_id", queryParams.curriculum_year_id);
  if (queryParams.search) {
    params.append("search", queryParams?.search);
  }
  try {
    const response = await fetchApi(
      `/academic/setting/subjects/search/program-head?${params.toString()}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
