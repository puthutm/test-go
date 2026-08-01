"use client";

import { useQuery } from "@tanstack/react-query";
import { getFinalProjectProposals } from "./get-final-project-proposals";

export const useGetFinalProjectProposal = () => {
  const query = useQuery({
    queryKey: ["final-project-proposal"],
    queryFn: async () => {
      const response = await getFinalProjectProposals();
      return response;
    },
  });

  return query;
};
