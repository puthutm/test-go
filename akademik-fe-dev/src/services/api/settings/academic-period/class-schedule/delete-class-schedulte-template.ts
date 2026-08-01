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
      `/academic/setting/academic-period/classes/${classId}/schedules/template/${classScheduleTemplateId}`,
      {
        method: "DELETE",
      }
    );

    revalidatePath(
      `/settings/academic-period/[academicPeriodId]/classes/[classId]`,
      "page"
    );
    revalidatePath(
      `/curriculum/academic-period/[academicPeriodId]/classes/[classId]`,
      "page"
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
