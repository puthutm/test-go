import { useModalContext } from "@/lib/hooks/use-modal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";
import { EditIcon } from "@/components/icons/edit";
import { DeleteIcon } from "@/components/icons/delete";

export const useGradeScaleColumns = () => {
  const { setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const handleDelete = async (id: string) => {
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "confirm",
      message: "Hapus Skala Nilai",
      id: id,
    }));
  };

  const columns: ColumnDef<IGradeScale>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return <p className="text-start">{row.index + 1}</p>;
      },
    },
    {
      accessorKey: "study_program_name",
      header: "Program Studi",
      cell: ({ row }) => (
        <p className="text-start">{row.original.study_program_name ?? "-"}</p>
      ),
    },
    {
      accessorKey: "grade_name",
      header: "Nilai",
      cell: ({ row }) => (
        <p className="text-start">{row.original.grade_name ?? "-"}</p>
      ),
    },
    {
      accessorKey: "weight_value",
      header: "Bobot",
      cell: ({ row }) => (
        <p className="text-start">{row.original.weight_value ?? "-"}</p>
      ),
    },
    {
      accessorKey: "lower_value",
      header: "Batas Bawah",
      cell: ({ row }) => (
        <p className="text-start">{row.original.lower_value ?? "-"}</p>
      ),
    },
    {
      accessorKey: "upper_value",
      header: "Batas Atas",
      cell: ({ row }) => (
        <p className="text-start">{row.original.upper_value ?? "-"}</p>
      ),
    },
    {
      accessorKey: "description",
      header: "Deskripsi",
      cell: ({ row }) => (
        <p className="text-start">{row.original.description ?? "-"}</p>
      ),
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-1 justify-content-center">
          <Button
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Ubah"
            onClick={() => {
              setModalState((prev) => ({
                ...prev,
                open: true,
                id: row.original.id,
                state: "edit",
              }));
            }}
          >
            <EditIcon />
          </Button>

          {/* <Link href={`/settings/academic-period/detail/${row.original.id}`}> */}
          {/* <Button
            className="bg-transparent border-0 text-black p-0 fs-4"
            title="Detail"
            onClick={() => {
              setModalState((prev) => ({
                ...prev,
                open: !prev.open,
                id: row.original.id,
                state: "detail",
              }));
            }}
          >
            <i className="ri-eye-line text-primary text-lg"></i>
          </Button> */}
          {/* </Link> */}

          <Button
            title="Hapus"
            className="bg-transparent border-0 text-black p-0"
            onClick={() => handleDelete(row.original.id)}
          >
            <DeleteIcon />
          </Button>
        </div>
      ),
    },
  ];

  return { columns };
};
