"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getSubjectById = async (
  subjectId: string
): Promise<ApiResponse<Subject>> => {
  try {
    const response = await fetchApi(`/academic/setting/subjects/${subjectId}`);

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
