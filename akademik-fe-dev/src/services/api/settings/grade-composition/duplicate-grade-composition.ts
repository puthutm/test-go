"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

interface Payload {
  academicPeriodIdSource: string;
  academicPeriodIdTarget: string;
  isOverWrite: boolean;
}

export const duplicateGradeComposition = async ({
  academicPeriodIdSource,
  academicPeriodIdTarget,
  isOverWrite,
}: Payload) => {
  try {
    const reqBody = {
      academic_period_id_source: academicPeriodIdSource,
      academic_period_id_target: academicPeriodIdTarget,
      is_overwrite: isOverWrite,
    };
    const response = await fetchApi(
      "/academic/setting/value-compositions/duplicate",
      {
        method: "POST",
        body: JSON.stringify(reqBody),
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    revalidatePath("/settings/grade-composition", "page");

    return response;
  } catch (error: any) {
    throw new Error(error);
  }
};
