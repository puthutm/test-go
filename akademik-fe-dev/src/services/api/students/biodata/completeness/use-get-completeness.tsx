"use client";

import { useQuery } from "@tanstack/react-query";

import { getCompletenessStudent } from "./get-completeness";

export const useCompletenessStudent = () => {
  return useQuery({
    queryKey: ["completeness-student"],
    queryFn: async () => {
      const data = await getCompletenessStudent();
      return data;
    },
  });
};
