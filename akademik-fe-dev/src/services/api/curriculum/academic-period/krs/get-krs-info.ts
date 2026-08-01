"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getKrsInfo = async (): Promise<ApiResponse<InfoKrs>> => {
  try {
    const response = await fetchApi(`/student/academic/filling-krs/info`);

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error?.message : "Something went wrong"
    );
  }
};
