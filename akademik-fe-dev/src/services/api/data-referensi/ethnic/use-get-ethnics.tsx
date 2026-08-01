"use client";

import { useQuery } from "@tanstack/react-query";

import { getEthnics } from "./get-ethnics";

export const useEthnics = () => {
  return useQuery({
    queryKey: ["ethnics"],
    queryFn: async () => {
      const data = await getEthnics();
      return data;
    },
  });
};
