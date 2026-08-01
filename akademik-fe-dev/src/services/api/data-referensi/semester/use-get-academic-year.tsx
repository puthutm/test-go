"use client";

import { useQuery } from "@tanstack/react-query";

import { getSemesters } from "./get-semester";

export const useSemesters = () => {
  return useQuery({
    queryKey: ["semesters"],
    queryFn: async () => {
      const data = await getSemesters();
      return data;
    },
  });
};
