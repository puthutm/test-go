"use client";

import { useQuery } from "@tanstack/react-query";

import { getBloodTypes } from "./get-blood-types";

export const useBloodTypes = () => {
  return useQuery({
    queryKey: ["blood-types"],
    queryFn: async () => {
      const data = await getBloodTypes();
      return data;
    },
  });
};
