"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const restoreGradeScale = async (
  gradeScaleId: string,
) => {
  try {
    const response = await fetchApi(`/academic/setting/value-scales/trash/${gradeScaleId}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
    });

    revalidatePath("/settings/grade-scale/trash", "page");

    return response;
  } catch (error: any) {
    throw new Error(error);
  }
};