"use server";

import { revalidatePath } from "next/cache";

import { FormInformationSchemaType } from "@/lib/validations/students/biodata/form-information-schema";
import { fetchApi } from "@/lib/utils/fetch-server";

export const updateInformationStudent = async (
  payload: FormInformationSchemaType
): Promise<ApiResponse<InformationStudent>> => {
  const reqBody: FormInformationStudent = {
    almamater_size_id: payload.almamater_size_id.value,
    citizenship_id: payload.citizenship_id.value,
    job_id: payload.job_id.value,
    phone: payload.phone,
    private_email: payload.private_email,
    transportation_id: payload.transportation_id.value,
  };

  try {
    const response = await fetchApi("/student/biodata/informations", {
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
