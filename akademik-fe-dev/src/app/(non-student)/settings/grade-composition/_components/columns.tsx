"use client";

import { useModalContext } from "@/lib/hooks/use-modal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";
import { DisabledByDefaultIcon } from "@/components/icons/disabled-by-default";
import { DoneIcon } from "@/components/icons/done";
import { EditIcon } from "@/components/icons/edit";
import { DeleteIcon } from "@/components/icons/delete";

export const useGradeCompositionColumns = () => {
  const { setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const handleDelete = async (id: string) => {
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "confirm",
      message: "Hapus Komposisi Nilai",
      id: id,
    }));
  };

  const columns: ColumnDef<IGradeComposition>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return <p className="text-start">{row.index + 1}</p>;
      },
    },
    {
      accessorKey: "academic_periode_name",
      header: "Periode Akademik",
    },
    {
      accessorKey: "value_element_name",
      header: "Elemen Nilai",
      cell: ({ row }) => (
        <p className="text-start">{row.original.value_element_name ?? "-"}</p>
      ),
    },
    {
      accessorKey: "percentage",
      header: "Persentase",
      cell: ({ row }) => (
        <p className="text-start">{row.original.percentage ?? "-"}%</p>
      ),
    },
    {
      accessorKey: "is_passing_requirement",
      header: "Memenuhi Persyaratan",
      cell: ({ row }) => (
        <p className="text-center">
          {row.original.is_passing_requirement ? (
            <DoneIcon />
          ) : (
            <DisabledByDefaultIcon />
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
            title="Ubah"
            onClick={() => {
              setModalState((prev) => ({
                ...prev,
                open: !prev.open,
                id: row.original.id,
                state: "edit",
              }));
            }}
          >
            <EditIcon />
          </Button>

          {/* <Link href={`/settings/academic-period/detail/${row.original.id}`}> */}
          {/* <Button
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Detail"
            onClick={() => {
              setModalState((prev) => ({
                ...prev,
                open: !prev.open,
                id: row.original.id,
                state: "detail",
              }));
            }}
          >
            <i className="ri-eye-line text-primary text-lg"></i>
          </Button> */}
          {/* </Link> */}

          <Button
            title="Hapus"
            className="bg-transparent border-0 text-black p-0"
            onClick={() => handleDelete(row.original.id)}
          >
            <DeleteIcon />
          </Button>
        </div>
      ),
    },
  ];

  return { columns };
};
