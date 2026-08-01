"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const deleteGradeScale = async (
  gradeScaleId: string
): Promise<ApiResponse<Subject>> => {
  try {
    const response = await fetchApi(`/academic/setting/value-scales/${gradeScaleId}`, {
      method: "DELETE",
    });

    revalidatePath("/settings/grade-scale", "page");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
