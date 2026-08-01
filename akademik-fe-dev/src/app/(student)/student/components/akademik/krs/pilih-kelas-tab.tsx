"use client";

import styles from "@/styles/krs.module.css";
import DataTables from "@/components/ui/datatable";
import { usePilihKelasColumns } from "./column-pilih-kelas";

interface PilihKelasTabProps {
    classes: KrsAvailableClass[];
    savedKrsItems: KrsItem[];
    semesterName: string;
    onTakeClass: (classId: string) => void;
    isLoading?: boolean;
}

export default function PilihKelasTab({
    classes,
    savedKrsItems,
    semesterName,
    onTakeClass,
    isLoading,
}: PilihKelasTabProps) {
    const { columns } = usePilihKelasColumns({ savedKrsItems, onTakeClass });
    const dataTable = { data: classes };

    return (
        <div>
            <p className="p-3">{semesterName}</p>
            <DataTables
                columns={columns}
                data={dataTable}
                pageCount={1}
                pagination={{}}
                setPagination={() => { }}
                total={classes.length}
                isPaginate={false}
                isLoading={isLoading}
            />
            <div
                className={`h5 fw-semibold py-3 px-2 m-0 row d-flex justify-content-between`}
                style={{ backgroundColor: "#FFE91D33" }}
            >
                <p className="col">Total SKS</p>
                <p className="col text-center">3/20</p>
            </div>
            <div className={`d-grid gap-2 ${styles.mt_20}`}>
                <button className="btn btn-primary" type="button">
                    Validasi Krs
                </button>
            </div>
        </div>
    );
}
