"use client";

import { useQuery } from "@tanstack/react-query";
import { getSearchCourseGroups } from "./get-search-course-group";

export const useSearchCourseGroup = () => {
  const query = useQuery({
    queryKey: ["search-course-group"],
    queryFn: async () => {
      const data = await getSearchCourseGroups({
        page: 1,
        filter: "",
        page_size: 100,
      });

      return data;
    },
  });

  return query;
};
