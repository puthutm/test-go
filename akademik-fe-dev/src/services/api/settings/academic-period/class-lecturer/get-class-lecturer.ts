"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getClassLecturerByClassId = async (
  classId: string
): Promise<ApiResponse<ClassLecturer>> => {
  try {
    const response = await fetchApi(
      `/academic/setting/academic-period/classes/${classId}/lecturers`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
