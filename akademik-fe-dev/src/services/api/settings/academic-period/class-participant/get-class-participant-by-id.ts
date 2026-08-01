"use server";

import { fetchApi } from "@/lib/utils/fetch-server";

export const getClassParticipantById = async (
  classId: string,
  participantId: string
): Promise<ApiResponse<ClassParticipant>> => {
  try {
    const response = await fetchApi(
      `/academic/setting/academic-period/classes/${classId}/participants/${participantId}`
    );

    return response;
  } catch (error) {
    console.log(
      error instanceof Error ? error.message : "Something went wrong",
      "<<<<< Error log"
    );

    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
