"use client";

import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";

import { DeleteIcon } from "@/components/icons/delete";
import { EditIcon } from "@/components/icons/edit";
import { VisibilityIcon } from "@/components/icons/visibility";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { useModalContext } from "@/lib/hooks/use-modal";
import { useWeekDaysOptions } from "@/lib/hooks/use-days";

export const useColumnClassScheduleTemplate = (isDetail = false) => {
  const { setModalConfirmationState } = useModalConfirmationContext();
  const { setModalState } = useModalContext();

  const dayOptions = useWeekDaysOptions();

  const handleDelete = (id: string) => {
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "confirm",
      message: "Hapus Jadwal",
      id: id,
    }));
  };
  const columns: ColumnDef<ClassScheduleTemplate>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return row.index + 1;
      },
    },
    {
      accessorKey: "day_name",
      header: () => "Hari",
      cell: ({ row }) => {
        return (
          <p className="text-start">
            {
              dayOptions.find((data) => data.value === row.original.day_name)
                ?.label
            }
          </p>
        );
      },
    },
    {
      accessorKey: "start_time",
      header: () => "Jam Mulai",
      cell: ({ row }) => {
        return <p className="text-start">{row.original.start_time}</p>;
      },
    },
    {
      accessorKey: "end_time",
      header: () => "Jam Selesai",
      cell: ({ row }) => {
        return <p className="text-start">{row.original.end_time}</p>;
      },
    },
    {
      accessorKey: "type_of_meeting",
      header: () => "Jenis Pertemuan",
      cell: ({ row }) => {
        return (
          <p className="text-start text-capitalize">
            {row.original.type_of_meeting}
          </p>
        );
      },
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-1 justify-content-center">
          {!isDetail ? (
            <>
              <Button
                className="bg-transparent border-0 text-black p-0 fs-4"
                title="Ubah"
                onClick={() =>
                  setModalState((prev) => ({
                    ...prev,
                    open: true,
                    state: "edit",
                    id: row.original.id,
                  }))
                }
              >
                <EditIcon height="20" width="20" />
              </Button>
              <Button
                className="bg-transparent border-0 text-black p-0 fs-4"
                title="Detail"
                onClick={() =>
                  setModalState((prev) => ({
                    ...prev,
                    open: true,
                    state: "detail",
                    id: row.original.id,
                  }))
                }
              >
                <VisibilityIcon />
              </Button>
              <Button
                title="Hapus"
                className="bg-transparent border-0 text-black p-0"
                onClick={() => handleDelete(row.original.id)}
              >
                <DeleteIcon height="20" width="20" />
              </Button>
            </>
          ) : null}
        </div>
      ),
    },
  ];

  return { columns };
};
