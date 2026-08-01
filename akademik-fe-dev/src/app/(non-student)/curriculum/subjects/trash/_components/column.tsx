"use client";

import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { useRestoreSubject } from "@/services/api/curriculum/subjects/trash/restore";
import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";

export const useSubjectTrashColumn = () => {
  const { setModalConfirmationState } = useModalConfirmationContext();

  const { mutateAsync: onRestore, isPending: isRestoreSubjectLoading } =
    useRestoreSubject();

  const handleRestoreSubject = async (id: string) => {
    await onRestore(id);
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      message: "Berhasil mengembalikan data mata kuliah",
      state: "success",
    }));
  };

  const columns: ColumnDef<Subject>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return row.index + 1;
      },
    },
    {
      accessorKey: "code",
      header: "Kode",
    },
    {
      accessorKey: "curriculum_year_name",
      header: "Tahun Ajaran",
    },
    {
      accessorKey: "name_id",
      header: "Mata Kuliah",
    },
    {
      accessorKey: "study_program_name",
      header: "Program Studi",
    },
    {
      accessorKey: "total_sks",
      header: "SKS",
    },
    {
      accessorKey: "course_type_name",
      header: "Jenis Mata Kuliah",
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex justify-content-center">
          <Button
            className="bg-transparent border-0 text-black p-0 me-1"
            title="Restore"
            onClick={async () => await handleRestoreSubject(row.original.id)}
            disabled={isRestoreSubjectLoading}
            key={row.original.id}
          >
            <i className="mdi mdi-backup-restore fs-4 text-secondary" />
          </Button>
        </div>
      ),
    },
  ];

  return {
    columns,
  };
};
