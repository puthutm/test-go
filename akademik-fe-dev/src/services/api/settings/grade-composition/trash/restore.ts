"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const restoreGradeComposition = async (
  gradeCompositionId: string,
) => {
  try {
    const response = await fetchApi(`/academic/setting/value-compositions/trash/${gradeCompositionId}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
    });

    revalidatePath("/settings/grade-composition/trash", "page");

    return response;
  } catch (error: any) {
    throw new Error(error);
  }
};