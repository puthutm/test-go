"use server";

import { revalidatePath } from "next/cache";

import { fetchApi } from "@/lib/utils/fetch-server";
import { StudentAccountSchemaType } from "@/lib/validations/academic/portal/student-account-schema";

export const createStudent = async (payload: StudentAccountSchemaType) => {
  try {
    const response = await fetchApi("/academic/portal/students", {
      method: "POST",
      body: JSON.stringify(payload),
      headers: {
        "Content-Type": "application/json",
      },
    });

    revalidatePath("/portal/students", "page");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Internal Server Error"
    );
  }
};
