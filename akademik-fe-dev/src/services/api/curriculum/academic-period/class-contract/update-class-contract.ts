"use server";

import { fetchApi } from "@/lib/utils/fetch-server";
import { revalidatePath } from "next/cache";

export const updateClassContractForProgramHead = async (
  classId: string,
  payload: FormData
) => {
  try {
    const response = await fetchApi(
      `/program-head/curriculum/academic-period/classes/${classId}/contract`,
      {
        method: "PUT",
        body: payload,
      }
    );

    revalidatePath(
      "/curriculum/academic-period/[academicPeriodId]/classes/[classId]/edit",
      "page"
    );

    return response;
  } catch (error) {
    console.log("Error update class contract", error);

    throw new Error("Something went wrong");
  }
};
