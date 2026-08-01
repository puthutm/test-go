"use server";

import { revalidatePath } from "next/cache";

import { FormAddressSchemaType } from "@/lib/validations/students/biodata/form-address-schema";
import { fetchApi } from "@/lib/utils/fetch-server";

export const updateAddressStudent = async (
  payload: FormAddressSchemaType
): Promise<ApiResponse<AddressStudent>> => {
  const reqBody: any = {
    country_id: payload?.country_id?.value,
    province_id: payload?.province_id?.value,
    city_id: payload?.city_id?.value,
    district_id: payload?.district_id?.value,
    village_id: payload?.village_id?.value,
    address: payload.address || null,
    distance: payload.distance || null,
    postal_code: payload.postal_code || null,
    rt: payload.rt || null,
    rw: payload.rw || null,
  };

  try {
    const response = await fetchApi("/student/biodata/addresses", {
      method: "PUT",
      body: JSON.stringify(reqBody),
      headers: {
        "Content-Type": "application/json",
      },
    });

    revalidatePath("/student", "page");

    return response;
  } catch (error: any) {
    console.log(error?.message, "<<<< ERROR");
    throw new Error(error);
  }
};
