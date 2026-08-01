"use client";
import { ColumnDef } from "@tanstack/react-table";

// import component
import Link from "next/link";
import { VisibilityIcon } from "@/components/icons/visibility";

// import hook
// import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { IDummyValueTable } from "../page";

interface iColumnsParams {
  (): { columns: ColumnDef<IDummyValueTable>[] };
}

const useColumnLecturerCollegeCurriculum: iColumnsParams = () => {
  // const { setModalConfirmationState } = useModalConfirmationContext();

  const columns: ColumnDef<IDummyValueTable>[] = [
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
      accessorKey: "curriculum_name",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">
            {row.original?.curriculum_name ?? "-"}
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
      accessorKey: "curriculum_name",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">
            {row.original?.curriculum_name ?? "-"}
          </p>
        );
      },
    },
    {
      header: "SKS",
      accessorKey: "sks",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original?.sks ?? "-"}</p>
        );
      },
    },
    {
      header: "Jenis MK",
      accessorKey: "type_subjects",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">
            {row.original?.type_subjects ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Prodi Pengampu",
      accessorKey: "teaching_programs",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">
            {row.original?.teaching_programs ?? "-"}
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
              {/*//! action edit */}
              {/* <a href="subject/edit/123"
                      className="bg-transparent border-0 text-black p-0">
                        <EditIcon color="#0AB39C" width='20' height='20'/>
                      </a> */}
              {/*//! action view ap */}
              <Link
                href="/curriculum/college-curriculum"
                className="bg-transparent border-0 text-black p-0"
              >
                <VisibilityIcon color="#2E3192" width="20" height="20" />
              </Link>
              {/*//! action delete */}
              {/* <Button
                        onClick={()=>{
                          setModalConfirmationState(()=>({
                            open:true,
                            state:'confirm',
                            message:'hapus data Mata Kuliah',
                            id:row.original.id,
                          }))
                        }}
                        className="bg-transparent border-0 text-black p-0">
                        <DeleteIcon color='#F06548' width='20' height='20'/>
                      </Button> */}
            </div>
          </>
        );
      },
    },
  ];

  return { columns };
};

export default useColumnLecturerCollegeCurriculum;
