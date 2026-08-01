"use client";

import { useAxios } from "@/lib/hooks/use-axios";
import { useQuery } from "@tanstack/react-query";

export const useGetAllSubject = (queryParam: QueryParam) => {
  const axios = useAxios();

  let params: QueryParam = {
    page: queryParam.page,
    limit: queryParam.limit,
    sort_by: queryParam.sort_by,
    sort_direction: queryParam.sort_direction,
  };

  if (queryParam.search) {
    params = {
      ...params,
      search: queryParam.search,
    };
  }

  const fetchSubjects = async (): Promise<
    ApiResponse<PaginationData<Subject>> | undefined
  > => {
    try {
      const { data } = await axios.get("/academic/setting/subjects", {
        params,
      });

      return data;
    } catch (error) {
      throw new Error(
        error instanceof Error ? error.message : "Something went wrong"
      );
    }
  };

  const query = useQuery({
    queryKey: [
      "get-all-subject",
      params.page,
      params.limit,
      params.sort_by,
      params.sort_direction,
    ],
    queryFn: fetchSubjects,
  });

  return { ...query };
};
