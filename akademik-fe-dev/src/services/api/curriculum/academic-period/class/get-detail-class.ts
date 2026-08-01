"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getDetailClassForProgramHead = async (
  classId: string
): Promise<ApiResponse<Class>> => {
  try {
    const response = await fetchApi(
      `/program-head/curriculum/academic-period/classes/${classId}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error?.message : "Something went wrong"
    );
  }
};
