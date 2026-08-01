"use client";

import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";
import { usePathname } from "next/navigation";

import { formatDate } from "@/lib/utils/format-date";

export const useColumnPresenceClassSession = () => {
  const pathname = usePathname();

  const columns: ColumnDef<ClassPresenceSession>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return row.index + 1;
      },
    },
    {
      accessorKey: "session_date",
      header: "Tanggal",
      cell: ({ row }) => {
        const date = formatDate(row.original.session_date);

        return date;
      },
    },
    {
      accessorKey: "session",
      header: "Sesi",
    },
    {
      accessorKey: "presence_percentage",
      header: "Kehadiran",
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-2 items-content-center justify-content-center">
          <Link
            href={`${pathname}/${row.original.session_id}/students
            `}
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
