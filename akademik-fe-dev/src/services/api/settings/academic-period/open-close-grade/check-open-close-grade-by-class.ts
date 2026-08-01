"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

interface Args {
  academicPeriodId: string;
  classId: string;
}

export const checkOpenCloseGradeByClass = async ({
  academicPeriodId,
  classId,
}: Args): Promise<ApiResponse<OpenCloseGrade> | undefined> => {
  try {
    return await fetchApi(
      `/academic/setting/academic-period/${academicPeriodId}/classes/${classId}/class-score/open-close-values`
    );
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
