"use client";

import { useQuery } from "@tanstack/react-query";
import { getDetailSksLimits } from "./get-detail-sks-limits";

export const useGetDetailSksLimit = (
  sksLimitId: string | null 
) => {
  const query = useQuery({
    queryKey: ["detail-sks-limit", sksLimitId],
    queryFn: async () => {
      const response = await getDetailSksLimits(sksLimitId as string);
      return response;
    },
    enabled: sksLimitId != null ? true : false,
    gcTime:0
  });

  return query;
};
