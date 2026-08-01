"use client";

import { useQuery } from "@tanstack/react-query";
import { getAcademicYearById } from "./get-academic-year-by-id";

export const useGetAcademicYearById = (academicYearId: string) => {
  return useQuery({
    queryKey: ["academic-year-by-id", academicYearId],
    queryFn: async () => await getAcademicYearById(academicYearId),
    enabled: !!academicYearId,
  });
};
