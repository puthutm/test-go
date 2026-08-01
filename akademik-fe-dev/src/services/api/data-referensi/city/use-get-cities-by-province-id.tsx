"use client";

import { useQuery } from "@tanstack/react-query";

import { getCitiesByProvinceId } from "./get-cities-by-province-id";

export const useCitiesByProvinceId = (provinceId: string) => {
  return useQuery({
    queryKey: ["cities-by-province-id", provinceId],
    queryFn: async () => {
      const data = await getCitiesByProvinceId(provinceId);
      return data;
    },
    enabled: !!provinceId,
  });
};
