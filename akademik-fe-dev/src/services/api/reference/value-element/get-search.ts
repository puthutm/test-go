"use client";

import { useAxiosDataReferensi } from "@/lib/hooks/use-axios";
import { useQuery } from "@tanstack/react-query";

export const useGetSearchValueElement = (queryParam: QueryParamDataRefensi) => {
  const axios = useAxiosDataReferensi();

  let params: QueryParamDataRefensi = {
    page: queryParam.page,
    page_size: queryParam.page_size,
    sort_by: queryParam.sort_by,
    sort_direction: queryParam.sort_direction,
  };

  if (queryParam.filter) {
    params = {
      ...params,
      filter: queryParam.filter,
    };
  }

  const fetchSearchValueElement = async (): Promise<
    ApiResponse<ValueElement[]> | undefined
  > => {
    try {
      const { data } = await axios.get("/academic/value-elements/search", {
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
    queryKey: ["get-search-value-element"],
    queryFn: fetchSearchValueElement,
  });

  return { ...query };
};
