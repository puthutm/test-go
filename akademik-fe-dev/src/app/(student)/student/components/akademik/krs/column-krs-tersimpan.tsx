"use client";

import { ColumnDef } from "@tanstack/react-table";

interface KrsTersimpanColumnProps {
    onDeleteKrs: (krsItemId: string, classId: string) => void;
}

export const useKrsTersimpanColumns = ({ onDeleteKrs }: KrsTersimpanColumnProps) => {
    const columns: ColumnDef<KrsItem>[] = [
        {
            header: "Kode",
            accessorKey: "subject_code",
            cell: ({ getValue }) => (
                <div className="text-center">{getValue() as string}</div>
            ),
        },
        {
            header: "Mata Kuliah",
            accessorKey: "subject_name",
            cell: ({ getValue }) => (
                <div className="text-center">{getValue() as string}</div>
            ),
        },
        {
            header: "Jadwal",
            accessorKey: "schedule",
            cell: ({ getValue }) => (
                <div className="text-center">{getValue() as string}</div>
            ),
        },
        {
            header: "Dosen Pengampu",
            accessorKey: "lecturer_names",
            cell: ({ getValue }) => (
                <div className="text-center">{getValue() as string}</div>
            ),
        },
        {
            header: "Kelas",
            accessorKey: "class_name",
            cell: ({ getValue }) => (
                <div className="text-center">{getValue() as string}</div>
            ),
        },
        {
            header: "SKS",
            accessorKey: "sks",
            cell: ({ getValue }) => (
                <div className="text-center">{getValue() as number}</div>
            ),
        },
        {
            id: "status",
            header: "Status",
            cell: ({ row }) => {
                const status = row.original.item_status?.toLowerCase();
                let bgColor = "#6c757d";
                let label = row.original.item_status;

                if (status === "approved") {
                    bgColor = "#66BB6A";
                    label = "Approved";
                } else if (status === "rejected") {
                    bgColor = "#dc3545";
                    label = "Rejected";
                } else if (status === "waiting") {
                    bgColor = "#9E9E9E";
                    label = "Waiting";
                }

                return (
                    <div className="d-flex align-items-center justify-content-center">
                        <span
                            className="badge"
                            style={{ backgroundColor: bgColor, fontSize: "12px", padding: "6px 12px" }}
                        >
                            {label}
                        </span>
                    </div>
                );
            },
        },
        {
            id: "aksi",
            header: "Aksi",
            cell: ({ row }) => {
                const status = row.original.item_status?.toLowerCase();
                return (
                    <div className="d-flex justify-content-center">
                        <button
                            className="btn text-white d-flex align-items-center justify-content-center p-3"
                            style={{
                                height: "40px",
                                backgroundColor: status === "approved" ? "#9E9E9E" : "#F06548",
                            }}
                            disabled={status === "approved"}
                            onClick={() => {
                                onDeleteKrs(row.original.krs_item_id, row.original.class_id);
                            }}
                        >
                            Hapus
                        </button>
                    </div>
                );
            },
        },
    ];

    return { columns };
};
