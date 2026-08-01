import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";

import { VisibilityIcon } from "@/components/icons/visibility";
import { ModeEditIcon } from "@/components/icons/mode-edit";

export const useWeeklyScheduleColumnTable = () => {
  const columns: ColumnDef<any>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return row.index + 1;
      },
    },
    {
      accessorKey: "days",
      header: "Hari",
      cell: ({ row }) => <p className="text-start">{row.original.days}</p>,
    },
    {
      accessorKey: "start_hour",
      header: "Jam Mulai",
      cell: ({ row }) => (
        <p className="text-start">{row.original.start_hour}</p>
      ),
    },
    {
      accessorKey: "end_hour",
      header: "Jam Selesai",
      cell: ({ row }) => <p className="text-start">{row.original.end_hour}</p>,
    },
    {
      accessorKey: "type_meeting",
      header: "Jenis Pertemuan",
      cell: ({ row }) => (
        <p className="text-start">{row.original.type_meeting}</p>
      ),
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: () => (
        <div className="d-flex gap-1 justify-content-center">
          <Button
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Ubah"
          >
            <ModeEditIcon />
          </Button>
          <Button
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Detail"
          >
            <VisibilityIcon />
          </Button>
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
