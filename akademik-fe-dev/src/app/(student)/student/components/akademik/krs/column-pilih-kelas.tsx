"use client";

import { ColumnDef } from "@tanstack/react-table";

interface PilihKelasColumnProps {
    savedKrsItems: KrsItem[];
    onTakeClass: (classId: string) => void;
}

export const usePilihKelasColumns = ({ savedKrsItems, onTakeClass }: PilihKelasColumnProps) => {
    const columns: ColumnDef<KrsAvailableClass>[] = [
        {
            header: "Kode",
            accessorKey: "class_code",
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
            cell: () => <div className="text-center">-</div>,
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
            header: "Kuota",
            accessorKey: "quota_text",
            cell: ({ getValue }) => (
                <div className="text-center">{getValue() as string}</div>
            ),
        },
        {
            id: "pilih_mk",
            header: "Pilih MK",
            cell: ({ row }) => {
                const cls = row.original;
                const isTaken = cls.button_state === "taken" ||
                    savedKrsItems.some((s) => s.class_id === cls.class_id);
                const isFull = cls.used_quota >= cls.capacity;

                if (isFull && !isTaken) {
                    return (
                        <div className="d-flex justify-content-center">
                            <button
                                className="btn text-white d-flex align-items-center justify-content-center p-3"
                                style={{ width: "100%", height: "40px", backgroundColor: "#dc3545" }}
                                disabled
                            >
                                Penuh
                            </button>
                        </div>
                    );
                }

                if (isTaken) {
                    return (
                        <div className="d-flex justify-content-center">
                            <button
                                className="btn text-white d-flex align-items-center justify-content-center p-3"
                                style={{ width: "100%", height: "40px", backgroundColor: "#6B9DC2" }}
                                disabled
                            >
                                Terambil
                            </button>
                        </div>
                    );
                }

                return (
                    <div className="d-flex justify-content-center">
                        <button
                            className="btn text-white d-flex align-items-center justify-content-center p-3"
                            style={{ width: "100%", height: "40px", backgroundColor: "#10487A" }}
                            onClick={() => onTakeClass(cls.class_id)}
                        >
                            Ambil
                        </button>
                    </div>
                );
            },
        },
    ];

    return { columns };
};
