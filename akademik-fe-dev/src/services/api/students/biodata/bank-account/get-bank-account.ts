"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getBankAccountStudent = async (): Promise<
  ApiResponse<BankAccountStudent>
> => {
  try {
    const response = await fetchApi("/student/biodata/bank-accounts");

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
