"use client";

import styles from "@/styles/krs.module.css";
import { useState } from "react";
import DataTables from "@/components/ui/datatable";
import { useMkMengulangColumns } from "./column-mk-mengulang";

const courseDataWithGrade = [
    {
        semester: "Semester 3",
        courses: [
            {
                kode: "KPIS502",
                mataKuliah: "Kalkulus",
                nilai: "D",
                jadwal: "Sen, 10.00-14.00 WIB",
                dosenPengampu: "Vika febri muliati, S.KOM, M.Kom",
                kelas: "402",
                sks: 3,
                kuota: "15/50",
                pilihMK: true,
                remidi: true,
            },
            {
                kode: "Text",
                mataKuliah: "Cloud Computing",
                nilai: "B-",
                jadwal: "Text",
                dosenPengampu: "Text",
                kelas: "Text",
                sks: "Text",
                kuota: "Text",
                pilihMK: false,
            },
        ],
    },
];

export default function MkMengulangTab() {
    const [courseWithNilai, setCourseWithNilai] = useState(courseDataWithGrade);

    const handleToggleCourse = (kode: string) => {
        const updatedCourseWithNilai = courseWithNilai.map((semester) => ({
            ...semester,
            courses: semester.courses.map((course) =>
                course.kode === kode ? { ...course, pilihMK: !course.pilihMK } : course
            ),
        }));
        setCourseWithNilai(updatedCourseWithNilai);
    };

    const remidiCourses = courseWithNilai
        .flatMap((s) => s.courses)
        .filter((c) => c.remidi);

    const { columns } = useMkMengulangColumns({ onToggleCourse: handleToggleCourse });
    const dataTable = { data: remidiCourses };

    return (
        <div>
            <p className="p-3">Mata Kuliah Mengulang</p>
            <DataTables
                columns={columns}
                data={dataTable}
                pageCount={1}
                pagination={{}}
                setPagination={() => { }}
                total={remidiCourses.length}
                isPaginate={false}
            />

            <div
                className={`h5 fw-semibold py-3 px-2 m-0 row d-flex justify-content-between`}
                style={{ backgroundColor: "#FFE91D33" }}
            >
                <p className="col">Total SKS Mengulang</p>
                <p className="col text-center">
                    {courseWithNilai
                        .flatMap((s) => s.courses)
                        .filter((c) => c.remidi && c.pilihMK)
                        .reduce((acc, c) => acc + (Number(c.sks) || 0), 0)}
                    /20
                </p>
            </div>

            <div className={`d-grid gap-2 ${styles.mt_20}`}>
                <button className="btn btn-primary" type="button">
                    Validasi KRS
                </button>
            </div>
        </div>
    );
}
