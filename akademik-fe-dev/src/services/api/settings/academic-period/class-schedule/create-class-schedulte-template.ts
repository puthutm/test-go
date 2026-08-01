"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";
import { ClassScheduleFormTemplateSchemaType } from "@/lib/validations/settings/academic-period/form-class-schedule";
import { getTimes } from "@/lib/utils/format-date";

export const createClassScheduleTemplate = async ({
  classId,
  payload,
}: {
  classId: string;
  payload: ClassScheduleFormTemplateSchemaType;
}) => {
  try {
    const reqBody = {
      day_name: payload.day_name.value,
      start_time: getTimes(payload.start_time[0].toString()),
      end_time: getTimes(payload.end_time[0].toString()),
      type_of_meeting: payload.type_of_meeting.value,
    };

    const response = await fetchApi(
      `/academic/setting/academic-period/classes/${classId}/schedules/template`,
      {
        method: "POST",
        body: JSON.stringify(reqBody),
        headers: {
          "Content-Type": "application/json",
        },
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
