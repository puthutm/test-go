"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getVillagesByDistrictId = async (
  districtId: string
): Promise<ApiResponse<Village[] | undefined>> => {
  try {
    const response = await fetchApiDatareferensi(
      `/region/villages/by-district/${districtId}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
