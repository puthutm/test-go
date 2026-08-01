"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getFinalProjectProposals = async (): Promise<
  ApiResponse<FinalProjectProposalStudent[]> | undefined
> => {
  try {
    const response = await fetchApi("/student/student/final-project-proposals");

    return response;
  } catch (error) {
    if (error instanceof Error) {
      console.log(error?.message, "<<<< ERROR");
      throw new Error(error?.message);
    }
  }
};
