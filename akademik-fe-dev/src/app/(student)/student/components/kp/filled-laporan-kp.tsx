// FilledLaporanKP.tsx
import React from "react";

interface FilledLaporanKPProps {
  data: {
    judul: string;
    sertifikat: string;
    laporan: string;
  };
  openEditModal: () => void;
}

export const FilledLaporanKP: React.FC<FilledLaporanKPProps> = ({ data }) => {
  return (
    <div className="row">
      <div className="col-md-6">
        <div className="mb-3">
          <label className="form-label text-muted small">
            Judul Kegiatan Kerja Praktik
          </label>
          <input
            type="text"
            className="form-control form-control-icon bg-white"
            disabled
            value={data.judul || ""}
          />
        </div>
      </div>
      <div className="col-md-6 mb-3">
        <div>
          <label className="form-label text-muted small">
            Sertifikat Pembicara 2
          </label>
          <div className="d-flex align-items-center justify-content-between gap-2">
            <input
              type="text"
              className="form-control form-control-icon bg-white"
              disabled
              value={data.sertifikat || ""}
            />
            <button className="btn btn-light" type="button" disabled>
              Lihat
            </button>
          </div>
          <div className="text-muted small">
            File dalam bentuk .pdf max 10mb.
          </div>
        </div>
      </div>
      <div className="col-md-6">
        <div className="mb-3">
          <label className="form-label text-muted small">Laporan KP</label>
          <div className="d-flex align-items-center justify-content-between gap-2">
            <input
              type="text"
              className="form-control form-control-icon bg-white"
              disabled
              value={data.laporan || ""}
            />
            <button className="btn btn-light" type="button" disabled>
              Lihat
            </button>
          </div>
          <div className="text-muted small">
            File dalam bentuk .pdf max 10mb.
          </div>
          <div className="text-end mt-1">
            <a href="#" className="text-primary small">
              Template Laporan
            </a>
          </div>
        </div>
      </div>
    </div>
  );
};
