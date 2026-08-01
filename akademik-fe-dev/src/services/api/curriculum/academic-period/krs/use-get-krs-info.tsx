"use client";

import { useQuery } from "@tanstack/react-query";

import { getKrsInfo } from "./get-krs-info";

export const useGetKrsInfo = () => {
  return useQuery({
    queryKey: ["krs-info"],
    queryFn: async () => await getKrsInfo(),
  });
};
