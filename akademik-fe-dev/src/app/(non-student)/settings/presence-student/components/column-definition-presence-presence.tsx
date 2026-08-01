"use client";

import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";

import { ContentCopyIcon } from "@/components/icons/content-copy";
import { useModalContext } from "@/lib/hooks/use-modal";
import { Dispatch, SetStateAction } from "react";

export const useColumnDefinitionPresence = ({
  setStudyProgramId,
}: {
  setStudyProgramId: Dispatch<SetStateAction<string>>;
}) => {
  const { setModalState } = useModalContext();
  const columns: ColumnDef<Presences>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return row.index + 1;
      },
    },
    {
      accessorKey: "academic_periode_name",
      header: "Periode",
      cell: ({ row }) => (
        <p className="text-start">{row.original.academic_periode_name}</p>
      ),
    },
    {
      accessorKey: "study_program_name",
      header: "Program Studi",
      cell: ({ row }) => (
        <p className="text-start">
          {row.original.study_program_name
            ? row.original.study_program_name
            : "Semua Program Studi"}
        </p>
      ),
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-2 items-content-center justify-content-center">
          <Button
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Duplikasi"
            onClick={() => {
              setModalState((prev) => ({
                ...prev,
                open: true,
                id: row.original.academic_periode_id,
                state: "edit",
              }));
              setStudyProgramId(row.original.study_program_id);
            }}
          >
            <ContentCopyIcon />
          </Button>
          {/* <Link
            href={`/settings/presence/${row.original.id}/detail`}
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Detail"
          >
            <i className="ri-eye-line text-primary text-lg"></i>
          </Link> */}
        </div>
      ),
    },
  ];

  return { columns };
};
