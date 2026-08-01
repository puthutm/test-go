"use client";

import Image from "next/image";
import React from "react";
import { Button, Modal, ModalBody } from "reactstrap";

import deleteImage from "@/assets/images/delete-emot.svg";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

export const ModalErrorConfirmation: React.FC = ({}) => {
  const { modalConfirmationState, setModalConfirmationState } =
    useModalConfirmationContext();

  const handleToggleModal = () => {
    setModalConfirmationState((prev) => ({
      ...prev,
      open: false,
    }));
  };

  return (
    <Modal
      isOpen={
        modalConfirmationState.open && modalConfirmationState.state === "failed"
      }
      centered
      className="mx-auto"
      style={{ maxWidth: "400px" }}
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
              <h3 className="fw-semibold text-black ">Gagal</h3>
              <p
                className="text-muted"
                style={{ fontSize: "14px", overflowX: "auto", width: "350px" }}
              >
                {modalConfirmationState.message}
              </p>
            </div>
            <div className="mx-auto">
              <Button
                type="button"
                className="btn-success waves-effect waves-light me-2"
                onClick={handleToggleModal}
                style={{ width: "100px" }}
              >
                Oke
              </Button>
            </div>
          </div>
        </div>
      </ModalBody>
    </Modal>
  );
};
