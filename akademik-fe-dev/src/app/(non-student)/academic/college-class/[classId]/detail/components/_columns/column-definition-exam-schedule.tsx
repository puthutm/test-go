"use client";
import { ColumnDef } from "@tanstack/react-table";

// import hook
import { IDummyValueTable } from "../_sections/section-exam-schedule";

import { Button } from "reactstrap";

import { DeleteIcon } from "@/components/icons/delete";
import { EditIcon } from "@/components/icons/edit";
import { VisibilityIcon } from "@/components/icons/visibility";

interface iColumnsParams {
  (): { columns: ColumnDef<IDummyValueTable>[] };
}

const useColumnExamSchedule: iColumnsParams = () => {
  const columns: ColumnDef<IDummyValueTable>[] = [
    {
      header: "No",
      accessorKey: "no",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return <p className="m-0 p-0 text-start">{row.index + 1}</p>;
      },
    },
    {
      header: "Jenis Ujian",
      accessorKey: "jenis_ujian",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.jenis_ujian ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Kelompok",
      accessorKey: "kelompok",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.kelompok ?? "-"}</p>
        );
      },
    },
    {
      header: "Tanggal",
      accessorKey: "tanggal",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.tanggal ?? "-"}</p>
        );
      },
    },
    {
      header: "Ruang",
      accessorKey: "ruang",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.ruang ?? "-"}</p>
        );
      },
    },
    {
      header: "Peserta",
      accessorKey: "peserta",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.peserta ?? "-"}</p>
        );
      },
    },
    {
      header: "Aksi",
      enableSorting: false,
      cell: () => {
        return (
          <div className="d-flex gap-2 justify-content-center align-items-center">
            {/*//! action edit */}
            <Button
              href="subject/edit/123"
              className="bg-transparent border-0 text-black p-0"
            >
              <EditIcon color="#0AB39C" width="20" height="20" />
            </Button>
            {/*//! action view ap */}
            <Button
              href="subject/detail/123?tab=student"
              className="bg-transparent border-0 text-black p-0"
            >
              <VisibilityIcon color="#2E3192" width="20" height="20" />
            </Button>
            {/*//! action delete */}
            <Button className="bg-transparent border-0 text-black p-0">
              <DeleteIcon color="#F06548" width="20" height="20" />
            </Button>
          </div>
        );
      },
    },
  ];

  return { columns };
};

export default useColumnExamSchedule;
