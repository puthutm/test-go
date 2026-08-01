"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getCities = async (
  search?: string
): Promise<ApiResponse<City[] | undefined>> => {
  try {
    const response = await fetchApiDatareferensi(
      `/region/cities/search?page_size=1000&filter=${search}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
