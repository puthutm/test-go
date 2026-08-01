"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const importStudent = async (payload: FormData) => {
  try {
    const response = await fetchApi("/academic/portal/students/bulk", {
      method: "POST",
      body: payload,
    });

    revalidatePath("/portal/students", "page");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Internal Server Error"
    );
  }
};
