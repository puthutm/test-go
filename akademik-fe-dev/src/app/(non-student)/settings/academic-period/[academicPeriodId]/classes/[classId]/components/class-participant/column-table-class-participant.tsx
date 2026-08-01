import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";

// import { ModeEditIcon } from "@/components/icons/mode-edit";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { VisibilityIcon } from "@/components/icons/visibility";
import { useModalContext } from "@/lib/hooks/use-modal";

export const useClassParticipantColumnTable = (isDetail = false) => {
  const { setModalConfirmationState } = useModalConfirmationContext();
  const { setModalState } = useModalContext();

  const handleDelete = (id: string) => {
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "confirm",
      message: "Hapus Peserta Kelas",
      id: id,
    }));
  };
  const columns: ColumnDef<ClassParticipant>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return row.index + 1;
      },
    },
    {
      accessorKey: "student_nim",
      header: "NIM",
      cell: ({ row }) => (
        <p className="text-start">{row.original.student_nim}</p>
      ),
    },
    {
      accessorKey: "student_name",
      header: "Nama",
      cell: ({ row }) => (
        <p className="text-start">{row.original.student_name}</p>
      ),
    },
    {
      accessorKey: "study_program_name",
      header: "Program Studi",
      cell: ({ row }) => (
        <p className="text-start">{row.original.study_program_name}</p>
      ),
    },
    {
      accessorKey: "year_of_entry",
      header: "Tahun Masuk",
      cell: ({ row }) => (
        <p className="text-start">{row.original.year_of_entry}</p>
      ),
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-1 justify-content-center">
          {/* <Button
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Ubah"
          >
            <ModeEditIcon />
          </Button> */}
          {!isDetail ? (
            <>
              <Button
                className="bg-transparent border-0 text-black p-0 fs-4"
                title="Detail"
                onClick={() => {
                  setModalState((prev) => ({
                    ...prev,
                    open: true,
                    id: row.original.id,
                    state: "detail",
                  }));
                }}
              >
                <VisibilityIcon />
              </Button>
              <Button
                title="Hapus"
                className="bg-transparent border-0 text-black p-0"
                onClick={() => handleDelete(row.original.id)}
              >
                <i className="mdi mdi mdi-trash-can-outline fs-4 text-danger" />
              </Button>
            </>
          ) : null}
        </div>
      ),
    },
  ];

  return { columns };
};
