"use client";

import { useQuery } from "@tanstack/react-query";

import { getReligions } from "./get-religions";

export const useReligions = () => {
  return useQuery({
    queryKey: ["religions"],
    queryFn: async () => {
      const data = await getReligions();
      return data;
    },
  });
};
