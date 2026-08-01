"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getDistrictsByCityId = async (
  cityId: string
): Promise<ApiResponse<District[] | undefined>> => {
  try {
    const response = await fetchApiDatareferensi(
      `/region/districts/by-city/${cityId}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
