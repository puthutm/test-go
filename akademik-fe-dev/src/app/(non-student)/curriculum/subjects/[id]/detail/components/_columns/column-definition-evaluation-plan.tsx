"use client";
import { ColumnDef } from "@tanstack/react-table";

// import component
// import {Button} from "reactstrap"
// import { VisibilityIcon } from "@/components/icons/visibility"

// import hook
import { IDummyValueTable } from "../_sections/section-evaluation-plan";

interface iColumnsParams {
  (): { columns: ColumnDef<IDummyValueTable>[] };
}

const useColumnEvaluationPlan: iColumnsParams = () => {
  const columns: ColumnDef<IDummyValueTable>[] = [
    {
      header: "No",
      accessorKey: "no",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return <p className="m-0 p-0 text-start">{row.index + 1}</p>;
      },
    },
    {
      header: "Basis Evaluation",
      accessorKey: "basis_evaluation",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.basis_evaluation ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Komponent Evaluasi",
      accessorKey: "component_evaluation",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.component_evaluation ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Bobot (%)",
      accessorKey: "bobot",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.bobot ?? "-"}</p>
        );
      },
    },
    {
      header: "Deskripsi",
      accessorKey: "deskripsi_ind",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.deskripsi_ind ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Deskripsi (ENG)",
      accessorKey: "deskripsi_eng",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.deskripsi_eng ?? "-"}
          </p>
        );
      },
    },
  ];

  return { columns };
};

export default useColumnEvaluationPlan;
