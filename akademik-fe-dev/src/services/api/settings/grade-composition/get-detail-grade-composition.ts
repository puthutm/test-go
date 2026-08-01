"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getDetailGradeComposition = async (GradeCompositionId:string): Promise<
  ApiResponse<IGradeComposition | undefined>
> => {
  try {
    const response = await fetchApi(
      `/academic/setting/value-compositions/${GradeCompositionId}`
    );

    return response;
  } catch (error: any) {
    throw new Error(error);
  }
};
