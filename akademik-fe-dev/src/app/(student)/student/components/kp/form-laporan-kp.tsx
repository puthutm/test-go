// FormLaporanKP.tsx
import React, { useState, ChangeEvent } from "react";
import { Modal, ModalHeader, ModalBody, ModalFooter } from "reactstrap";
import { FileUploadIcon } from "@/components/icons/file-upload";

interface FormLaporanKPProps {
  isOpen: boolean;
  toggle: () => void;
  onSubmit: (data: {
    judul: string;
    laporan: File | null;
    sertifikat: File | null;
  }) => void;
  initialData?: {
    judul: string;
    laporan: string;
    sertifikat: string;
  };
}

export const FormLaporanKP: React.FC<FormLaporanKPProps> = ({
  isOpen,
  toggle,
  onSubmit,
  initialData,
}) => {
  const [judulKegiatan, setJudulKegiatan] = useState<string>(
    initialData?.judul || ""
  );
  const [laporanFile, setLaporanFile] = useState<File | null>(null);
  const [sertifikatFile, setSertifikatFile] = useState<File | null>(null);

  // File input handler
  const handleLaporanFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      setLaporanFile(e.target.files[0]);
    }
  };

  const handleSertifikatFileChange = (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      setSertifikatFile(e.target.files[0]);
    }
  };

  // Handle form submission
  const handleSubmit = () => {
    onSubmit({
      judul: judulKegiatan,
      laporan: laporanFile,
      sertifikat: sertifikatFile,
    });
    toggle();
  };

  return (
    <Modal isOpen={isOpen} toggle={toggle} centered>
      <ModalHeader
        toggle={toggle}
        className="border-2 border-bottom py-4 px-2 mx-4 "
      >
        Tambah Laporan KP
      </ModalHeader>
      <ModalBody>
        <form>
          <div className="mb-3">
            <label htmlFor="judulKegiatan" className="form-label">
              Judul Kegiatan Kerja Praktik
            </label>
            <input
              type="text"
              className="form-control border-light"
              id="judulKegiatan"
              placeholder="Text"
              value={judulKegiatan}
              onChange={(e) => setJudulKegiatan(e.target.value)}
            />
          </div>
          <div className="mb-3">
            <label htmlFor="laporanKP" className="form-label">
              Laporan KP
            </label>
            <div className="d-flex align-items-center justify-content-between gap-2">
              <input
                type="text"
                className="form-control form-control-icon border-light"
                id="laporanKP"
                placeholder="Upload File"
                readOnly
                value={
                  laporanFile ? laporanFile.name : initialData?.laporan || ""
                }
              />
              <label
                className="input-group-text border-0 btn btn-light"
                htmlFor="laporanFileInput"
              >
                <span className="d-flex align-items-center gap-1">
                  <FileUploadIcon />
                  Upload
                </span>
              </label>
              <input
                type="file"
                id="laporanFileInput"
                className="d-none"
                accept=".pdf"
                onChange={handleLaporanFileChange}
              />
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
          <div className="mb-3">
            <label htmlFor="sertifikatPembicara" className="form-label">
              Sertifikat Pembicara 2
            </label>
            <div className="d-flex align-items-center justify-content-between gap-2">
              <input
                type="text"
                className="form-control form-control-icon border-light"
                id="sertifikatPembicara"
                placeholder="Upload File"
                readOnly
                value={
                  sertifikatFile
                    ? sertifikatFile.name
                    : initialData?.sertifikat || ""
                }
              />
              <label
                className="input-group-text ml-4 border-0 btn btn-light"
                htmlFor="sertifikatFileInput"
              >
                <span className="d-flex align-items-center gap-1">
                  <FileUploadIcon />
                  Upload
                </span>
              </label>
              <input
                type="file"
                id="sertifikatFileInput"
                className="d-none"
                accept=".pdf"
                onChange={handleSertifikatFileChange}
              />
            </div>
            <div className="text-muted small">
              File dalam bentuk .pdf max 10mb.
            </div>
          </div>
        </form>
      </ModalBody>
      <ModalFooter>
        <button type="button" className="btn btn-light" onClick={toggle}>
          Batal
        </button>
        <button
          type="button"
          className="btn btn-primary"
          onClick={handleSubmit}
        >
          Submit
        </button>
      </ModalFooter>
    </Modal>
  );
};
