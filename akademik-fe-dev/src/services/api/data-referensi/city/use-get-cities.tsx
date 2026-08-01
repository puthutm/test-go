"use client";

import { useQuery } from "@tanstack/react-query";

import { getCities } from "./get-cities";

export const useCities = (search?: string) => {
  return useQuery({
    queryKey: ["cities", search],
    queryFn: async () => {
      const data = await getCities(search);
      return data;
    },
  });
};
