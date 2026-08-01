"use client";

import { useQuery } from "@tanstack/react-query";
import { getCurriculumYear } from "./get-curriculum-year";

export const useGetOptionCurriculumYear = (params:QueryParamDataRefensi,statusEnabled:boolean=false) => {
  return useQuery({
    queryKey: ["get-option-curriculum-year",params.page,params.page_size,params.filter],
    queryFn: async () => {
      const data = await getCurriculumYear(params);
      return data;
    },
    retry:0,
    enabled:statusEnabled
  });
};
