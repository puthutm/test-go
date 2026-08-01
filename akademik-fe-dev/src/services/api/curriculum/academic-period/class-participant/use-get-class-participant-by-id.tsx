"use client";

import { useQuery } from "@tanstack/react-query";

import { getClassParticipantByIdForProgramHead } from "./get-class-participant-by-id";

export const useGetClassParticipantByIdForProgramHead = (
  classId: string,
  participantId: string
) => {
  return useQuery({
    queryKey: ["class-participant-by-id", classId, participantId],
    queryFn: async () =>
      await getClassParticipantByIdForProgramHead(classId, participantId),
    enabled: !!classId && !!participantId,
  });
};
