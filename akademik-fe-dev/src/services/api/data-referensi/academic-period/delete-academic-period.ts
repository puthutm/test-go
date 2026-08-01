"use server";

import { revalidatePath } from "next/cache";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const deleteAcademicPeriod = async (
  id: string
): Promise<ApiResponse<null>> => {
  try {
    const response = await fetchApiDatareferensi(
      `/pmb/academic-periods/${id}`,
      {
        method: "DELETE",
      }
    );

    revalidatePath("/settings/academic-period", "page");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Internal Server Error"
    );
  }
};
