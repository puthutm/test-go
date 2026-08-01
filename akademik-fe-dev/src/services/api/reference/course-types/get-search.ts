"use client";

import { useAxiosDataReferensi } from "@/lib/hooks/use-axios";
import { useQuery } from "@tanstack/react-query";

export const useGetSearchCourseTypes = (queryParam: QueryParamDataRefensi) => {
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

  const fetchCourseTypes = async (): Promise<
    ApiResponse<CourseType[]> | undefined
  > => {
    try {
      const { data } = await axios.get("/academic/course-types/search", {
        params,
      });

      return data;
    } catch (error) {
      throw new Error(
        error instanceof Error
          ? error.message
          : "Something went wrong while fetching course types"
      );
    }
  };

  const query = useQuery({
    queryKey: ["get-search-course-types"],
    queryFn: fetchCourseTypes,
    enabled: false,
  });

  return { ...query };
};
