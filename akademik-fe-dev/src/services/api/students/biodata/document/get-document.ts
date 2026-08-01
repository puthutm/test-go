"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getDocumentStudent = async (): Promise<
  ApiResponse<DocumentStudent>
> => {
  try {
    const response = await fetchApi("/student/biodata/documents");

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
