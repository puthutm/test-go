"use client";
import { ColumnDef } from "@tanstack/react-table";

// import hook
import { IDummyValueTable } from "../page";

interface iColumnsParams {
  (): { columns: ColumnDef<IDummyValueTable>[] };
}

const useColumnDetailPresenceClass: iColumnsParams = () => {
  const columns: ColumnDef<IDummyValueTable>[] = [
    {
      header: "Nim",
      accessorKey: "nim",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return <p className="m-0 p-0 text-center">{row.original.nim ?? "-"}</p>;
      },
    },
    {
      header: "Nama Mahasiswa",
      accessorKey: "nama",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original?.nama ?? "-"}</p>
        );
      },
    },
    {
      header: "Keterangan",
      accessorKey: "keterangan",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">
            {row.original?.keterangan ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Status",
      accessorKey: "status",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original?.status ?? "-"}</p>
        );
      },
    },
  ];

  return { columns };
};

export default useColumnDetailPresenceClass;
