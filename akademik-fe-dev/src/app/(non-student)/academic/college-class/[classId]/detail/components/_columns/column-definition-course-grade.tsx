"use client";
import { ColumnDef } from "@tanstack/react-table";

// import component

import { IDummyValueTable } from "../_sections/section-course-grades";
interface iColumnsParams {
  (): { columns: ColumnDef<IDummyValueTable>[] };
}
const useColumnDefinitionCourseGrade: iColumnsParams = () => {
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
      header: "NIM",
      accessorKey: "nim",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return <p className="m-0 p-0 text-start">{row.original?.nim ?? "-"}</p>;
      },
    },
    {
      header: "Nama",
      accessorKey: "nama",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.nama ?? "-"}</p>
        );
      },
    },
    {
      header: "Hadir",
      accessorKey: "hadir",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <>
            <div className="position-relative d-flex justify-content-center align-items-center">
              <p
                className="m-0 py-1 px-2  position-relative rounded-2  text-center"
                style={{
                  display: "inline-block",
                  background: "#F7B84B3B",
                  color: "#F7B84B",
                }}
              >
                {row.original?.hadir ?? "-"}
              </p>
            </div>
          </>
        );
      },
    },
    {
      header: "Tugas",
      accessorKey: "tugas",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original?.tugas ?? "-"}</p>
        );
      },
    },
    {
      header: "UTS",
      accessorKey: "uts",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original?.uts ?? "-"}</p>
        );
      },
    },
    {
      header: "UAS",
      accessorKey: "uas",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original?.uas ?? "-"}</p>
        );
      },
    },
    {
      header: "Kehadiran",
      accessorKey: "kehadiran",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">
            {row.original?.kehadiran ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Nilai",
      accessorKey: "nilai",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original?.nilai ?? "-"}</p>
        );
      },
    },
    {
      header: "Grade",
      accessorKey: "grade",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original?.grade ?? "-"}</p>
        );
      },
    },
    {
      header: "Lulus",
      accessorKey: "lulus",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original?.lulus ?? "-"}</p>
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
  ];

  return { columns };
};

export default useColumnDefinitionCourseGrade;
