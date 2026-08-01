import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";
import Link from "next/link";

import { VisibilityIcon } from "@/components/icons/visibility";
import { formatDate } from "@/lib/utils/format-date";

export const useTableAcademicPeriodColumns = () => {
  const columns: ColumnDef<AcademicPeriod>[] = [
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
      header: "Nama Periode",
      cell: ({ row }) => <p className="text-start">{row.original.fullname}</p>,
    },
    {
      accessorKey: "college_start",
      header: "Tgl. Mulai Kuliah",
      cell: ({ row }) => (
        <p className="text-start">
          {formatDate(row.original.start_date_of_college ?? "")}
        </p>
      ),
    },
    {
      accessorKey: "college_end",
      header: "Tgl. Berakhir Kuliah",
      cell: ({ row }) => (
        <p className="text-start">
          {formatDate(row.original.end_date_of_college ?? "")}
        </p>
      ),
    },
    {
      accessorKey: "uts_date",
      header: "Tgl. UTS Kuliah",
      cell: ({ row }) => (
        <p className="text-start">{`${formatDate(
          row.original.start_date_of_uts ?? ""
        )} - ${formatDate(row.original.end_date_of_uts ?? "")}`}</p>
      ),
    },
    {
      accessorKey: "uas_date",
      header: "Tgl. UAS Kuliah",
      cell: ({ row }) => (
        <p className="text-start">{`${formatDate(
          row.original.start_date_of_uas ?? ""
        )} - ${formatDate(row.original.end_date_of_uas ?? "")}`}</p>
      ),
    },
    {
      accessorKey: "is_active",
      header: "Sedang Berjalan",
      cell: ({ row }) => (
        <p className="d-flex justify-content-center text-white">
          {row.original.is_active ? (
            <span className="bg-success px-3 py-1 rounded-pill">Aktif</span>
          ) : (
            <span className="bg-danger px-3 py-1 rounded-pill">Nonaktif</span>
          )}
        </p>
      ),
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
            <Link
              href={`/curriculum/academic-period/${row.original.id}/classes`}
            >
              <VisibilityIcon />
            </Link>
          </Button>
        </div>
      ),
    },
  ];

  return { columns };
};
