"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getBloodTypes = async (): Promise<
  ApiResponse<BloodType[] | undefined>
> => {
  try {
    const response = await fetchApiDatareferensi(
      "/biodata/blood-types/search?page_size=1000"
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
