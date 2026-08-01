"use client";

import { useQuery } from "@tanstack/react-query";

import { getClassScheduleTemplateById } from "./get-detail-class-schedule-template";

export const useGetClassScheduleTemplateById = ({
  classId,
  classScheduleTemplateId,
}: {
  classId: string;
  classScheduleTemplateId: string;
}) => {
  return useQuery({
    queryKey: [
      "class-schedule-template-by-id",
      classId,
      classScheduleTemplateId,
    ],
    queryFn: async () =>
      await getClassScheduleTemplateById({
        classId,
        classScheduleTemplateId,
      }),
    enabled: false,
  });
};
