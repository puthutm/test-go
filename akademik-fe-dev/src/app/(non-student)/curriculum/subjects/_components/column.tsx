import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";

export const useSubjectColumns = () => {
  const { setModalConfirmationState } = useModalConfirmationContext();
  const handleDelete = async (id: string) => {
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "confirm",
      message: "Hapus Mata Kuliah",
      id: id,
    }));
  };

  const columns: ColumnDef<Subject>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return row.index + 1;
      },
    },
    {
      accessorKey: "code",
      header: "Kode",
    },
    {
      accessorKey: "curriculum_year_name",
      header: "Tahun Ajaran",
    },
    {
      accessorKey: "name_id",
      header: "Mata Kuliah",
    },
    {
      accessorKey: "study_program_name",
      header: "Program Studi",
    },
    {
      accessorKey: "total_sks",
      header: "SKS",
    },
    {
      accessorKey: "course_type_name",
      header: "Jenis Mata Kuliah",
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-1 justify-content-center">
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
