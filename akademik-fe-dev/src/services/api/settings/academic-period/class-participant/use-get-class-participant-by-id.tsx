"use client";

import { useQuery } from "@tanstack/react-query";

import { getClassParticipantById } from "./get-class-participant-by-id";

export const useGetClassParticipantById = (
  classId: string,
  participantId: string
) => {
  return useQuery({
    queryKey: ["class-participant-by-id", classId, participantId],
    queryFn: async () => await getClassParticipantById(classId, participantId),
    enabled: !!classId && !!participantId,
  });
};
