"use client";

import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";

import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { EditIcon } from "@/components/icons/edit";
import { useModalContext } from "@/lib/hooks/use-modal";

export const useColumnCurriculumStudyProgram = () => {
  const { setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const handleDelete = (id: string) => {
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "confirm",
      message: "Hapus Mata Kuliah",
      id: id,
    }));
  };

  const columns: ColumnDef<CurriculumStudyProgramTable>[] = [
    {
      header: "No",
      accessorKey: "no",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return <p className="m-0 p-0 text-center">{row.index + 1}</p>;
      },
    },
    {
      header: "Kode",
      accessorKey: "subject_code",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original.subject_code}</p>
        );
      },
    },
    {
      header: "Mata Kuliah",
      accessorKey: "subject_name_id",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">{row.original.subject_name_id}</p>
        );
      },
    },
    {
      header: "SKS",
      accessorKey: "subject_total_sks",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <p className="m-0 p-0 text-center">
            {row.original.subject_total_sks}
          </p>
        );
      },
    },
    {
      header: "Status",
      accessorKey: "is_mandatory",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <div className="d-flex justify-content-center">
            {row.original.is_mandatory ? "Wajib" : "Pilihan"}
          </div>
        );
      },
    },
    {
      header: "Nilai Min",
      accessorKey: "grade_code",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return <p className="m-0 p-0 text-center">{row.original.grade_code}</p>;
      },
    },
    {
      header: "Prasayarat",
      accessorKey: "subject_prerequisites",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return (
          <div className="d-flex justify-content-center gap-1">
            {row.original.subject_prerequisites.map((data) => (
              <span
                className="rounded-3 px-3 py-1"
                style={{ border: "1px solid #DEE5EC", color: "#495057" }}
                key={data.id}
              >
                {data.subject_name_id}
              </span>
            ))}
          </div>
        );
      },
    },
    {
      header: "Konsentrasi Bidang Studi",
      accessorKey: "field_study_concentration_name",
      enableColumnFilter: false,
      cell: ({ row }) => {
        return row.original.field_study_concentration_name ? (
          <div className="text-center">
            <span
              className="rounded-3 px-3 py-1 text-center"
              style={{ border: "1px solid #DEE5EC", color: "#495057" }}
            >
              {row.original.field_study_concentration_name}
            </span>
          </div>
        ) : null;
      },
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-1 justify-content-center align-items-center">
          <Button
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Detail"
            onClick={() =>
              setModalState((prev) => ({
                ...prev,
                open: true,
                state: "edit",
                id: row.original.id,
              }))
            }
          >
            <EditIcon />
          </Button>
          <Button
            title="Hapus"
            className="bg-transparent border-0 text-black p-0"
            onClick={() => handleDelete(row.original.id)}
          >
            <i className="mdi mdi mdi-trash-can-outline fs-4 text-danger" />
          </Button>
        </div>
      ),
    },
  ];

  return { columns };
};
