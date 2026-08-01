"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getTransportations = async (): Promise<
  ApiResponse<Transportation[] | undefined>
> => {
  try {
    const response = await fetchApiDatareferensi(
      "/biodata/transportations/search?page_size=1000"
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
