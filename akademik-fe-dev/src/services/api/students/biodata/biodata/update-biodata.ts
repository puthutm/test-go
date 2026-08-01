"use server";
import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";
import { formatDateNumeric } from "@/lib/utils/format-date";
import { FormBiodataSchemaType } from "@/lib/validations/students/biodata/form-biodata-schema";

export const updateBiodataStudent = async (
  payload: FormBiodataSchemaType
): Promise<ApiResponse<BiodataStudent>> => {
  const reqBody: FormBiodataStudent = {
    birth_date: formatDateNumeric(payload.birth_date[0].toString()),
    birth_place_id: payload.birth_place_id.value,
    ethnic_id: payload.ethnic_id?.value as string,
    gender: payload.gender.value,
    name: payload.name,
    nik: payload.nik,
    no_kk: payload.no_kk,
    religion_id: payload.religion_id.value,
    status_id: payload.status_id?.value as string,
    back_degree: payload.back_degree || null,
    blood_type_id: payload.blood_type_id?.value || null,
    height: payload.height ? parseInt(payload.height) : null,
    weight: payload.weight ? parseInt(payload.weight) : null,
  };

  try {
    const response = await fetchApi("/student/biodata/biodatas", {
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
