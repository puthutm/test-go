"use client";
import { ColumnDef } from "@tanstack/react-table";

interface iColumnsParams {
  (): { columns: ColumnDef<IClassParticipant>[] };
}

const useColumnClassParticipants: iColumnsParams = () => {
  const columns: ColumnDef<IClassParticipant>[] = [
    {
      header: "Nim",
      accessorKey: "student_nim",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original.student_nim ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Nama Mahasiswa",
      accessorKey: "student_name",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.student_name ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Program Studi",
      accessorKey: "study_program_name",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.study_program_name ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Angkatan",
      accessorKey: "angkatan",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.year_of_entry ?? "-"}
          </p>
        );
      },
    },
  ];

  return { columns };
};

export default useColumnClassParticipants;
