"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const deleteSksLimit = async (
  id: string
): Promise<ApiResponse<ISksLimit>> => {
  try {
    const response = await fetchApi(`/academic/setting/sks-limits/${id}`, {
      method: "DELETE",
    });

    revalidatePath("/settings/credit-limit", "page");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
