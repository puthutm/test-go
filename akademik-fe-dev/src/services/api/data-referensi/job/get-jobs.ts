"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getJobs = async (): Promise<ApiResponse<Job[] | undefined>> => {
  try {
    const response = await fetchApiDatareferensi(
      "/biodata/jobs/search?page_size=1000"
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
