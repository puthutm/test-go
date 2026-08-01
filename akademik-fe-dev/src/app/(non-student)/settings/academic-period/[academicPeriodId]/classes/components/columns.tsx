import { DeleteIcon } from "@/components/icons/delete";
import { EditIcon } from "@/components/icons/edit";
import { VisibilityIcon } from "@/components/icons/visibility";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";
import { Button } from "reactstrap";

export const useAcademicPeriodClassesColumns = (academicPeriodId: string) => {
  const { setModalConfirmationState } = useModalConfirmationContext();

  const handleDelete = (id: string) => {
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "confirm",
      message: "Hapus Jadwal",
      id: id,
    }));
  };
  const columns: ColumnDef<Class>[] = [
    // code, name, subject, capacity, participant, lecturer, start_date, end_date, created_at, updated_at
    // all date should be formatted to id-ID (ex: 14 Maret 2025)
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
      cell: ({ row }) => <p className="text-start">{row.original.code}</p>,
    },
    {
      accessorKey: "name",
      header: "Nama",
      cell: ({ row }) => <p className="text-start">{row.original.name}</p>,
    },
    {
      accessorKey: "subject_name_id",
      header: "Mata Kuliah",
      cell: ({ row }) => (
        <p className="text-start">{row.original.subject_name_id}</p>
      ),
    },
    {
      accessorKey: "capacity",
      header: "Kapasitas",
      cell: ({ row }) => <p className="text-start">{row.original.capacity}</p>,
    },
    {
      accessorKey: "participant",
      header: "Peserta",
      cell: ({ row }) => (
        <p className="text-start">{row.original.total_participant}</p>
      ),
    },
    {
      accessorKey: "lecturer_name",
      header: "Dosen",
      cell: ({ row }) => (
        <p className="text-start">{row.original.lecturer_name}</p>
      ),
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex gap-1 justify-content-center">
          <Link
            href={`/settings/academic-period/${academicPeriodId}/classes/${row.original.id}/edit`}
          >
            <Button
              className="bg-transparent border-0 text-black p-0 fs-4"
              title="Ubah"
            >
              <EditIcon height="20" width="20" />
            </Button>
          </Link>
          <Link
            href={`/settings/academic-period/${academicPeriodId}/classes/${row.original.id}/detail`}
          >
            <Button
              className="bg-transparent border-0 text-black p-0 fs-4"
              title="Detail"
            >
              <VisibilityIcon />
            </Button>
          </Link>

          <Button
            title="Hapus"
            className="bg-transparent border-0 text-black p-0"
            onClick={() => handleDelete(row.original.id)}
          >
            <DeleteIcon height="20" width="20" />
          </Button>
        </div>
      ),
    },
  ];

  return { columns };
};
