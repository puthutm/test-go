"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";
import { FormKrsStatusSchemaType } from "@/lib/validations/lecturer/update-status-reject-approve-krs";

export const updateKrsStatus = async (
  krsItemId: string,
  payload: FormKrsStatusSchemaType
) => {
  try {
    const response = await fetchApi(
      `/lecturer/lectures/krs-requests/${krsItemId}`,
      {
        method: "PUT",
        body: JSON.stringify(payload),
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    revalidatePath("/settings/credit-limit", "page");

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR CREATE SUBJECT");
    throw new Error(error);
  }
};
