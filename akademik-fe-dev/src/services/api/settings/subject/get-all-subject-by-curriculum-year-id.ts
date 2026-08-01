"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

interface SubjectParams extends QueryParam {
  study_program_id?: string;
}

export const getSubjectsByCurriculumYearId = async ({
  queryParam,
  curriculumYearId,
}: {
  queryParam: SubjectParams;
  curriculumYearId: string;
}): Promise<ApiResponse<PaginationData<Subject[]>>> => {
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
    const response = await fetchApi(
      `/academic/setting/curriculum_year/${curriculumYearId}/subjects?${params.toString()}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
