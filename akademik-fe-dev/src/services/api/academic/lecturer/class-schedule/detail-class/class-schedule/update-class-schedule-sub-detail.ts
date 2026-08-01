"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const editClassScheduleSubDetail = async (
  idClass: string,
    idClassSchedule: string,
    payload: FormData
) => {
  try {
    const response = await fetchApi(`/lecturer/academic/class-schedules/${idClass}/class-schedules/${idClassSchedule}`, {
      method: "POST",
        body: payload,
    });

    revalidatePath(`/academic/college-class/${idClass}/detail/${idClassSchedule}/detail-class-schedule`, "page");
    revalidatePath(`/academic/college-class/${idClass}/detail/${idClassSchedule}/detail-class-schedule/edit`, "page");

    return response;
  } catch (error: any) {
    throw new Error(error);
  }
};