"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getEducationalLevels = async (): Promise<
  ApiResponse<EducationalLevel[] | undefined>
> => {
  try {
    const response = await fetchApiDatareferensi(
      "/education/educational-levels/search?page_size=1000"
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
