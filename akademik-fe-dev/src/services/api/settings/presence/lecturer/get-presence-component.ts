"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getAllPresenceComponent = async ({
  academicPeriodId,
  subjectId,
  studyProgramId,
}: {
  academicPeriodId: string;
  subjectId: string;
  studyProgramId: string;
}): Promise<ApiResponse<PresenceComponent> | undefined> => {
  try {
    const params = new URLSearchParams();

    params.append("study_program_id", studyProgramId);

    const response = await fetchApi(
      `/lecturer/academic/presence/academic-periods/${academicPeriodId}/subjects/${subjectId}/component?${params.toString()}`
    );

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error?.message ?? "Internal Server Error");
  }
};
