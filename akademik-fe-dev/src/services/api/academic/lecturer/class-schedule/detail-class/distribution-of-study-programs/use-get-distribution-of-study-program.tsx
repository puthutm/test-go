"use client";
import { useQuery } from "@tanstack/react-query";

import { getDistributionOfStudyProgram } from "./get-distribution-of-study-program";
export const useGetDistributionOfStudyProgram = (idClass: string) => {
  return useQuery({
    queryKey: ["get-detail-distribution-of-programs", idClass],
    queryFn: async () => {
      const data = await getDistributionOfStudyProgram(idClass);
      return data;
    },
    retry: 0,
  });
};
