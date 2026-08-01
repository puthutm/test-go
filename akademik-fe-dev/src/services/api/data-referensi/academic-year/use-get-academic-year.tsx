"use client";

import { useQuery } from "@tanstack/react-query";

import { getSearchAcademicYears } from "./get-search-academic-year";

export const useAcademicYears = () => {
  return useQuery({
    queryKey: ["academic-years"],
    queryFn: async () => {
      const data = await getSearchAcademicYears();
      return data;
    },
  });
};
