"use client";
import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";

import { Button } from "reactstrap";
import { VisibilityIcon } from "@/components/icons/visibility";

export const useTableDefinitionKrsColumn = () => {
  const columns: ColumnDef<KRS>[] = [
    {
      header: "No",
      accessorKey: "no",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return <p className="m-0 p-0 text-center">{row.index + 1}</p>;
      },
    },
    {
      header: "Periode Akademik",
      accessorKey: "academic_periode_name",
    },
    {
      header: "Nama Mahasiswa",
      accessorKey: "student_name",
    },
    {
      header: "NIM",
      accessorKey: "student_nim",
    },
    {
      header: "Jumlah SKS",
      accessorKey: "total_sks",
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-1 justify-content-center">
          <Button
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Detail"
          >
            <Link href={`/lectures/krs/${row.original.krs_header_id}/requests`}>
              <VisibilityIcon />
            </Link>
          </Button>
        </div>
      ),
    },
  ];

  return { columns };
};
