"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const generateClassSchedule = async ({
  classId,
}: {
  classId: string;
}) => {
  try {
    const response = await fetchApi(
      `/program-head/curriculum/academic-period/classes/${classId}/schedules/generate`
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
