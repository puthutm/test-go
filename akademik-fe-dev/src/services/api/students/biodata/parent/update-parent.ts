"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";
import { formatDateNumeric } from "@/lib/utils/format-date";
import { FormParentSchemaType } from "@/lib/validations/students/biodata/form-parent-schema";

type UpdateArgs = {
  parentType: "father" | "mother";
  payload: FormParentSchemaType;
};

export const updateParent = async ({
  parentType,
  payload,
}: UpdateArgs): Promise<ApiResponse<FormParentStudent>> => {
  const reqBody: FormParentStudent = {
    name: payload.name,
    nik: payload.nik,
    email: payload?.email || null,
    phone: payload?.phone || null,
    phone2: payload?.phone2 || null,
    address: payload?.address || null,
    income: payload?.income || null,
    kinship: payload?.kinship || null,
    job_id: payload?.job_id?.value || null,
    birth_place_id: payload?.birth_place_id?.value || null,
    education_level_id: payload?.education_level_id?.value || null,
    status_kinship: payload?.status_kinship?.value || null,
    life_status: payload?.life_status?.value || null,
    birth_date: payload?.birth_date?.length
      ? formatDateNumeric(payload?.birth_date[0].toString())
      : null,
  };

  try {
    const response = await fetchApi(`/student/biodata/parents/${parentType}`, {
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
