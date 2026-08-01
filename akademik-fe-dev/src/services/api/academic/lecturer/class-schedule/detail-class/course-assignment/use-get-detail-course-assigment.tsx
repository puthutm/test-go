"use client";
import { useQuery } from "@tanstack/react-query";

import { getDetailCourseAssignment } from "./get-detail-course-assignment";

export const useGetDetailCourseAssignment = (idClass: string,idCourseAssigment:string | undefined | null) => {
  return useQuery({
    queryKey: ["get-detail-course-assignment", idClass,idCourseAssigment],
    queryFn: async () => {
      const data = await getDetailCourseAssignment(idClass,idCourseAssigment);
      return data;
    },
    retry: 0,
    enabled:idCourseAssigment != null ? true :false
  });
};
