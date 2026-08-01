"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const deleteClassScheduleSubDetail = async (
  classId: string,
  class_schedule_id: string
): Promise<ApiResponse<[]>> => {
  try {
    const response = await fetchApi(
      `/lecturer/academic/class-schedules/${classId}/class-schedules/${class_schedule_id}`,
      {
        method: "DELETE",
      }
    );

    revalidatePath(`/academic/college-class/${classId}/detail`, "page");
    revalidatePath(
      `/academic/college-class/${classId}/detail/${class_schedule_id}/detail-class-schedule`,
      "page"
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
