"use client";

import { useState } from "react";
import {
  Button,
  Input,
  Label,
  Modal,
  ModalBody,
  Spinner,
} from "reactstrap";

import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { importStudent } from "@/services/api/portal/academic/import-student";

interface ModalImportStudentProps {
  isOpen: boolean;
  toggle: () => void;
}

export const ModalImportStudent = ({ isOpen, toggle }: ModalImportStudentProps) => {
  const { setModalConfirmationState } = useModalConfirmationContext();
  const [file, setFile] = useState<File | null>(null);
  const [loading, setLoading] = useState(false);

  const handleToggleModal = () => {
    toggle();
    setFile(null);
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      setFile(e.target.files[0]);
    }
  };

  const onSubmit = async () => {
    if (!file) return;

    setLoading(true);
    const formData = new FormData();
    formData.append("file", file);

    try {
      const response = await importStudent(formData);

      if (response.error) {
        setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "failed",
          message: response.message,
        }));
      } else {
        setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          message: "Import data mahasiswa berhasil",
          state: "success",
        }));
        handleToggleModal();
      }
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "failed",
        message: error.toString(),
      }));
    } finally {
      setLoading(false);
    }
  };

  return (
    <Modal isOpen={isOpen} toggle={handleToggleModal} centered>
      <div className="d-flex justify-content-between align-items-center pt-3 px-3">
        <p className="fs-4 fw-semibold mb-0 text-black">Import Data Mahasiswa</p>
        <Button className="bg-white border-0 p-0" onClick={handleToggleModal}>
          <i className="ri-close-fill text-black fs-4"></i>
        </Button>
      </div>
      <ModalBody>
        <div className="mb-3">
          <Label htmlFor="file" className="form-label">
            Pilih File CSV
          </Label>
          <Input
            type="file"
            id="file"
            accept=".csv"
            onChange={handleFileChange}
          />
          <small className="text-muted mt-1 d-block">
            Format file harus .csv
          </small>
        </div>

        <div className="d-flex justify-content-end mt-4">
          <Button
            type="button"
            className="btn-light waves-effect waves-light me-2"
            onClick={handleToggleModal}
          >
            Batal
          </Button>
          <Button
            disabled={!file || loading}
            color="primary"
            onClick={onSubmit}
          >
            {loading ? <Spinner size="sm" /> : "Simpan"}
          </Button>
        </div>
      </ModalBody>
    </Modal>
  );
};
