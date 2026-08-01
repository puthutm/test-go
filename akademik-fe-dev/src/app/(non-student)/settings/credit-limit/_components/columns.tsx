import { useModalContext } from "@/lib/hooks/use-modal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";
import { DeleteIcon } from "@/components/icons/delete";
import { EditIcon } from "@/components/icons/edit";

export const useCreditLimitColumns = () => {
  const { setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const handleDelete = async (id: string) => {
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "confirm",
      message: "Hapus Batas SKS",
      id: id,
    }));
  };

  const columns: ColumnDef<ISksLimit>[] = [
    {
      accessorKey: "id",
      header: () => "No.",
      cell: ({ row }) => {
        return <p className="text-start">{row.index + 1}</p>;
      },
    },
    {
      accessorKey: "ips_min",
      header: "IPS Min",
      cell: ({ row }) => (
        <p className="text-start">{row.original.ips_min ?? "-"}</p>
      ),
    },
    {
      accessorKey: "ips_max",
      header: "IPS Max",
      cell: ({ row }) => (
        <p className="text-start">{row.original.ips_max ?? "-"}</p>
      ),
    },
    {
      accessorKey: "sks_limit",
      header: "SKS Limit",
      cell: ({ row }) => (
        <p className="text-start">{row.original.sks_limit ?? "-"}</p>
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
                open: !prev.open,
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
