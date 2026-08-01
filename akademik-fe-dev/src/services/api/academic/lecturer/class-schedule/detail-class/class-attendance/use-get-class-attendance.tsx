"use client";
import { useQuery } from "@tanstack/react-query";

import { getClassAttendance } from "./get-class-attendance";

export const useGetClassAttendance = (idClass: string) => {
  return useQuery({
    queryKey: ["get-class-attendance", idClass],
    queryFn: async () => {
      const data = await getClassAttendance(idClass);
      return data;
    },
    retry: 0,
  });
};
