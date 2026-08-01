"use client";

import { useAxiosDataReferensi } from "@/lib/hooks/use-axios";
import { useQuery } from "@tanstack/react-query";

export const useGetSearchCourseGroups = (queryParam: QueryParamDataRefensi) => {
  const axios = useAxiosDataReferensi();

  const params: QueryParamDataRefensi = {
    page: queryParam.page,
    // page_size: queryParam.page_size,
    sort_by: queryParam.sort_by,
    sort_direction: queryParam.sort_direction,
  };

  // if (queryParam.filter) {
  //   params = {
  //     ...params,
  //     filter: queryParam.filter,
  //   };
  // }

  const fetchCourseGroups = async (): Promise<
    ApiResponse<CourseGroup[]> | undefined
  > => {
    try {
      const { data } = await axios.get("/academic/course-groups/search", {
        params,
      });

      return data;
    } catch (error) {
      throw new Error(
        error instanceof Error
          ? error.message
          : "Something went wrong while fetching course groups"
      );
    }
  };

  const query = useQuery({
    queryKey: ["get-search-course-groups"],
    queryFn: fetchCourseGroups,
    enabled: false,
  });

  return { ...query };
};
