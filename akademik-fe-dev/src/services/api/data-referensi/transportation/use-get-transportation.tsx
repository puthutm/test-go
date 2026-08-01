"use client";

import { useQuery } from "@tanstack/react-query";

import { getTransportations } from "./get-transportation";

export const useTransportations = () => {
  return useQuery({
    queryKey: ["transportation"],
    queryFn: async () => {
      const data = await getTransportations();
      return data;
    },
  });
};
