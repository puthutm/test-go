"use client";

import { useQuery } from "@tanstack/react-query";

import { getAlmamaterSizes } from "./get-almamater-sizes";

export const useAlmamaterSizes = () => {
  return useQuery({
    queryKey: ["almamater-sizes"],
    queryFn: async () => {
      const data = await getAlmamaterSizes();
      return data;
    },
  });
};
