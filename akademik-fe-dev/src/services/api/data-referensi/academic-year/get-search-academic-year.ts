"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getSearchAcademicYears = async (): Promise<
  ApiResponse<AcademicYear[] | undefined>
> => {
  try {
    const response = await fetchApiDatareferensi(
      `/pmb/academic-years/search?page_size=1000`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
