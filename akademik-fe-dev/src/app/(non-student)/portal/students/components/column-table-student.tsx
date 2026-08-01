// import { VisibilityIcon } from "@/components/icons/visibility";
// import { ModeEditIcon } from "@/components/icons/mode-edit";
import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";

export const useTableStudentColumns = () => {
  const columns: ColumnDef<any>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return row.index + 1;
      },
    },
    {
      accessorKey: "student_nim",
      header: "NIM",
      cell: ({ row }) => (
        <p className="text-start">
          {row.original.student_nim ?? row.original.nim}
        </p>
      ),
    },
    {
      accessorKey: "student_name",
      header: "Nama",
      cell: ({ row }) => (
        <p className="text-start">{row.original.student_name}</p>
      ),
    },
    {
      accessorKey: "study_program_name",
      header: "Program Studi",
      cell: ({ row }) => (
        <p className="text-start">{row.original.study_program_name}</p>
      ),
    },
    {
      accessorKey: "student_status",
      header: "Status",
      cell: ({ row }) => (
        <div className="d-flex justify-content-center">
          {row.original.student_status === "active" ? (
            <span
              className="p-1 rounded mx-auto"
              style={{
                backgroundColor: "#6CBE401A",
                color: "#6CBE40",
                textTransform: "capitalize",
              }}
            >
              {row.original.student_status}
            </span>
          ) : null}
          {row.original.student_status === "inactive" ? (
            <span
              className="p-1 rounded mx-auto"
              style={{ backgroundColor: "#F065481A", color: "#F06548" }}
            >
              {row.original.student_status}
            </span>
          ) : null}
        </div>
      ),
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: () => (
        <div className="d-flex gap-1 justify-content-center">
          {/* <Button
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
          </Button> */}
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
