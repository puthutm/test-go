"use client";
import { useQuery } from "@tanstack/react-query";

import { getWeeklySchedule } from "./get-weekly-schedule";

export const useGetWeeklySchedule = (idClass: string) => {
  return useQuery({
    queryKey: ["get-detail-weekly-schedule", idClass],
    queryFn: async () => {
      const data = await getWeeklySchedule(idClass);
      return data;
    },
    retry: 0,
  });
};
