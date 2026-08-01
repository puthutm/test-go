"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getCountries = async (): Promise<
  ApiResponse<Country[] | undefined>
> => {
  try {
    const response = await fetchApiDatareferensi(
      `/region/countries/search?page_size=1000`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
