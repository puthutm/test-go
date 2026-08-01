"use client";
import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";

import { EyeAkademikIcon } from "@/components/icons/eye-akademik";

import { formatDate } from "@/lib/utils/format-date";

export const useColumnDefThesisProposal = () => {
  const columns: ColumnDef<FinalProjectProposal>[] = [
    {
      header: "No",
      accessorKey: "no",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return <p className="m-0 p-0 text-center">{row.index + 1}</p>;
      },
    },
    {
      header: "Nama Mahasiswa",
      accessorKey: "student_name",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">
            {row.original?.student_name ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Prodi",
      accessorKey: "study_program_name",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">
            {row.original?.study_program_name ?? "-"}
          </p>
        );
      },
    },
    {
      header: "Judul Skripsi",
      accessorKey: "title_id",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original?.title_id ?? "-"}</p>
        );
      },
    },
    {
      header: "Topik",
      accessorKey: "topic",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original?.topic ?? "-"}</p>
        );
      },
    },
    {
      header: "Pembimbing",
      accessorKey: "dosen_pembimbing",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return <p className="m-0 p-0 text-center">{row.original?.id ?? "-"}</p>;
      },
    },
    {
      header: "Tgl Pengajuan",
      accessorKey: "date",
      enableColumnFilter: false,
      cell: ({ row }) => {
        const date = new Date(row.original?.date);
        return (
          <p className="m-0 p-0 text-center">
            {row.original?.date ? formatDate(date.toString()) : "-"}
          </p>
        );
      },
    },
    {
      header: "Status",
      accessorKey: "status",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p
            className="m-0 p-0 text-center p-1 fs-6 rounded-3 "
            style={{
              color: "#6CBE40",
              background: "#6CBE401A",
            }}
          >
            {row.original?.status ?? "-"}
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
              {/* <Link href="subjects/edit/123"
                      className="bg-transparent border-0 text-black p-0">
                        <EditIcon color="#0AB39C" width='20' height='20'/>
                      </Link> */}
              {/*//! action view ap */}
              <Link
                href={`thesis-proposal/${row.original.id}?tabs=proposal`}
                className="bg-transparent border-0 text-black p-0"
              >
                <EyeAkademikIcon color="#2E3192" width="20" height="20" />
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
