"use client";

import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";
import { Button } from "reactstrap";

import { EditIcon } from "@/components/icons/edit";
import { VisibilityIcon } from "@/components/icons/visibility";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { DeleteIcon } from "@/components/icons/delete";

export const useColumnDefinitionSubject = (curriculumYearId: string) => {
  const { setModalConfirmationState } = useModalConfirmationContext();

  const handleDelete = (id: string) => {
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "confirm",
      message: "Hapus Mata Kuliah",
      id: id,
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
          <Link
            href={`/curriculum/curriculum-year/${curriculumYearId}/subjects/${row.original.id}/edit`}
          >
            <Button
              className="bg-transparent border-0 text-black p-0 fs-4"
              title="Ubah"
            >
              <EditIcon height="20" width="20" />
            </Button>
          </Link>
          <Link
            href={`/curriculum/curriculum-year/${curriculumYearId}/subjects/${row.original.id}/detail`}
          >
            <Button
              className="bg-transparent border-0 text-black p-0 fs-4"
              title="Detail"
            >
              <VisibilityIcon />
            </Button>
          </Link>
          <Button
            title="Hapus"
            className="bg-transparent border-0 text-black p-0"
            onClick={() => handleDelete(row.original.id)}
          >
            <DeleteIcon width="20" height="20" />
          </Button>
        </div>
      ),
    },
  ];

  return { columns };
};
