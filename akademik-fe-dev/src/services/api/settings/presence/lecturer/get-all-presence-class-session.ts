"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getAllPresenceClassSession = async ({
  academicPeriodId,
  subjectId,
  classId,
}: {
  academicPeriodId: string;
  subjectId: string;
  classId: string;
}): Promise<ApiResponse<ClassPresenceSession[]> | undefined> => {
  try {
    const response = await fetchApi(
      `/lecturer/academic/presence/academic-periods/${academicPeriodId}/subjects/${subjectId}/class/${classId}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error?.message ?? "Internal Server Error");
  }
};
