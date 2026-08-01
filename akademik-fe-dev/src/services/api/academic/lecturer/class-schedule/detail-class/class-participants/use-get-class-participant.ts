"use client";
import { useQuery } from "@tanstack/react-query";

import { getClassParticipantClassSchedule } from "./get-class-participants";

export const useGetClassParticipantDetailClass = (idClass: string) => {
  return useQuery({
    queryKey: ["get-student-class-participant-detail-class", idClass],
    queryFn: async () => {
      const data = await getClassParticipantClassSchedule(idClass);
      return data;
    },
    retry: 0,
  });
};
