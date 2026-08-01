"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getReligions = async (): Promise<
  ApiResponse<Religion[] | undefined>
> => {
  try {
    const response = await fetchApiDatareferensi(
      "/biodata/religions/search?page_size=1000"
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
