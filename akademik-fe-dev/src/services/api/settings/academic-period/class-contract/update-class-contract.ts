"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";

export const updateClassContract = async (
  classId: string,
  payload: FormData
) => {
  try {
    const response = await fetchApi(
      `/academic/setting/academic-period/classes/${classId}/contract`,
      {
        method: "PUT",
        body: payload,
      }
    );

    revalidatePath(
      "/settings/academic-period/[academicPeriodId]/classes/[classId]/edit",
      "page"
    );

    return response;
  } catch (error) {
    console.log("Error update class contract", error);

    throw new Error("Something went wrong");
  }
};
