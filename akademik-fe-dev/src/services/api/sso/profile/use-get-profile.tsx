"use client";

import { useQuery } from "@tanstack/react-query";

import { getProfile } from "./get-profle";

export const useGetProfile = () => {
  return useQuery({
    queryKey: ["profile"],
    queryFn: async () => {
      const data = await getProfile();
      return data;
    },
  });
};
