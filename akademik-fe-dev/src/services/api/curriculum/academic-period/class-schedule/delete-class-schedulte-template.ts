"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const deleteClassScheduleTemplate = async ({
  classId,
  classScheduleTemplateId,
}: {
  classId: string;
  classScheduleTemplateId: string;
}) => {
  try {
    const response = await fetchApi(
      `/program-head/curriculum/academic-period/classes/${classId}/schedules/template/${classScheduleTemplateId}`,
      {
        method: "DELETE",
      }
    );

    revalidatePath(
      `/curriculum/academic-period/[academicPeriodId]/classes/[classId]/edit`,
      "page"
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
