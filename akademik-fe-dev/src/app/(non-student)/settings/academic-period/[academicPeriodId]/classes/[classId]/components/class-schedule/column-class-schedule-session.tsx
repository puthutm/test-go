"use client";

import { ColumnDef } from "@tanstack/react-table";

import { useWeekDaysOptions } from "@/lib/hooks/use-days";
import { formatDate } from "@/lib/utils/format-date";

export const useColumnClassScheduleSession = () => {
  const dayOptions = useWeekDaysOptions();

  const columns: ColumnDef<ClassScheduleSession>[] = [
    {
      accessorKey: "session",
      header: () => "Sesi",
      cell: ({ row }) => {
        return row.original.session;
      },
    },
    {
      accessorKey: "day_name",
      header: () => "Hari",
      cell: ({ row }) => {
        return (
          <p className="text-start">
            {`${
              dayOptions.find((data) => data.value === row.original.day_name)
                ?.label
            }, ${formatDate(row.original.date)}`}
          </p>
        );
      },
    },
    {
      accessorKey: "times",
      header: () => "Waktu",
      cell: ({ row }) => {
        return (
          <p className="text-start">{`${row.original.start_time} s/d ${row.original.end_time}`}</p>
        );
      },
    },
    {
      accessorKey: "type_of_meeting",
      header: () => "Jenis",
      cell: ({ row }) => {
        return (
          <p className="text-start text-capitalize">
            {row.original.type_of_meeting}
          </p>
        );
      },
    },
    {
      accessorKey: "status",
      header: () => "Status",
    },
  ];

  return { columns };
};
