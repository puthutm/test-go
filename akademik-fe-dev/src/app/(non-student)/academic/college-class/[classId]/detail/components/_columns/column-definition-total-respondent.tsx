"use client";
import { ColumnDef } from "@tanstack/react-table";

// import hook
import { IDummyValueTable } from "../_sections/section-questionnaire-recap";

interface iColumnsParams {
  (): { columns: ColumnDef<IDummyValueTable>[] };
}

const useColumnTotalRespondent: iColumnsParams = () => {
  const columns: ColumnDef<IDummyValueTable>[] = [
    {
      accessorKey: "no",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return <p className="m-0 p-0 text-start">{row.index + 1}</p>;
      },
    },
    {
      header: "Soal",
      accessorKey: "soal",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.soal ?? "-"}</p>
        );
      },
    },
  ];

  return { columns };
};

export default useColumnTotalRespondent;
