// EmptyStateLaporanKP.tsx
import React from "react";
import { AddIcon } from "@/components/icons/add";
import { HighlightOffIcon } from "@/components/icons/highlight-off";

interface EmptyStateLaporanKPProps {
  openModal: () => void;
}

export const EmptyStateLaporanKP: React.FC<EmptyStateLaporanKPProps> = ({
  openModal,
}) => {
  return (
    <div className="d-flex flex-column align-items-center py-5">
      <div className="p-3 mb-3">
        <HighlightOffIcon />
      </div>
      <h5 className="text-center mb-2">
        Belum Ada Informasi Laporan Kerja Praktik
      </h5>
      <p className="text-center text-muted mb-4">
        Silakan isi form untuk melengkapi Laporan Kerja Praktik
      </p>
      <button className="btn btn-primary w-100" onClick={openModal}>
        <AddIcon color="white" /> Tambah Laporan Kerja Praktik
      </button>
    </div>
  );
};
