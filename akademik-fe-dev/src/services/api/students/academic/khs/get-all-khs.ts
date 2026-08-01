"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getKhsStudent = async (): Promise<ApiResponse<Khs>> => {
  try {
    const response = await fetchApi("/student/academic/khs");

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
