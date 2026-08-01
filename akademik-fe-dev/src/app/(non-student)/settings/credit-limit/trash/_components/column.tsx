"use client";

import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { useRestoreCreditLimit } from "@/services/api/settings/sks-limit/trash/restore";
import { ColumnDef } from "@tanstack/react-table";
import { Button } from "reactstrap";

export const useCreditLimitTrashColumn = () => {
  const { setModalConfirmationState } = useModalConfirmationContext();

  const { mutateAsync, isPending: isRestoreCreditLimitLoading } =
    useRestoreCreditLimit();

  const handleRestoreCreditLimit = async (id: string) => {
    await mutateAsync(id);
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      message: "Berhasil mengembalikan batas SKS",
      state: "success",
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
      cell: ({ row }) => <p className="text-start">{row.original.ips_min}</p>,
    },
    {
      accessorKey: "ips_max",
      header: "IPS Max",
      cell: ({ row }) => <p className="text-start">{row.original.ips_max}</p>,
    },
    {
      accessorKey: "sks_limit",
      header: "SKS Limit",
      cell: ({ row }) => <p className="text-start">{row.original.sks_limit}</p>,
    },
    {
      accessorKey: "id",
      header: "Aksi",
      cell: ({ row }) => (
        <div className="d-flex justify-content-center">
          <Button
            className="bg-transparent border-0 text-black p-0 me-1"
            title="Restore"
            onClick={async () =>
              await handleRestoreCreditLimit(row.original.id)
            }
            disabled={isRestoreCreditLimitLoading}
            key={row.original.id}
          >
            <i className="mdi mdi-backup-restore fs-4 text-secondary" />
          </Button>
        </div>
      ),
    },
  ];

  return { columns };
};
