"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getProvincesByCountryId = async (
  countryId: string
): Promise<ApiResponse<Province[] | undefined>> => {
  try {
    const response = await fetchApiDatareferensi(
      `/region/provinces/by-country/${countryId}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
