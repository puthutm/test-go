"use client";

import { ColumnDef } from "@tanstack/react-table";

interface CourseWithGrade {
    kode: string;
    mataKuliah: string;
    jadwal: string;
    dosenPengampu: string;
    kelas: string;
    sks: number | string;
    kuota: string;
    nilai: string;
    remidi?: boolean;
    pilihMK: boolean;
}

interface MkMengulangColumnProps {
    onToggleCourse: (kode: string) => void;
}

export const useMkMengulangColumns = ({ onToggleCourse }: MkMengulangColumnProps) => {
    const columns: ColumnDef<CourseWithGrade>[] = [
        {
            header: "Kode",
            accessorKey: "kode",
            cell: ({ getValue }) => (
                <div className="text-center">{getValue() as string}</div>
            ),
        },
        {
            header: "Mata Kuliah",
            accessorKey: "mataKuliah",
            cell: ({ getValue }) => (
                <div className="text-center">{getValue() as string}</div>
            ),
        },
        {
            id: "nilai",
            header: "Nilai",
            cell: ({ row }) => (
                <div className="text-center">
                    <span className="badge bg-danger">{row.original.nilai}</span>
                </div>
            ),
        },
        {
            header: "Jadwal",
            accessorKey: "jadwal",
            cell: ({ getValue }) => (
                <div className="text-center">{getValue() as string}</div>
            ),
        },
        {
            header: "Kelas",
            accessorKey: "kelas",
            cell: ({ getValue }) => (
                <div className="text-center">{getValue() as string}</div>
            ),
        },
        {
            header: "SKS",
            accessorKey: "sks",
            cell: ({ getValue }) => (
                <div className="text-center">{String(getValue())}</div>
            ),
        },
        {
            header: "Kuota",
            accessorKey: "kuota",
            cell: ({ getValue }) => (
                <div className="text-center">{getValue() as string}</div>
            ),
        },
        {
            id: "pilih_mk",
            header: "Pilih MK",
            cell: ({ row }) => {
                const course = row.original;
                return (
                    <div className="d-flex justify-content-center">
                        <button
                            className="btn text-white d-flex align-items-center justify-content-center p-3"
                            style={{
                                height: "40px",
                                backgroundColor: course.pilihMK ? "#10487A" : "#6c757d",
                            }}
                            onClick={() => onToggleCourse(course.kode)}
                        >
                            {course.pilihMK ? "Ambil" : "Batal"}
                        </button>
                    </div>
                );
            },
        },
    ];

    return { columns };
};
