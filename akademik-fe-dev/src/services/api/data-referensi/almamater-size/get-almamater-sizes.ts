"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getAlmamaterSizes = async (): Promise<
  ApiResponse<AlmamaterSize[] | undefined>
> => {
  try {
    const response = await fetchApiDatareferensi(
      "/biodata/almamater-sizes/search?page_size=1000"
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
