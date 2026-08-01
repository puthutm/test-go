"use client";

import { useQuery } from "@tanstack/react-query";

import { getVillagesByDistrictId } from "./get-village-by-district-id";

export const useVillagesByDistrictId = (districtId: string) => {
  return useQuery({
    queryKey: ["village-by-district-id", districtId],
    queryFn: async () => {
      const data = await getVillagesByDistrictId(districtId);
      return data;
    },
    enabled: !!districtId,
  });
};
