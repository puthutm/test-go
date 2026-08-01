"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getDetailSksLimits = async (SksLimitsId:string): Promise<
  ApiResponse<ISksLimit | undefined>
> => {
  try {
    const response = await fetchApi(
      `/academic/setting/sks-limits/${SksLimitsId}`
    );

    return response;
  } catch (error: any) {
    throw new Error(error);
  }
};

