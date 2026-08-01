"use client";

import { useQuery } from "@tanstack/react-query";

import { getDetailGradeComposition } from "./get-detail-grade-composition";

export const useGetDetailGradeComposition = (
  gradeCompositionId: string | null
) => {
  const query = useQuery({
    queryKey: ["detail-grade-composition", gradeCompositionId],
    queryFn: async () => {
      const response = await getDetailGradeComposition(gradeCompositionId as string);
      return response;
    },
    enabled: gradeCompositionId != null ? true : false,
    gcTime: 0
  });

  return query;
};
