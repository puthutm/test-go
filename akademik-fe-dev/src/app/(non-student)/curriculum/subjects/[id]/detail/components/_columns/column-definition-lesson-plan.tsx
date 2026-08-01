"use client";
import { ColumnDef } from "@tanstack/react-table";

import { Dispatch, SetStateAction } from "react";

// import component
import { Button } from "reactstrap";
import { DeleteIcon } from "@/components/icons/delete";
import { VisibilityIcon } from "@/components/icons/visibility";
// import { SaveIcon } from "@/components/icons/save"

// import hook
// import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation"
import { IDummyValueTable } from "../_sections/section-lesson-plan";

import { IModalDetailLessonPlan } from "../_sections/section-lesson-plan";
interface iColumnsParams {
  (setShowModalDetail: Dispatch<SetStateAction<IModalDetailLessonPlan>>): {
    columns: ColumnDef<IDummyValueTable>[];
  };
}

const useColumnLessonPlan: iColumnsParams = (setShowModalDetail) => {
  // const {setModalConfirmationState} = useModalConfirmationContext()

  const columns: ColumnDef<IDummyValueTable>[] = [
    {
      header: "Sesi",
      accessorKey: "sesi",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.sesi ?? "-"}</p>
        );
      },
    },
    {
      header: "Sub-CPK",
      accessorKey: "sub_cpk",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.sub_cpk ?? "-"}</p>
        );
      },
    },
    {
      header: "penilaian",
      accessorKey: "penilaian",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.penilaian ?? "-"}</p>
        );
      },
    },
    {
      header: "Metode Pembelajaran",
      accessorKey: "metode_pembelajaran",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.metode_pembelajaran ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Materi Pembelajaran",
      accessorKey: "materi_pembelajaran",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.materi_pembelajaran ?? "-"}
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
      header: "Action",
      enableSorting: false,
      cell: ({ row }: any) => {
        return (
          <>
            <div className="d-flex gap-2 justify-content-center align-items-center">
              {/*//! action view ap */}
              <Button
                onClick={() => {
                  setShowModalDetail(() => ({
                    status: true,
                    title: "Detail Rencana Pembelajaran",
                    data: row.original,
                  }));
                }}
                className="bg-transparent border-0 text-black p-0"
              >
                <VisibilityIcon color="#2E3192" width="20" height="20" />
              </Button>
              {/*//! action delete */}
              <Button
                // onClick={()=>{
                //   setModalConfirmationState(()=>({
                //     open:true,
                //     state:'confirm',
                //     message:'hapus data Mata Kuliah',
                //     id:row.original.id,
                //   }))
                // }}
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

export default useColumnLessonPlan;
