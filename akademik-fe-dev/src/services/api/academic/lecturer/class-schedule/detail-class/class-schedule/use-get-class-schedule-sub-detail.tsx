"use client";
import { useQuery } from "@tanstack/react-query";

import { getClassScheduleSubDetail } from "./get-class-schedule-sub-detail";

export const useGetClassScheduleSubDetail = (idClass: string,statusEnable:boolean = true) => {
  return useQuery({
    queryKey: ["get-class-schedule-sub-detail", idClass],
    queryFn: async () => {
      const data = await getClassScheduleSubDetail(idClass);
      return data;
    },
    retry: 0,
    enabled:statusEnable
  });
};
