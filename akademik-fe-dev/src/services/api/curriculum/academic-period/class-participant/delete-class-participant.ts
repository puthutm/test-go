"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const deleteClassParticipantForProgramHead = async (
  classId: string,
  participantId: string
) => {
  try {
    const response = await fetchApi(
      `/program-head/curriculum/academic-period/classes/${classId}/participants/${participantId}`,
      {
        method: "DELETE",
      }
    );

    revalidatePath(
      "/curriculum/academic-period/[academicPeriodId]/classes/[classId]/edit",
      "page"
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
