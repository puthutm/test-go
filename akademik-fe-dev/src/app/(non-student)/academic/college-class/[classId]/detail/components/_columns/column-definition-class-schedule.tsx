"use client";
import { ColumnDef } from "@tanstack/react-table";
import { useParams } from "next/navigation";
// import component
import { Button } from "reactstrap";
import Link from "next/link";
import { DeleteIcon } from "@/components/icons/delete";
import { VisibilityIcon } from "@/components/icons/visibility";
// import { Input } from "reactstrap";

// import hook
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

import { getHourAndMinute } from "@/lib/utils/format-date";

interface iColumnsParams {
  (): { columns: ColumnDef<IClassScheduleSubDetail>[] };
}

const useColumnClassSchedule: iColumnsParams = () => {
  const { setModalConfirmationState } = useModalConfirmationContext();
  const params = useParams();
  const columns: ColumnDef<IClassScheduleSubDetail>[] = [
    // {
    //     header: ({ table }) => {
    //        table.getSelectedRowModel().rows.forEach((el)=>{
    //         console.log(el.original.id)
    //        })
    //         return (
    //       <input
    //       type="checkbox"
    //       checked={table.getIsAllRowsSelected()}
    //       onChange={table.getToggleAllRowsSelectedHandler()}
    //       />
    //     )
    //     },
    //   accessorKey: "no",
    //   enableColumnFilter: false,
    //   cell: ({ row }) => {

    //     return <input type="checkbox" checked={row.getIsSelected()} disabled={!row.getCanSelect()}
    //         onChange={row.getToggleSelectedHandler()}
    //     />;
    //   },
    // },
    {
      header: "Sesi",
      accessorKey: "session",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.session ?? "-"}</p>
        );
      },
    },
    {
      header: "Hari",
      accessorKey: "day_name",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.day_name ?? "-"}</p>
        );
      },
    },
    {
      header: "Waktu",
      accessorKey: "session_schedule",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {`${getHourAndMinute(
              row.original.start_time
            )} s/d ${getHourAndMinute(row.original.end_time)}`}
          </p>
        );
      },
    },
    {
      header: "Jenis",
      accessorKey: "type_of_meeting",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.type_of_meeting ?? "-"}
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
              {/*//! action view  */}
              <Link
                href={`/academic/college-class/${params.classId}/detail/${row.original.id}/detail-class-schedule`}
                className="bg-transparent border-0 text-black p-0"
              >
                <VisibilityIcon color="#2E3192" width="20" height="20" />
              </Link>
              {/*//! action delete */}
              <Button
                onClick={() => {
                  setModalConfirmationState(() => ({
                    open: true,
                    state: "confirm",
                    message: "hapus data jadwal kuliah",
                    id: row.original.id,
                  }));
                }}
                className="bg-transparent border-0 text-black p-0"
              >
                <DeleteIcon color="#F06548" width="20" height="20" />
              </Button>
            </div>
          </>
        );
      },
    },
  ];

  return { columns };
};

export default useColumnClassSchedule;
