"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getAddressStudent = async (): Promise<
  ApiResponse<AddressStudent>
> => {
  try {
    const response = await fetchApi("/student/biodata/addresses");

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
