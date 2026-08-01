"use client";
import { useQuery } from "@tanstack/react-query";
import { getSubjectDetailLecture } from "./get-detail-subject-lecture";


export const useGetDetailSubjectLecturer = (idSubject:string) => {
  return useQuery({
    queryKey: ["get-detail-subject-lecturer",idSubject],
    queryFn: async () => {
      const data = await getSubjectDetailLecture(idSubject);
      return data;
    },
    retry:0,
  });
};
