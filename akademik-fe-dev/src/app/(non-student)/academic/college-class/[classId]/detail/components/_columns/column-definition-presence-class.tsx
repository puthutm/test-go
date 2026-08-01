"use client";
import { ColumnDef } from "@tanstack/react-table";

// import hook
// import { VisibilityIcon } from "@/components/icons/visibility";
// import Link from "next/link";

// import { useParams } from "next/navigation";

import { getHourAndMinute } from "@/lib/utils/format-date";

interface iColumnsParams {
  (): { columns: ColumnDef<IClassAttendance>[] };
}

const useColumnPresenceClass: iColumnsParams = () => {
  // const params = useParams();

  const columns: ColumnDef<IClassAttendance>[] = [
    {
      header: "Sesi",
      accessorKey: "session",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start" style={{ color: "#495057" }}>
            {row.original.session ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Hari",
      accessorKey: "day_name",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start" style={{ color: "#495057" }}>
            {row.original?.day_name ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Waktu",
      accessorKey: "date",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start" style={{ color: "#495057" }}>
            {`${getHourAndMinute(
              row.original?.start_time ?? "-"
            )} s/d ${getHourAndMinute(row.original?.end_time ?? "-")}`}
          </p>
        );
      },
    },
    {
      header: "Rencana dan Realisasi Materi",
      accessorKey: "type_of_meeting",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <>
            <div className="position-relative d-flex flex-column">
              {/* title */}
              <p
                className="m-0 p-0 fw-bold"
                style={{ color: "#909090", fontSize: "13px" }}
              >
                {row.original.material_plan ?? "-"}
              </p>
              {/* subtitle */}
              <p
                className="m-0 p-0"
                style={{
                  color: "#495057",
                  fontSize: "13px",
                  fontWeight: "500",
                }}
              >
                {row.original.material_realization ?? "-"}
              </p>
            </div>
          </>
        );
      },
    },
    // {
    //   header: "Ruang",
    //   accessorKey: "ruang",
    //   enableColumnFilter: false,
    //   cell: ({ row }) => {
    //     return (
    //       <p className="m-0 p-0 text-center" style={{ color: "#495057" }}>
    //         {row.original?.ruang ?? "-"}
    //       </p>
    //     );
    //   },
    // },
    // {
    //   header: "Hadir",
    //   accessorKey: "hadir",
    //   enableColumnFilter: false,
    //   cell: ({ row }) => {
    //     return (
    //       <p className="m-0 p-0 text-center" style={{ color: "#495057" }}>
    //         {row.original?.hadir ?? "-"}
    //       </p>
    //     );
    //   },
    // },
    // {
    //   header: "%",
    //   accessorKey: "persen",
    //   enableColumnFilter: false,
    //   cell: ({ row }) => {
    //     return (
    //       <p className="m-0 p-0 text-center" style={{ color: "#495057" }}>
    //         {row.original?.persen ?? "-"}
    //       </p>
    //     );
    //   },
    // },
    // {
    //   header: "Presensi",
    //   enableSorting: false,
    //   cell: ({ row }) => {
    //     return (
    //       <div className="d-flex gap-2 justify-content-center align-items-center">
    //         {/*//! action view ap */}
    //         <Link
    //           href={`${params.catchAll?.length - 1}/detail-presence/${
    //             row.original.id
    //           }`}
    //           className="bg-transparent border-0 text-black p-0"
    //         >
    //           <VisibilityIcon color="#2E3192" width="20" height="20" />
    //         </Link>
    //       </div>
    //     );
    //   },
    // },
  ];

  return { columns };
};

export default useColumnPresenceClass;
