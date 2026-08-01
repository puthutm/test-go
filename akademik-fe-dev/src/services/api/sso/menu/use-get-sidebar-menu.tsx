"use client";

import { useQuery } from "@tanstack/react-query";

import { getSidebarMenu } from "./get-sidebar.-menu";

export const useGetSidebarMenu = () => {
  return useQuery({
    queryKey: ["sidebar-menu"],
    queryFn: async () => {
      const data = await getSidebarMenu();
      return data;
    },
  });
};
