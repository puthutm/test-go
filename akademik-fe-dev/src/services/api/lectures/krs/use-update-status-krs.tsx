"use client";

import { useMutation } from "@tanstack/react-query";

import { updateKrsStatus } from "./update-status-krs";
import { FormKrsStatusSchemaType } from "@/lib/validations/lecturer/update-status-reject-approve-krs";

export const useUpdateStatusKrs = () => {
  return useMutation({
    mutationFn: async ({
      krsItemId,
      payload,
    }: {
      krsItemId: string;
      payload: FormKrsStatusSchemaType;
    }) => await updateKrsStatus(krsItemId, payload),
  });
};
