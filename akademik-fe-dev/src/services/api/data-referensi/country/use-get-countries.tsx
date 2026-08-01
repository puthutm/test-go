"use client";

import { useQuery } from "@tanstack/react-query";

import { getCountries } from "./get-countries";

export const useCountries = () => {
  return useQuery({
    queryKey: ["countries"],
    queryFn: async () => {
      const data = await getCountries();
      return data;
    },
  });
};
