"use client";

import { useQuery } from "@tanstack/react-query";

import { getGrade } from "./get-grade";

export const useGrade = (status:boolean = true) => {
  return useQuery({
    queryKey: ["option-grade"],
    queryFn: async () => {
      const data = await getGrade();
      return data;
    },
    enabled:status
  });
};
