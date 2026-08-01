"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getFinalProjectProposalByProposalId = async (
  proposalId: string
): Promise<ApiResponse<FinalProjectProposal>> => {
  try {
    const response = await fetchApi(
      `/student/student/final-project-proposals/${proposalId}`
    );

    return response;
  } catch (error) {
    console.error("Error fetching final project proposal:", error);
    throw error;
  }
};
