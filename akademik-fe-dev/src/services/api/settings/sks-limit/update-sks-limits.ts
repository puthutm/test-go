"use server";

import { revalidatePath } from "next/cache";
import { CreditLimitFormType } from "@/lib/validations/settings/credit-limit";
import { fetchApi } from "@/lib/utils/fetch-server";

export const editSksLimit = async (
  sksLimitId: string,
  payload: CreditLimitFormType
) => {
  const reqBody: ISksLimitForm = {
          ips_min: payload.ips_min,
      ips_max:payload.ips_max,
      sks_limit:payload.sks_limit,
    };

  try {
    const response = await fetchApi(`/academic/setting/sks-limits/${sksLimitId}`, {
      method: "PUT",
      body: JSON.stringify(reqBody),
      headers: {
        "Content-Type": "application/json",
      },
    });

    revalidatePath("/settings/credit-limit", "page");

    return response;
  } catch (error: any) {
    throw new Error(error);
  }
};