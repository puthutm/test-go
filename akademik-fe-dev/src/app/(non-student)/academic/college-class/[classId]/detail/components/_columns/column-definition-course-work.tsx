"use client";
import { ColumnDef } from "@tanstack/react-table";

// import component
// import { Button } from "reactstrap";
// import { DeleteIcon } from "@/components/icons/delete";
// import { EditIcon } from "@/components/icons/edit";
import { VisibilityIcon } from "@/components/icons/visibility";
// import { Input } from "reactstrap";

// import hook
import { useModalContext } from "@/lib/hooks/use-modal";
interface iColumnsParams {
  (): { columns: ColumnDef<ICourseAssignment>[] };
}

const useColumnCourseWork: iColumnsParams = () => {
  const { setModalState } = useModalContext();

  const columns: ColumnDef<ICourseAssignment>[] = [
    {
      header: "No",
      accessorKey: "no",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return <p className="m-0 p-0 text-start">{row.index + 1}</p>;
      },
    },
    {
      header: "Judul",
      accessorKey: "title",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.title ?? "-"}</p>
        );
      },
    },
    {
      header: "Keterangan",
      accessorKey: "description",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.description ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Tugas Pertemuan Ke -",
      accessorKey: "session_schedule",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.session_schedule ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Awal Pengumpulan",
      accessorKey: "start_of_assignment_submission",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.time_to_open ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Akhir Pengumpulan",
      accessorKey: "end_of_assignment_submission",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.deadline_of_assignment_submission ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Jumlah Pengumpulan",
      accessorKey: "total_collect",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.total_collect ?? "-"}
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
              {/*//! action edit */}
              {/* <button className="bg-transparent border-0 text-black p-0">
                <EditIcon color="#0AB39C" width="20" height="20" />
              </button> */}
              {/*//! action view  */}
              <button
                onClick={() => {
                  setModalState({
                    open: true,
                    state: "detail",
                    id: row.original.id,
                  });
                }}
                className="bg-transparent border-0 text-black p-0"
              >
                <VisibilityIcon color="#2E3192" width="20" height="20" />
              </button>
              {/*//! action delete */}
              {/* <Button
                onClick={() => {
                  setModalConfirmationState(() => ({
                    open: true,
                    state: "confirm",
                    message: "hapus data Tugas Kuliah",
                    id: row.original.id,
                  }));
                }}
                className="bg-transparent border-0 text-black p-0"
              >
                <DeleteIcon color="#F06548" width="20" height="20" />
              </Button> */}
            </div>
          </>
        );
      },
    },
  ];

  return { columns };
};

export default useColumnCourseWork;
