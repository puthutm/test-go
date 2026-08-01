"use client";

import React from "react";
import { Button, Modal, ModalBody, Spinner } from "reactstrap";
import Image from "next/image";

import deleteImage from "@/assets/images/delete-emot.svg";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

interface ModalDeleteConfirmationProps {
  isLoading: boolean;
  onDelete: () => Promise<void>;
}

export const ModalDeleteConfirmation: React.FC<
  ModalDeleteConfirmationProps
> = ({ isLoading, onDelete }) => {
  const { modalConfirmationState, setModalConfirmationState } =
    useModalConfirmationContext();

  const handleToogleModal = () => {
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "confirm",
    }));
  };
  return (
    <Modal
      isOpen={
        modalConfirmationState.open &&
        modalConfirmationState.state === "confirm"
      }
      centered
      style={{ width: "400px" }}
    >
      <ModalBody>
        <div className="d-flex justify-content-center items-content-center p-4">
          <div className="d-flex flex-column gap-2 ">
            <Image
              src={deleteImage}
              width={120}
              height={120}
              alt="Icon"
              className="mx-auto"
            />
            <div className="d-flex flex-column gap-1 text-center">
              <h3 className="fw-semibold text-black ">
                {modalConfirmationState.message}?
              </h3>
              <p className="text-muted">Anda yakin ingin menghapus? </p>
            </div>
            <div className="mx-auto">
              <Button
                type="button"
                className="btn-light waves-effect waves-light me-2"
                onClick={handleToogleModal}
              >
                Tidak, Batal
              </Button>
              <Button
                className="btn btn-danger waves-effect waves-light"
                disabled={isLoading}
                onClick={onDelete}
              >
                {isLoading ? (
                  <>
                    <Spinner size={"sm"} /> Ya,Hapus
                  </>
                ) : (
                  "Ya, Hapus"
                )}
              </Button>
            </div>
          </div>
        </div>
      </ModalBody>
    </Modal>
  );
};
