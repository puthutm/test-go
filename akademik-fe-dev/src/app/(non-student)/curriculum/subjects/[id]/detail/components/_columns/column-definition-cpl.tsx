"use client";
import { ColumnDef } from "@tanstack/react-table";

// import component
import { Button } from "reactstrap";
import { SaveIcon } from "@/components/icons/save";

// import hook
// import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { IDummyValueTable } from "../_sections/section-cpl";

interface iColumnsParams {
  (): { columns: ColumnDef<IDummyValueTable>[] };
}

const useColumnCPL: iColumnsParams = () => {
  // const { setModalConfirmationState } = useModalConfirmationContext();

  const columns: ColumnDef<IDummyValueTable>[] = [
    // {
    //     header: 'No',
    //     accessorKey: "no",
    //     enableColumnFilter: false,
    //     cell: ({row}) => {
    //       return <p className="m-0 p-0 text-center">{row.index ?? '-'}</p>;
    //     },
    // },
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
      header: "Deskripsi",
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
      header: "Action",
      enableSorting: false,
      cell: () => {
        return (
          <>
            <div className="d-flex gap-2 justify-content-center align-items-center">
              {/*//! action view ap */}
              <Button
                style={{ background: "#007AFF" }}
                className=" text-black p-1 px-2"
              >
                <SaveIcon />
              </Button>
            </div>
          </>
        );
      },
    },
  ];

  return { columns };
};

export default useColumnCPL;
