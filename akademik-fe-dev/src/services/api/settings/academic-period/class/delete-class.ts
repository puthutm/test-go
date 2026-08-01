"use server";

import { fetchApi } from "@/lib/utils/fetch-server";
import { revalidatePath } from "next/cache";

export const deleteClass = async (classId: string) => {
  try {
    const response = await fetchApi(
      `/academic/setting/academic-period/classes/${classId}`,
      {
        method: "DELETE",
      }
    );

    revalidatePath(
      `/settings/academic-period/[academicPeriod]/classes/classes`,
      "page"
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error?.message : "Something went wrong"
    );
  }
};
