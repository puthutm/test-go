"use client";

import Image from "next/image";
import { Button, Modal, ModalBody } from "reactstrap";

import smileImage from "@/assets/images/smile-emote.svg";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

export const ModalSuccessConfirmation: React.FC = () => {
  const { modalConfirmationState, setModalConfirmationState } =
    useModalConfirmationContext();

  const handleToogleModal = () => {
    setModalConfirmationState((prev) => ({
      ...prev,
      open: !prev.open,
      state: "success",
    }));
  };

  return (
    <Modal
      isOpen={
        modalConfirmationState.open &&
        modalConfirmationState.state === "success"
      }
      centered
      style={{ width: "400px" }}
    >
      <ModalBody>
        <div className="d-flex justify-content-center items-content-center p-4">
          <div className="d-flex flex-column gap-2 ">
            <Image
              src={smileImage}
              width={120}
              height={120}
              alt="Icon"
              className="mx-auto"
            />
            <div className="d-flex flex-column gap-1 text-center">
              <h3 className="fw-semibold text-black ">Berhasil</h3>
              <p className="text-muted">{modalConfirmationState.message}</p>
            </div>
            <div className="mx-auto">
              <Button
                type="button"
                className="btn-success waves-effect waves-light me-2"
                onClick={handleToogleModal}
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
