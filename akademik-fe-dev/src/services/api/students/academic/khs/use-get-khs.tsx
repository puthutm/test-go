"use client";

import { useQuery } from "@tanstack/react-query";
import { getKhsStudent } from "./get-all-khs";

export const useGetKhs = () => {
  return useQuery({
    queryKey: ["khs"],
    queryFn: async () => await getKhsStudent(),
  });
};
