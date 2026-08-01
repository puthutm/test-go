"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getSemesters = async (): Promise<
  ApiResponse<Semester[] | undefined>
> => {
  try {
    const response = await fetchApiDatareferensi(
      `/pmb/semesters/search?page_size=1000`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
