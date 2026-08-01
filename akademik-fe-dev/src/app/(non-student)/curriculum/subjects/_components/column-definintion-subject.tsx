"use client";
import { ColumnDef } from "@tanstack/react-table";

// import component
import { Button } from "reactstrap";
import Link from "next/link";
import { DeleteIcon } from "@/components/icons/delete";
import { EditIcon } from "@/components/icons/edit";
import { VisibilityIcon } from "@/components/icons/visibility";

// import hook
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

interface iColumnsParams {
  (): { columns: ColumnDef<ILectureSubjects>[] };
}

const useColumnLecturerSubject: iColumnsParams = () => {
  const { setModalConfirmationState } = useModalConfirmationContext();

  const columns: ColumnDef<ILectureSubjects>[] = [
    {
      header: "No",
      accessorKey: "no",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return <p className="m-0 p-0 text-start">{row.index + 1}</p>;
      },
    },
    {
      header: "Kur .",
      accessorKey: "curriculum_name",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.curriculum_year_name ?? "-"}
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
          <p className="m-0 p-0 text-start">{row.original?.code ?? "-"}</p>
        );
      },
    },
    {
      header: "Nama Mata Kuliah",
      accessorKey: "curriculum_name",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.name_id ?? "-"}</p>
        );
      },
    },
    {
      header: "SKS",
      accessorKey: "sks",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">{row.original?.total_sks ?? "-"}</p>
        );
      },
    },
    {
      header: "Jenis MK",
      accessorKey: "type_subjects",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-start">
            {row.original?.course_group_name ?? "-"}
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
            {row.original?.supporting_lecturers?.map(
              (data) =>
                `${data.lecturer_front_title} ${data.lecturer_name} ${data.lecturer_back_title}`
            ) ?? "-"}
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
              <Link
                href={`/settings/subject/${row.original.id}/edit`}
                className="bg-transparent border-0 text-black p-0"
              >
                <EditIcon color="#0AB39C" width="20" height="20" />
              </Link>
              {/*//! action view ap */}
              <Link
                href={`subjects/${row.original.id}/detail?tab=subject`}
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
                    message: "hapus data Mata Kuliah",
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

export default useColumnLecturerSubject;
