"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

interface Args {
  academicPeriodId: string;
}

export const checkOpenCloseGradeByAcademicPeriod = async ({
  academicPeriodId,
}: Args): Promise<ApiResponse<OpenCloseGrade> | undefined> => {
  try {
    return await fetchApi(
      `/academic/setting/academic-period/${academicPeriodId}/classes/open-close-values`
    );
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
