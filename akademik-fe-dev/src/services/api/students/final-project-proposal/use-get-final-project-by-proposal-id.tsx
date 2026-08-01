"use client";

import { useQuery } from "@tanstack/react-query";
import { getFinalProjectProposalByProposalId } from "./get-final-project-proposal-by-proposal-id";

export const useGetFinalProjectByProposalId = (proposalId: string) => {
  const query = useQuery({
    queryKey: ["final-project-proposal-by-id", proposalId],
    queryFn: async () => {
      const response = await getFinalProjectProposalByProposalId(proposalId);

      return response;
    },
    enabled: false,
  });

  return query;
};
