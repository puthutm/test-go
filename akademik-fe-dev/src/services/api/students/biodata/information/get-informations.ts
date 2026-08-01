"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getInformationStudent = async (): Promise<
  ApiResponse<InformationStudent>
> => {
  try {
    const response = await fetchApi("/student/biodata/informations");

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
