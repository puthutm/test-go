"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getDetailClass = async (
  classId: string
): Promise<ApiResponse<Class>> => {
  try {
    const response = await fetchApi(
      `/academic/setting/academic-period/classes/${classId}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error?.message : "Something went wrong"
    );
  }
};
