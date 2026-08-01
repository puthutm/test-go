"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";
import { FormPresenceComponentSchemaType } from "@/lib/validations/settings/presence-student/create-presence-student-for-lecturer-validation";

export const createOrUpdatePresenceComponentStudent = async ({
  academicPeriodId,
  subjectId,
  payload,
}: {
  academicPeriodId: string;
  subjectId: string;
  payload: FormPresenceComponentSchemaType;
}) => {
  try {
    const response = await fetchApi(
      `/lecturer/academic/presence/academic-periods/${academicPeriodId}/subjects/${subjectId}/class`,
      {
        method: "POST",
        body: JSON.stringify(payload),
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    revalidatePath("/settings/presence/[subjectId]", "page");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error?.message : "Something went wrong"
    );
  }
};
