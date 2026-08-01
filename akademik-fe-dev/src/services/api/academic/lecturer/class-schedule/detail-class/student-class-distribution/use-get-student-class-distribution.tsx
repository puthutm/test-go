"use client";
import { useQuery } from "@tanstack/react-query";

import { getStudentClassDistribution } from "./get-student-class-distribution";

export const useGetStudentClassDistribution = (idClass: string) => {
  return useQuery({
    queryKey: ["get-student-class-distribution", idClass],
    queryFn: async () => {
      const data = await getStudentClassDistribution(idClass);
      return data;
    },
    retry: 0,
  });
};
