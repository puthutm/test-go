"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const deleteGradeComposition = async (
  gradeCompositionId: string
): Promise<ApiResponse<IGradeComposition>> => {
  try {
    const response = await fetchApi(`/academic/setting/value-compositions/${gradeCompositionId}`, {
      method: "DELETE",
    });

    revalidatePath("/settings/grade-composition", "page");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
