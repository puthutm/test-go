"use client";

import { useQuery } from "@tanstack/react-query";

import { getProvincesByCountryId } from "./get-provinces-by-country-id";

export const useProvincesByCountryId = (countryId: string) => {
  return useQuery({
    queryKey: ["provinces-by-country-id", countryId],
    queryFn: async () => {
      const data = await getProvincesByCountryId(countryId);
      return data;
    },
    enabled: !!countryId,
  });
};
