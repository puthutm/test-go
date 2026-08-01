"use client";

import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";

export const useColumnPresenceSubject = () => {
  const columns: ColumnDef<SubjectsPresence>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return row.index + 1;
      },
    },
    {
      accessorKey: "academic_periode_name",
      header: "Periode",
      cell: ({ row }) => (
        <p className="text-start">{row.original.academic_periode_name}</p>
      ),
    },
    {
      accessorKey: "subject_name_id",
      header: "Nama Mata Kuliah",
    },
    {
      accessorKey: "class_count",
      header: "Jumlah Kelas",
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-2 items-content-center justify-content-center">
          <Link
            href={`/settings/presence/${row.original.subject_id}/classes?period=${row.original.academic_periode_id}&studyProgram=${row.original.study_program_id}`}
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Detail"
          >
            <i className="ri-eye-line text-primary text-lg"></i>
          </Link>
        </div>
      ),
    },
  ];

  return { columns };
};
