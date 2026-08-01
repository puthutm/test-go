"use client";

import { useQuery } from "@tanstack/react-query";
import { getCurriculumYearById } from "./get-curriculum-year-by-id";

export const useGetCurriculumYearById = (id: string) => {
  return useQuery({
    queryKey: ["curriculum-year", id],
    queryFn: async () => getCurriculumYearById(id),
    enabled: !!id,
  });
};
