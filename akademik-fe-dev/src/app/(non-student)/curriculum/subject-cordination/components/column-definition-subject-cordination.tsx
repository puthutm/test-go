"use client";
import { ColumnDef } from "@tanstack/react-table";

// import component
import Link from "next/link";
import { VisibilityIcon } from "@/components/icons/visibility";

// import hook

interface iColumnsParams {
  (): { columns: ColumnDef<ILectureSubjectsCordinator>[] };
}

const useColumnSubjectCordination: iColumnsParams = () => {
  const columns: ColumnDef<ILectureSubjectsCordinator>[] = [
    {
      header: "No",
      accessorKey: "no",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return <p className="m-0 p-0 text-center">{row.index ?? "-"}</p>;
      },
    },
    {
      header: "Kur .",
      accessorKey: "kur",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">
            {row.original?.curriculum_year_name ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Kode",
      accessorKey: "code",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original?.code ?? "-"}</p>
        );
      },
    },
    {
      header: "Nama Mata Kuliah",
      accessorKey: "nama_mata_kuliah",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original?.name_id ?? "-"}</p>
        );
      },
    },
    {
      header: "SKS",
      accessorKey: "sks",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">
            {row.original?.total_sks ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Jenis",
      accessorKey: "jenis",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">
            {row.original?.course_type_name ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Prasyarat",
      accessorKey: "prasyarat",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">
            {row.original?.supporting_lecturers?.map(
              (data) =>
                `${data.lecturer_front_title} ${data.lecturer_name} ${data.lecturer_back_title}`
            ) ?? "-"}
          </p>
        );
      },
    },

    {
      header: "Action",
      enableSorting: false,
      cell: () => {
        return (
          <>
            <div className="d-flex gap-2 justify-content-center align-items-center">
              {/*//! action view  */}
              <Link
                href="subject/detail/123"
                className="bg-transparent border-0 text-black p-0"
              >
                <VisibilityIcon color="#2E3192" width="20" height="20" />
              </Link>
            </div>
          </>
        );
      },
    },
  ];

  return { columns };
};

export default useColumnSubjectCordination;
