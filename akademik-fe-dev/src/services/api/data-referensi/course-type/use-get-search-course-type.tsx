"use client";

import { useQuery } from "@tanstack/react-query";
import { getSearchCourseTypes } from "./get-search-course-type";

export const useSearchCourseType = () => {
  const query = useQuery({
    queryKey: ["search-course-type"],
    queryFn: async () => {
      const data = await getSearchCourseTypes({
        page: 1,
        filter: "",
        page_size: 100,
      });

      return data;
    },
  });

  return query;
};
