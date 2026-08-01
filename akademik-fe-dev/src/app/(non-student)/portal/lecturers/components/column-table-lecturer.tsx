import { VisibilityIcon } from "@/components/icons/visibility";
import { ModeEditIcon } from "@/components/icons/mode-edit";
import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";

export const useTableLecturersColumns = () => {
  const columns: ColumnDef<PortalLecturer>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return row.index + 1;
      },
    },
    {
      accessorKey: "nip",
      header: "NIP",
      cell: ({ row }) => <p className="text-start">{row.original.nip}</p>,
    },
    {
      accessorKey: "name",
      header: "Nama",
      cell: ({ row }) => <p className="text-start">{row.original.name}</p>,
    },
    {
      accessorKey: "nidn",
      header: "NIDN",
      cell: ({ row }) => <p className="text-start">{row.original.nidn}</p>,
    },
    {
      accessorKey: "study_program_name",
      header: "Program Studi",
      cell: ({ row }) => (
        <p className="text-start">{row.original.study_program_name}</p>
      ),
    },
    {
      accessorKey: "gender",
      header: "L/P",
      cell: ({ row }) => <p className="text-start">{row.original.gender}</p>,
    },
    {
      accessorKey: "phone",
      header: "No Hp",
      cell: ({ row }) => <p className="text-start">{row.original.phone}</p>,
    },
    {
      accessorKey: "email",
      header: "Email",
      cell: ({ row }) => <p className="text-start">{row.original.email}</p>,
    },
    {
      accessorKey: "status",
      header: "Status",
      cell: ({ row }) => (
        <div className="d-flex justify-content-center">
          <span
            className="p-1 rounded mx-auto"
            style={{ backgroundColor: "#6CBE401A", color: "#6CBE40" }}
          >
            {row.original.status}
          </span>
        </div>
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
