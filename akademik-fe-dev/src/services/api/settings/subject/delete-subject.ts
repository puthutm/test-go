"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const deleteSubject = async (
  subjectId: string
): Promise<ApiResponse<Subject>> => {
  try {
    const response = await fetchApi(`/academic/setting/subjects/${subjectId}`, {
      method: "DELETE",
    });

    revalidatePath("/settings/subject", "page");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
