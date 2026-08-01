"use client";
import { useQuery } from "@tanstack/react-query";

import { getAcademicSystemDistribution } from "./get-academic-system-distribution";

export const useGetAcademicSystemDistribution = (idClass: string) => {
  return useQuery({
    queryKey: ["get-academic-system-distribution", idClass],
    queryFn: async () => {
      const data = await getAcademicSystemDistribution(idClass);
      return data;
    },
    retry: 0,
  });
};
