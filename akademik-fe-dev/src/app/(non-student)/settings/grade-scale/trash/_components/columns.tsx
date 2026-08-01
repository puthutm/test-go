"use client";

import { useState } from "react";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { restoreGradeScale } from "@/services/api/settings/grade-scale/trash/restore";
import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";
import { ReplayIcon } from "@/components/icons/replay";

export const useGradeScaleTrashColumns = () => {
  const { setModalConfirmationState } = useModalConfirmationContext();
  const [isRestoreGradeScaleLoading, setIsRestoreGradeScaleLoading] =
    useState<boolean>(false);

  const handleRestoreGradeScale = async (id: string) => {
    setIsRestoreGradeScaleLoading(true);
    try {
      const response = await restoreGradeScale(id);
      if (response.error) {
        throw new Error(response.error);
      }
      setModalConfirmationState((prev) => ({
        ...prev,
        open: !prev.open,
        state: "success",
        message: "Data skala nilai berhasil dikembalikan.",
      }));
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: !prev.open,
        state: "failed",
        message:
          error.message ||
          "Terjadi kesalahan saat mengembalikan data skala nilai.",
      }));
    } finally {
      setIsRestoreGradeScaleLoading(false);
    }
  };

  const columns: ColumnDef<IGradeScale>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return <p className="text-start">{row.index + 1}</p>;
      },
    },
    {
      accessorKey: "study_program_name",
      header: "Program Studi",
      cell: ({ row }) => (
        <p className="text-start">{row.original.study_program_name ?? "-"}</p>
      ),
    },
    {
      accessorKey: "grade_name",
      header: "Nilai",
      cell: ({ row }) => (
        <p className="text-start">{row.original.grade_name ?? "-"}</p>
      ),
    },
    {
      accessorKey: "weight_value",
      header: "Bobot",
      cell: ({ row }) => (
        <p className="text-start">{row.original.weight_value ?? "-"}</p>
      ),
    },
    {
      accessorKey: "lower_value",
      header: "Batas Bawah",
      cell: ({ row }) => (
        <p className="text-start">{row.original.lower_value ?? "-"}</p>
      ),
    },
    {
      accessorKey: "upper_value",
      header: "Batas Atas",
      cell: ({ row }) => (
        <p className="text-start">{row.original.upper_value ?? "-"}</p>
      ),
    },
    {
      accessorKey: "description",
      header: "Deskripsi",
      cell: ({ row }) => (
        <p className="text-start">{row.original.description ?? "-"}</p>
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
            onClick={async () => await handleRestoreGradeScale(row.original.id)}
            disabled={isRestoreGradeScaleLoading}
            key={row.original.id}
          >
            <ReplayIcon />
          </Button>
        </div>
      ),
    },
  ];

  return { columns };
};
