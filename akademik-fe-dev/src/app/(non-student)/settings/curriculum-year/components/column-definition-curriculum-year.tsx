import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";
import { Button } from "reactstrap";

import { useModalContext } from "@/lib/hooks/use-modal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { formatDate } from "@/lib/utils/format-date";

export const useColumnDefinitionCurriculumYears = () => {
  const { setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const handleDelete = (id: string) => {
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "confirm",
      message: "Hapus Tahun Kurikulum",
      id: id,
    }));
  };
  const columns: ColumnDef<CurriculumYear>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return row.index + 1;
      },
    },
    {
      accessorKey: "year",
      header: "Tahun",
      cell: ({ row }) => <p className="text-start">{row.original.years}</p>,
    },
    {
      accessorKey: "academic_periode_name",
      header: "Mulai Berlaku",
      cell: ({ row }) => (
        <p className="text-start">{row.original.academic_periode_name}</p>
      ),
    },
    {
      accessorKey: "start_date",
      header: "Tanggal Mulai",
      cell: ({ row }) => (
        <p className="text-start">{formatDate(row.original.start_date)}</p>
      ),
    },
    {
      accessorKey: "end_date",
      header: "Tanggal Selesai",
      cell: ({ row }) => (
        <p className="text-start">{formatDate(row.original.end_date)}</p>
      ),
    },
    {
      accessorKey: "description",
      header: "Deskripsi",
      cell: ({ row }) => (
        <p className="text-start">{row.original.description}</p>
      ),
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-1 justify-content-center">
          <Button
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Ubah"
            onClick={() =>
              setModalState((prev) => ({
                ...prev,
                open: true,
                id: row.original.id,
                state: "edit",
              }))
            }
          >
            <i className="las la-edit text-success text-lg pointer" />
          </Button>
          <Link
            href={`/settings/curriculum-year/${row.original.id}/subjects`}
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Detail"
          >
            <i className="ri-eye-line text-primary text-lg"></i>
          </Link>

          <Button
            title="Hapus"
            className="bg-transparent border-0 text-black p-0"
            onClick={() => handleDelete(row.original.id)}
          >
            <i className="mdi mdi mdi-trash-can-outline fs-4 text-danger" />
          </Button>
        </div>
      ),
    },
  ];

  return { columns };
};
