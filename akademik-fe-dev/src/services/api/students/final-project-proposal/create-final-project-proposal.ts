"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const createFinalProjectProposal = async (
  payload: FormData
): Promise<
  ApiResponse<
    Pick<FinalProjectProposal, "id" | "title_id" | "title_en" | "status">
  >
> => {
  try {
    const response = await fetchApi(
      "/student/student/final-project-proposals",
      {
        method: "POST",
        body: payload,
      }
    );

    return response;
  } catch (error) {
    console.error("Error creating final project proposal:", error);
    throw error;
  }
};
