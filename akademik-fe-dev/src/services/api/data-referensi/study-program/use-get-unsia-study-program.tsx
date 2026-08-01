"use client";

import { useQuery } from "@tanstack/react-query";

import { getUnsiaStudyProgram } from "./get-unsia-study-program";

export const useGetUnsiaStudyProgram = (statusActive: boolean = true) => {
  const query = useQuery({
    queryKey: ["unsia-study-program"],
    queryFn: async () => {
      const response = await getUnsiaStudyProgram();
      return response;
    },
    enabled: statusActive,
  });

  return query;
};
