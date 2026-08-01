"use client";

import { useAxiosDataReferensi } from "@/lib/hooks/use-axios";
import { useQuery } from "@tanstack/react-query";

export const useGetSearchStudyProgram = (queryParam: QueryParamDataRefensi,statusEnabled:boolean = true) => {
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

  const fetchStudyProgram = async (): Promise<
    ApiResponse<StudyProgram[]> | undefined
  > => {
    try {
      const { data } = await axios.get("/education/study-programs/search", {
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
    queryKey: ["get-search-study-program"],
    queryFn: fetchStudyProgram,
    enabled: statusEnabled,
  });

  return { ...query };
};
