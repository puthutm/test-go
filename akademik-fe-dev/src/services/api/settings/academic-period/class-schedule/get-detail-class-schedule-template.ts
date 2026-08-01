"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getClassScheduleTemplateById = async ({
  classId,
  classScheduleTemplateId,
}: {
  classId: string;
  classScheduleTemplateId: string;
}): Promise<ApiResponse<ClassScheduleTemplate>> => {
  try {
    const response = await fetchApi(
      `/academic/setting/academic-period/classes/${classId}/schedules/template/${classScheduleTemplateId}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
