"use client";
import { useQuery } from "@tanstack/react-query";

import { getCourseAssignment } from "./get-course-assignment";

export const useGetCourseAssignment = (idClass: string) => {
  return useQuery({
    queryKey: ["get-course-assignment", idClass],
    queryFn: async () => {
      const data = await getCourseAssignment(idClass);
      return data;
    },
    retry: 0,
  });
};
