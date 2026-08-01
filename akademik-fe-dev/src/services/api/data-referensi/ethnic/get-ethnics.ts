"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getEthnics = async (): Promise<
  ApiResponse<Ethnic[] | undefined>
> => {
  try {
    const response = await fetchApiDatareferensi(
      "/biodata/ethnics/search?page_size=1000"
    );

    return response;
  } catch (error: any) {
    throw new Error(error?.response?.data?.message || "Something went wrong");
  }
};
