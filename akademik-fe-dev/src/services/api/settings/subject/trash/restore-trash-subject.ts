"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const restoreSubject = async (
  subjectId: string
): Promise<ApiResponse<Subject>> => {
  try {
    const response = await fetchApi(
      `/academic/setting/subjects/trash/${subjectId}`,
      {
        method: "PUT",
      }
    );

    revalidatePath("/settings/subject", "page");
    revalidatePath("/settings/subject/trash", "page");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
