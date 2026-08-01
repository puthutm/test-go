"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getDetailGradeScale = async (GradeScaleId:string): Promise<
  ApiResponse<IGradeScale | undefined>
> => {
  try {
    const response = await fetchApi(
      `/academic/setting/value-scales/${GradeScaleId}`
    );

    return response;
  } catch (error: any) {
    throw new Error(error);
  }
};
