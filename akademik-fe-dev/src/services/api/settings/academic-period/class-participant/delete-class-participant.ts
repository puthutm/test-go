"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const deleteClassParticipant = async (
  classId: string,
  participantId: string
) => {
  try {
    const response = await fetchApi(
      `/academic/setting/academic-period/classes/${classId}/participants/${participantId}`,
      {
        method: "DELETE",
      }
    );

    revalidatePath(
      "/settings/academic-period/[academicPeriodId]/classes/[classId]/edit",
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
