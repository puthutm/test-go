import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";
import Link from "next/link";

import { VisibilityIcon } from "@/components/icons/visibility";
import { ModeEditIcon } from "@/components/icons/mode-edit";

export const useTableSubjectAcademicPeriodColumns = (params: any) => {
  const columns: ColumnDef<Class>[] = [
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
      accessorKey: "name",
      header: "Nama Mata Kuliah",
      cell: ({ row }) => <p className="text-start">{row.original.name}</p>,
    },
    {
      accessorKey: "capacity",
      header: "Kapasitas",
      cell: ({ row }) => <p className="text-start">{row.original.capacity}</p>,
    },
    {
      accessorKey: "total_participant",
      header: "Peserta",
      cell: ({ row }) => (
        <p className="text-start">{row.original.total_participant}</p>
      ),
    },
    {
      accessorKey: "lecturer_name",
      header: "Dosen Pengajar",
      cell: ({ row }) => (
        <p className="text-start">{row.original.lecturer_name}</p>
      ),
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-1 justify-content-center">
          <Link
            href={`/curriculum/academic-period/${params.academicPeriodId}/classes/${row.original.id}/edit`}
            title="Ubah"
            className="bg-transparent border-0 text-black p-0 fs-4"
          >
            <ModeEditIcon />
          </Link>
          <Link
            href={`/curriculum/academic-period/${params.academicPeriodId}/classes/${row.original.id}/detail`}
            title="Detail"
            className="bg-transparent border-0 text-black p-0 fs-4"
          >
            <VisibilityIcon />
          </Link>
          <Button
            title="Hapus"
            className="bg-transparent border-0 text-black p-0"
          >
            <i className="mdi mdi mdi-trash-can-outline fs-4 text-danger" />
          </Button>
        </div>
      ),
    },
  ];

  return { columns };
};
