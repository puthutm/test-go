"use server";

import { revalidatePath } from "next/cache";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";
import { FormAcademicYearSchemaType } from "@/lib/validations/settings/academic-year";

export const createAcademicYear = async (
  payload: FormAcademicYearSchemaType
) => {
  const reqBody = {
    name: payload.name,
    years: payload.years,
  };
  try {
    const response = await fetchApiDatareferensi("/pmb/academic-years", {
      method: "POST",
      body: JSON.stringify(reqBody),
      headers: {
        "Content-Type": "application/json",
      },
    });

    revalidatePath("/settings/academic-year", "page");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Internal Server Error"
    );
  }
};
