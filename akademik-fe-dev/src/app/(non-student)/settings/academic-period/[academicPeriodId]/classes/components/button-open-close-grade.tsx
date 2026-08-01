"use client";

import { Button } from "reactstrap";

import { LockIcon } from "@/components/icons/lock";
import { LockOpenIcon } from "@/components/icons/lock-open";
import { useMutation } from "@tanstack/react-query";
import { updateOpenCloseGradeByAcademicPeriod } from "@/services/api/settings/academic-period/open-close-grade/update-open-close-grade-by-academic-period";
import { FormOpenCloseGradeSchemaType } from "@/lib/validations/settings/academic-period/form-open-close-grade";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

interface Props {
  academicPeriodId: string;
  statusLock: boolean;
}

export const ButtonOpenCloseGrade = ({
  statusLock,
  academicPeriodId,
}: Props) => {
  const { setModalConfirmationState } = useModalConfirmationContext();

  const { mutateAsync: toggleOpenCloseGrade, isPending } = useMutation({
    mutationFn: async ({
      academicPeriodId,
      payload,
    }: {
      academicPeriodId: string;
      payload: FormOpenCloseGradeSchemaType;
    }) =>
      await updateOpenCloseGradeByAcademicPeriod({ academicPeriodId, payload }),
    onSuccess(data) {
      if (data?.status === 200) {
        setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "success",
          message: statusLock
            ? "Nilai berhasil dibuka"
            : "Nilai berhasil ditutup",
        }));
      }
    },
    onError(error) {
      if (error) {
        setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "success",
          message: error.message,
        }));
      }
    },
  });
  return (
    <Button
      color={statusLock ? "danger" : "success"}
      onClick={async () =>
        await toggleOpenCloseGrade({
          academicPeriodId,
          payload: { status_locked: statusLock ? false : true },
        })
      }
      disabled={isPending}
    >
      {statusLock ? (
        <>
          <LockOpenIcon /> Open Nilai
        </>
      ) : (
        <>
          <LockIcon /> Lock Nilai
        </>
      )}
    </Button>
  );
};
