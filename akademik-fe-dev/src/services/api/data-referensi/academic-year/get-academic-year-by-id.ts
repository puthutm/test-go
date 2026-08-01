"use server";

import { fetchApiDatareferensi } from "@/lib/utils/fetch-server";

export const getAcademicYearById = async (
  academicYearId: string
): Promise<ApiResponse<AcademicYear>> => {
  try {
    const response = await fetchApiDatareferensi(
      `/pmb/academic-years/${academicYearId}`
    );

    return response;
  } catch (error) {
    throw new Error(
      error instanceof Error ? error.message : "Something went wrong"
    );
  }
};
