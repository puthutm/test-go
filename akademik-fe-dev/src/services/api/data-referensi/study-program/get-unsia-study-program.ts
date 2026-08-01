"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getUnsiaStudyProgram = async (): Promise<
  ApiResponse<UnsiaStudyProgram[] | undefined>
> => {
  try {
    const response = await fetchApiDatareferensi(
      `/education/study-programs/search?page_size=1000&filter=unsia`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
