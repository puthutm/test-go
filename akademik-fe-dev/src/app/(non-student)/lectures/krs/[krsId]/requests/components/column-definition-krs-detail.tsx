"use client";

import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";

import { useModalContext } from "@/lib/hooks/use-modal";
import { useUpdateStatusKrs } from "@/services/api/lectures/krs/use-update-status-krs";

export const useTableDefinitionKrsDetailColumn = () => {
  const { mutateAsync, isPending } = useUpdateStatusKrs();
  const { setModalState } = useModalContext();
  const columns: ColumnDef<KrsItem>[] = [
    {
      id: "select",
      header: ({ table }) => (
        <input
          type="checkbox"
          checked={table.getIsAllPageRowsSelected()}
          onChange={table.getToggleAllPageRowsSelectedHandler()}
        />
      ),
      cell: ({ row }) => (
        <input
          type="checkbox"
          checked={row.getIsSelected()}
          disabled={!row.getCanSelect()}
          onChange={row.getToggleSelectedHandler()}
        />
      ),
    },
    {
      header: "Mata Kuliah",
      accessorKey: "subject_name_id",
    },
    {
      header: "Kelas",
      accessorKey: "class_name",
    },
    {
      header: "SKS",
      accessorKey: "total_sks",
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-1 justify-content-center">
          <Button
            className="btn-success border-0"
            title="Detail"
            onClick={async () =>
              await mutateAsync({
                krsItemId: row.original.krs_item_id,
                payload: {
                  item_status: "approved",
                },
              })
            }
            disabled={isPending}
          >
            Terima
          </Button>
          <Button
            disabled={isPending}
            className="btn-danger border-0"
            title="Detail"
            onClick={() =>
              setModalState((prev) => ({
                ...prev,
                open: true,
                id: row.original.krs_item_id,
              }))
            }
          >
            Tolak
          </Button>
        </div>
      ),
    },
  ];

  return { columns };
};
