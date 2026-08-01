"use client";

import { useQuery } from "@tanstack/react-query";
import { getSubjects } from "./get-all-subject";

export const useGetAllSubjects = (queryParam: QueryParam) => {
  const query = useQuery({
    queryKey: [
      "subject",
      queryParam.search,
      queryParam.page,
      queryParam.limit,
      queryParam.sort_by,
      queryParam.sort_direction,
    ],
    queryFn: async () => {
      const data = await getSubjects(queryParam);
      return data;
    },
  });

  return query;
};
