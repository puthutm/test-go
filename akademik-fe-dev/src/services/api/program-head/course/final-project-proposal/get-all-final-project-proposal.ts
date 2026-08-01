"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

interface FinalProjectProposalParams extends QueryParam {
  study_program_id?: string;
  status?: string;
  academic_period_id?: string;
}

export const getFinalProjectProposalForProgramHead = async (
  queryParam: FinalProjectProposalParams
) => {
  const params = new URLSearchParams();

  if (queryParam.study_program_id) {
    params.append("study_program_id", queryParam.study_program_id);
  }

  if (queryParam.status) {
    params.append("status", queryParam.status);
  }

  if (queryParam.academic_period_id) {
    params.append("academic_period_id", queryParam.academic_period_id);
  }

  if (queryParam.page !== undefined) {
    params.append("page", String(queryParam.page));
  }

  if (queryParam.limit !== undefined) {
    params.append("limit", String(queryParam.limit));
  }

  try {
    const response = await fetchApi(
      `/program-head/course/final-project-proposals?${params.toString()}`
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
