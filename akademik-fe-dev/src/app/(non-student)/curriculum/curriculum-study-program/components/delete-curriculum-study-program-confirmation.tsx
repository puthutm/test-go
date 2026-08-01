"use client";

import { useQueryClient } from "@tanstack/react-query";
import { useState } from "react";

import { ModalDeleteConfirmation } from "@/components/ui/modal-delete-confirmation";
import { AKADEMIK } from "@/lib/constants/role";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { deleteCurriculumStudyProgramForAcademic } from "@/services/api/curriculum/curriculum-study-program/academic/delete-curriculum-study-program-for-academic";
import { deleteCurriculumStudyProgramForProgramHead } from "@/services/api/curriculum/curriculum-study-program/program-head/delete-curriculum-study-program-for-program-head";

export const DeleteCurriculumStudyProgram = ({ role }: { role: string }) => {
  const [loading, setLoading] = useState<boolean>(false);

  const { modalConfirmationState, setModalConfirmationState } =
    useModalConfirmationContext();

  const queryClient = useQueryClient();
  const onDelete = async () => {
    try {
      setLoading(true);
      const response =
        role === AKADEMIK
          ? await deleteCurriculumStudyProgramForAcademic(
              modalConfirmationState?.id as string
            )
          : await deleteCurriculumStudyProgramForProgramHead(
              modalConfirmationState?.id as string
            );

      if (!response.error) {
        setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "success",
          message: "Data berhasil dihapus",
          id: undefined,
        }));

        queryClient.refetchQueries({
          queryKey: ["curriculum-study-program-for-program-head"],
        });

        queryClient.refetchQueries({
          queryKey: ["curriculum-study-program-for-academic"],
        });

        return;
      }

      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "failed",
        message: "Data gagal dihapus",
        id: undefined,
      }));
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: !prev.open,
        state: "failed",
        message: error,
        id: undefined,
      }));
    } finally {
      setLoading(false);
    }
  };

  return <ModalDeleteConfirmation isLoading={loading} onDelete={onDelete} />;
};
