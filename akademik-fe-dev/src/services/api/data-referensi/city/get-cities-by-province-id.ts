"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getCitiesByProvinceId = async (
  provinceId: string
): Promise<ApiResponse<City[] | undefined>> => {
  try {
    const response = await fetchApiDatareferensi(
      `/region/cities/by-province/${provinceId}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
