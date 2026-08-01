"use client";

import { useQuery } from "@tanstack/react-query";

import { getEducationalLevels } from "./get-educational-level";

export const useEducationalLevels = () => {
  return useQuery({
    queryKey: ["educational-level"],
    queryFn: async () => {
      const data = await getEducationalLevels();
      return data;
    },
  });
};
