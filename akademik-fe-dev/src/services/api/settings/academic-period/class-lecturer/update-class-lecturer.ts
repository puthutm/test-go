"use server";

import { fetchApi } from "@/lib/utils/fetch-server";
import { FormClassLecturerSchemaType } from "@/lib/validations/settings/academic-period/form-class-lecturer";
import { revalidatePath } from "next/cache";

export const updateClassLecturer = async (
  classId: string,
  classLecturerId: string,
  payload: FormClassLecturerSchemaType
) => {
  const reqBody = {
    lecturer_id: payload.lecturer_id,
    subtitute_lecturer_id: payload.subtitute_lecturer_id,
  };
  try {
    const response = await fetchApi(
      `/academic/setting/academic-period/classes/${classId}/lecturers/${classLecturerId}`,
      {
        method: "PUT",
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
      error instanceof Error ? error?.message : "Something went wrong"
    );
  }
};
