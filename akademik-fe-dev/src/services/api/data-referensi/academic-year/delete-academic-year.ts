"use server";

import { revalidatePath } from "next/cache";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const deleteAcademicYear = async (id: string) => {
  try {
    const response = await fetchApiDatareferensi(`/pmb/academic-years/${id}`, {
      method: "DELETE",
    });

    revalidatePath("/settings/academic-year", "page");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Internal Server Error"
    );
  }
};
