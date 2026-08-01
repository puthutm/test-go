"use client";

import { useQuery } from "@tanstack/react-query";

import { getDetailGradeScale } from "./get-detail-grade-scale";

export const useGetDetailGradeScale = (
  gradeScaleId: string | null 
) => {
  const query = useQuery({
    queryKey: ["detail-grade-scale",gradeScaleId],
    queryFn: async () => {
      const response = await getDetailGradeScale(gradeScaleId as string);
      return response;
    },
    enabled: gradeScaleId != null ? true : false,
    gcTime:0
  });

  return query;
};
