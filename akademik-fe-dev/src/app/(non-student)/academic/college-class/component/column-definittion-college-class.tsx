"use client";
import { ColumnDef } from "@tanstack/react-table";

// import component
import Link from "next/link";
import { VisibilityIcon } from "@/components/icons/visibility";
import { DoneIcon } from "@/components/icons/done";

import { getHourAndMinute } from "@/lib/utils/format-date";
import { useWeekDaysOptions } from "@/lib/hooks/use-days";

// import hook
// import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

interface iColumnsParams {
  (): { columns: ColumnDef<ClassSchedule>[] };
}

const useColumnCollegeClass: iColumnsParams = () => {
  // const { setModalConfirmationState } = useModalConfirmationContext();
  const days = useWeekDaysOptions();
  const columns: ColumnDef<ClassSchedule>[] = [
    {
      header: "Kode",
      accessorKey: "code",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.code ?? "-"}</p>
        );
      },
    },
    {
      header: "Mata Kuliah",
      accessorKey: "subject_name_id",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.subject_name_id ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Program Studi",
      accessorKey: "study_program_id",
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
      header: "Kelas",
      accessorKey: "name",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.name ?? "-"}</p>
        );
      },
    },
    {
      header: "Jadwal Kelas",
      accessorKey: "name",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {`${
              row.original?.day_name === ""
                ? "-"
                : days.find((el) => el.value === row.original.day_name)?.label
            }, 
            ${getHourAndMinute(
              row.original?.start_time
            )} s/d ${getHourAndMinute(row.original?.end_time)}`}
          </p>
        );
      },
    },
    {
      header: "Nilai Dikunci",
      accessorKey: "nilai_dikunci",
      enableColumnFilter: false,
      cell: () => {
        return (
          <p className="m-0 p-0 text-center ">
            <DoneIcon />
          </p>
        );
      },
    },
    {
      header: "PJMK",
      accessorKey: "pjmk",
      enableColumnFilter: false,
      cell: () => {
        return (
          <p className="m-0 p-0 text-center ">
            <DoneIcon />
          </p>
        );
      },
    },
    {
      header: "Action",
      enableSorting: false,
      cell: ({ row }) => {
        return (
          <>
            <div className="d-flex gap-2 justify-content-center align-items-center">
              {/*//! action view ap */}
              <Link
                href={`/academic/college-class/${row.original.id}/detail?tab=class`}
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

export default useColumnCollegeClass;
