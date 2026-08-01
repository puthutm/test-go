"use server";

import { revalidatePath } from "next/cache";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";
import { formatDateNumeric } from "@/lib/utils/format-date";
import { CurriculumYearFormType } from "@/lib/validations/academic/settings/curriculum-year";

export const updateCurriculumYear = async (
  curriculumYearId: string,
  payload: CurriculumYearFormType
) => {
  const reqBody: CurriculumYearForm = {
    years: payload.years,
    starts: payload.starts.value,
    start_date: formatDateNumeric(payload.start_date[0].toString()),
    end_date: formatDateNumeric(payload.end_date[0].toString()),
    description: payload.description ?? null,
  };

  try {
    const response = await fetchApiDatareferensi(
      `/academic/curriculum-years/${curriculumYearId}`,
      {
        method: "PUT",
        body: JSON.stringify(reqBody),
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    revalidatePath("/settings/curriculum-year", "page");

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Internal Server Error"
    );
  }
};
