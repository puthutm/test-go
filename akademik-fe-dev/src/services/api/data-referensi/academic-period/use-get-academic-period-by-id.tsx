"use client";

import { useQuery } from "@tanstack/react-query";

import { getAcademicPeriodById } from "./get-academic-period-by-id";

export const useGetAcademicPeriodById = (id: string) => {
  return useQuery({
    queryKey: ["academic-period", id],
    queryFn: () => getAcademicPeriodById(id),
    enabled: !!id,
  });
};
