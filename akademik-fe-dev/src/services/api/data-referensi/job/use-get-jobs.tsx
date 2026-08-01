"use client";

import { useQuery } from "@tanstack/react-query";

import { getJobs } from "./get-jobs";

export const useJobs = () => {
  return useQuery({
    queryKey: ["jobs"],
    queryFn: async () => {
      const data = await getJobs();
      return data;
    },
  });
};
