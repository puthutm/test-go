"use client";

import { useQuery } from "@tanstack/react-query";

import { getStudentStatuses } from "./get-student-statuses";

export const useStudentStatuses = () => {
  return useQuery({
    queryKey: ["student-statuses"],
    queryFn: async () => {
      const data = await getStudentStatuses();
      return data;
    },
  });
};
