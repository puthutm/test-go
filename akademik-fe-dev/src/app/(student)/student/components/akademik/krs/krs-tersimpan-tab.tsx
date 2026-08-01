"use client";

import styles from "@/styles/krs.module.css";
import { Button } from "reactstrap";
import DataTables from "@/components/ui/datatable";
import { useKrsTersimpanColumns } from "./column-krs-tersimpan";

interface KrsTersimpanTabProps {
    savedKrsItems: KrsItem[];
    onDeleteKrs: (krsItemId: string, classId: string) => void;
    isLoading?: boolean;
}

export default function KrsTersimpanTab({
    savedKrsItems,
    onDeleteKrs,
    isLoading,
}: KrsTersimpanTabProps) {
    const { columns } = useKrsTersimpanColumns({ onDeleteKrs });
    const dataTable = { data: savedKrsItems };

    return (
        <div>
            <p className="p-3">KRS Tersimpan</p>
            <DataTables
                columns={columns}
                data={dataTable}
                pageCount={1}
                pagination={{}}
                setPagination={() => { }}
                total={savedKrsItems.length}
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

            <div
                className={`alert border-0 text-center ${styles.alert_gray} ${styles.m_20}`}
                role="alert"
            >
                <p> + Tambah Kelas</p>
            </div>

            <div className="d-flex justify-content-between mt-3 gap-3">
                <button
                    onClick={() => { }}
                    className="bg-transparent text-primary rounded px-3"
                    type="button"
                    style={{ border: "1px solid #10487A" }}
                >
                    <span>Batal</span>
                </button>
                <Button
                    color="primary"
                    className="d-flex flex-grow-0 justify-content-center align-items-center"
                >
                    <span>Update</span>
                </Button>
            </div>
        </div>
    );
}
