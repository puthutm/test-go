"use client";
import { useQuery } from "@tanstack/react-query";

import { getDetailClassScheduleSubDetail } from "./get-detail-class-scedule-sub-detail";

export const useGetDetailClassScheduleSubDetail = (
  idClass: string,
  idClassSchedule: string
) => {
  return useQuery({
    queryKey: [
      "get-detail-class-schedule-sub-detail",
      idClass,
      idClassSchedule,
    ],
    queryFn: async () => {
      const data = await getDetailClassScheduleSubDetail(
        idClass,
        idClassSchedule
      );
      return data;
    },
    retry: 0,
  });
};
