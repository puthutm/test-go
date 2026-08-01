"use client";
import { ColumnDef } from "@tanstack/react-table";

// import hook
import { getHourAndMinute } from "@/lib/utils/format-date";
import { useWeekDaysOptions } from "@/lib/hooks/use-days";

interface iColumnsParams {
  (): { columns: ColumnDef<WeeklySchedule>[] };
}

const useColumnWeeklySchedule: iColumnsParams = () => {
  const days = useWeekDaysOptions();
  const columns: ColumnDef<WeeklySchedule>[] = [
    {
      header: "No",
      accessorKey: "no",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return <p className="m-0 p-0 text-start">{row.index + 1}</p>;
      },
    },

    {
      header: "Hari",
      accessorKey: "day_name",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {days.find((el) => el.value === row.original.day_name)?.label}
          </p>
        );
      },
    },
    {
      header: "Jam Mulai",
      accessorKey: "start_time",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {getHourAndMinute(row.original?.start_time)}
          </p>
        );
      },
    },
    {
      header: "Jam Selesai",
      accessorKey: "end_time",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {getHourAndMinute(row.original?.end_time)}
          </p>
        );
      },
    },
    {
      header: "Jenis Pertemuan",
      accessorKey: "jenis_pertemuan",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.type_of_meeting ?? "-"}
          </p>
        );
      },
    },
    // {
    //   header: "Metode Pembayaran",
    //   accessorKey: "metode_pembelajaran",
    //   enableColumnFilter: false,
    //   cell: ({ row }) => {
    //     return (
    //       <p className="m-0 p-0 text-center">
    //         {row.original?. ?? "-"}
    //       </p>
    //     );
    //   },
    // },
  ];

  return { columns };
};

export default useColumnWeeklySchedule;
