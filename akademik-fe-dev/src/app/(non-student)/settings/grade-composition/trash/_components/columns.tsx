"use client";

import { useState } from "react";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { restoreGradeComposition } from "@/services/api/settings/grade-composition/trash/restore";
import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";
import { DisabledByDefaultIcon } from "@/components/icons/disabled-by-default";
import { DoneIcon } from "@/components/icons/done";

export const useGradeCompositionTrashColumns = () => {
  const { setModalConfirmationState } = useModalConfirmationContext();

  const [
    isRestoreGradeCompositionLoading,
    setIsRestoreGradeCompositionLoading,
  ] = useState(false);

  const handleRestoreGradeComposition = async (id: string) => {
    setIsRestoreGradeCompositionLoading(true);
    try {
      const response = await restoreGradeComposition(id);
      if (response.error) {
        throw new Error(response.error);
      }
      setModalConfirmationState((prev) => ({
        ...prev,
        open: !prev.open,
        state: "success",
        message: "Data komposisi nilai berhasil dikembalikan.",
      }));
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: !prev.open,
        state: "failed",
        message:
          error.message ||
          "Terjadi kesalahan saat mengembalikan data komposisi nilai.",
      }));
    } finally {
      setIsRestoreGradeCompositionLoading(false);
    }
  };

  const columns: ColumnDef<IGradeComposition>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return <p className="text-start">{row.index + 1}</p>;
      },
    },
    {
      accessorKey: "academic_periode_name",
      header: "Periode Akademik",
    },
    {
      accessorKey: "value_element_name",
      header: "Elemen Nilai",
      cell: ({ row }) => (
        <p className="text-start">{row.original.value_element_name ?? "-"}</p>
      ),
    },
    {
      accessorKey: "percentage",
      header: "Persentase",
      cell: ({ row }) => (
        <p className="text-start">{row.original.percentage ?? "-"}</p>
      ),
    },
    {
      accessorKey: "is_passing_requirement",
      header: "Memenuhi Persyaratan",
      cell: ({ row }) => (
        <p className="text-center">
          {row.original.is_passing_requirement ? (
            <DoneIcon />
          ) : (
            <DisabledByDefaultIcon />
          )}
        </p>
      ),
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex justify-content-center">
          <Button
            className="bg-transparent border-0 text-black p-0 me-1"
            title="Restore"
            onClick={async () =>
              await handleRestoreGradeComposition(row.original.id)
            }
            disabled={isRestoreGradeCompositionLoading}
            key={row.original.id}
          >
            <i className="mdi mdi-backup-restore fs-4 text-secondary" />
          </Button>
        </div>
      ),
    },
  ];

  return { columns };
};
