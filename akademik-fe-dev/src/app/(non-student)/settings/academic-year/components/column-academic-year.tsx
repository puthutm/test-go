"use client";

import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";

import { EditIcon } from "@/components/icons/edit";
import { VisibilityIcon } from "@/components/icons/visibility";
import { useModalContext } from "@/lib/hooks/use-modal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

export const useAcademicYearColumn = () => {
  const { setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const handleDelete = (id: string) => {
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "confirm",
      message: "Hapus Tahun Ajaran",
      id: id,
    }));
  };

  const columns: ColumnDef<AcademicYear>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return row.index + 1;
      },
    },
    {
      accessorKey: "years",
      header: "Tahun",
      cell: ({ row }) => <p className="text-start">{row.original.years}</p>,
    },
    {
      accessorKey: "name",
      header: "Nama",
      cell: ({ row }) => <p className="text-start">{row.original.name}</p>,
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-1 justify-content-center align-items-center">
          <Button
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Ubah"
            onClick={() =>
              setModalState({ open: true, state: "edit", id: row.original.id })
            }
          >
            <EditIcon height="20" width="20" />
          </Button>
          <Button
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Detail"
            onClick={() =>
              setModalState({
                open: true,
                state: "detail",
                id: row.original.id,
              })
            }
          >
            <VisibilityIcon />
          </Button>
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
