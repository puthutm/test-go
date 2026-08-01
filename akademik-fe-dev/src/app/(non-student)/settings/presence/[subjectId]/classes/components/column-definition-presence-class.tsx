"use client";

import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";
import { usePathname, useSearchParams } from "next/navigation";

export const useColumnPresenceClass = () => {
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const academicPeriod = searchParams.get("period");
  const columns: ColumnDef<ClassPresence>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return row.index + 1;
      },
    },
    {
      accessorKey: "class_name",
      header: "Kelas",
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-2 items-content-center justify-content-center">
          <Link
            href={`${pathname}/${row.original.class_id}/sessions?period=${academicPeriod}&studyProgram=${row.original.study_program_id}`}
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
