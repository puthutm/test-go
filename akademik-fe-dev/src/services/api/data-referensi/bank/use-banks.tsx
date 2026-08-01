"use client";

import { useQuery } from "@tanstack/react-query";

import { getBanks } from "./get-banks";

export const useBanks = () => {
  return useQuery({
    queryKey: ["banks"],
    queryFn: async () => {
      const data = await getBanks();
      return data;
    },
  });
};
