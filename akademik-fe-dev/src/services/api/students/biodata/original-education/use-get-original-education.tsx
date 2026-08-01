"use client";

import { useQuery } from "@tanstack/react-query";

import { getOriginalEducationStudent } from "./get-original-education-student";

export const useOriginalEducationStudent = () => {
  return useQuery({
    queryKey: ["original-education-student"],
    queryFn: async () => {
      const data = await getOriginalEducationStudent();
      return data;
    },
  });
};
