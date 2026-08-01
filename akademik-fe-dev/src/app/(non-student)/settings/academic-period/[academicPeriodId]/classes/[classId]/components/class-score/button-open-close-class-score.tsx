"use client";

import { Button } from "reactstrap";
import { useMutation } from "@tanstack/react-query";

import { LockIcon } from "@/components/icons/lock";
import { LockOpenIcon } from "@/components/icons/lock-open";
import { FormOpenCloseGradeSchemaType } from "@/lib/validations/settings/academic-period/form-open-close-grade";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { updateOpenCloseGradeByClassId } from "@/services/api/settings/academic-period/open-close-grade/update-open-close-grade-by-class";

interface Props {
  academicPeriodId: string;
  classId: string;
  statusLock: boolean;
}

export const ButtonOpenCloseClassScore = ({
  statusLock,
  academicPeriodId,
  classId,
}: Props) => {
  const { setModalConfirmationState } = useModalConfirmationContext();

  const { mutateAsync: toggleOpenCloseGrade, isPending } = useMutation({
    mutationFn: async ({
      academicPeriodId,
      classId,
      payload,
    }: {
      academicPeriodId: string;
      classId: string;
      payload: FormOpenCloseGradeSchemaType;
    }) =>
      await updateOpenCloseGradeByClassId({
        academicPeriodId,
        classId,
        payload,
      }),
    onSuccess(data) {
      if (data?.status === 200) {
        setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "success",
          message: statusLock
            ? "Nilai kelas berhasil dibuka"
            : "Nilai kelas berhasil ditutup",
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
          classId,
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
