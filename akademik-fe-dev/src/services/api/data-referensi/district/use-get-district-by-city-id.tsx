"use client";

import { useQuery } from "@tanstack/react-query";

import { getDistrictsByCityId } from "./get-districts-by-city-id";

export const useDistrictsByCityId = (cityId: string) => {
  return useQuery({
    queryKey: ["district-by-city-id", cityId],
    queryFn: async () => {
      const data = await getDistrictsByCityId(cityId);
      return data;
    },
    enabled: !!cityId,
  });
};
