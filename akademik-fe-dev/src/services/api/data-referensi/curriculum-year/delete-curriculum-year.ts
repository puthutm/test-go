"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";
import { revalidatePath } from "next/cache";

export const deleteCurriculumYear = async (
  id: string
): Promise<ApiResponse<CurriculumYear>> => {
  try {
    const response = await fetchApiDatareferensi(
      `/academic/curriculum-years/${id}`,
      {
        method: "DELETE",
      }
    );

    revalidatePath("/academic/curriculum-years", "page");

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
