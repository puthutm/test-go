"use client";

import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";
import { useState } from "react";

import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { RefreshIcon } from "@/components/icons/refresh";
import { restoreSubject } from "@/services/api/settings/subject/trash/restore-trash-subject";

export const useColumnDefinitionSubjectTrash = () => {
  const [loadingRestore, setLoadingRestore] = useState<boolean>(false);

  const { setModalConfirmationState } = useModalConfirmationContext();

  const onRestore = async (id: string) => {
    setLoadingRestore(true);
    const response = await restoreSubject(id as string);

    if (!response.error) {
      setModalConfirmationState((prev) => ({
        ...prev,
        id: undefined,
        open: true,
        message: response?.message.toString(),
      }));
    }
    setModalConfirmationState((prev) => ({
      ...prev,
      open: true,
      state: "success",
      message: "Data berhasil di-restore",
    }));

    setLoadingRestore(false);
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
      cell: ({ row }) => <p className="text-start">{row.original.code}</p>,
    },
    {
      accessorKey: "name_id",
      header: "Nama Mata Kuliah",
      cell: ({ row }) => <p className="text-start">{row.original.name_id}</p>,
    },
    {
      accessorKey: "study_program_name",
      header: "Prodi Pengampu",
      cell: ({ row }) => (
        <p className="text-start">{row.original.study_program_name}</p>
      ),
    },
    {
      accessorKey: "total_sks",
      header: "SKS",
      cell: ({ row }) => <p className="text-start">{row.original.total_sks}</p>,
    },
    {
      accessorKey: "course_type_name",
      header: "Jenis Mata Kuliah",
      cell: ({ row }) => (
        <p className="text-start">{row.original.course_type_name}</p>
      ),
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-1 justify-content-center align-items-center">
          <Button
            title="Restore"
            className="bg-transparent border-0 text-black p-0"
            onClick={() => onRestore(row.original.id)}
            disabled={loadingRestore}
          >
            <RefreshIcon />
          </Button>
        </div>
      ),
    },
  ];

  return { columns };
};
