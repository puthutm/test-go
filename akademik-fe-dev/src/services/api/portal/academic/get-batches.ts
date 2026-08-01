"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getBatches = async (): Promise<ApiResponse<Batch[]>> => {
  try {
    const response = await fetchApi(`/academic/portal/students/batches`);

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error
        ? error.message
        : "An error occurred while fetching data"
    );
  }
};
