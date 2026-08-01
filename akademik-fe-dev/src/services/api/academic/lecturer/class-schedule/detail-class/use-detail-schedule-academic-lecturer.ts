"use client";
import { useQuery } from "@tanstack/react-query";
import { getDetailClassScheduleLecturer } from "./get-detail-schedule-academic-lecturer";


export const useGetDetailClassScheduleLecturer = (idClass:string) => {
  return useQuery({
    queryKey: ["get-detail-class-schedule-lecturer",idClass],
    queryFn: async () => {
      const data = await getDetailClassScheduleLecturer(idClass);
      return data;
    },
    retry:0,
  });
};
