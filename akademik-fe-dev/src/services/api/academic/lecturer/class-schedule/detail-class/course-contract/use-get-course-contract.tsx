"use client";
import { useQuery } from "@tanstack/react-query";
import { getCourseContract } from "./get-course-contract";

export const useGetCourseContract = (idClass: string) => {
  return useQuery({
    queryKey: ["get-course-contract", idClass],
    queryFn: async () => {
      const data = await getCourseContract(idClass);
      return data;
    },
    retry: 0,
  });
};
