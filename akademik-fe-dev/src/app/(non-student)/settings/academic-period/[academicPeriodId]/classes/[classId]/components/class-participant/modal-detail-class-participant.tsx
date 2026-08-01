"use client";

import { Button, Col, Modal, ModalBody, ModalHeader, Row } from "reactstrap";

import { useModalContext } from "@/lib/hooks/use-modal";
import { useGetClassParticipantById } from "@/services/api/settings/academic-period/class-participant/use-get-class-participant-by-id";

export const ModalDetailClassParticipant = ({
  classId,
}: {
  classId: string;
}) => {
  const { modalState, setModalState } = useModalContext();

  const { data } = useGetClassParticipantById(
    classId,
    modalState?.id as string
  );

  const toggleModal = () => {
    setModalState((prev) => ({
      ...prev,
      open: !modalState.open,
      id: undefined,
    }));
  };

  return (
    <Modal isOpen={modalState.open && modalState.state === "detail"} centered>
      <ModalHeader>Detail Peserta Kelas</ModalHeader>
      <ModalBody>
        <Row>
          <Col sm={12}>NIM</Col>
          <Col sm={12}>{data?.data.student_nim || "NIM"}</Col>
          <Col sm={12} className="mt-1">
            Nama
          </Col>
          <Col sm={12}>{data?.data.student_name || "Nama"}</Col>
        </Row>
        <div className="d-flex justify-content-end mt-3 gap-2">
          <Button color="primary" onClick={toggleModal}>
            Tutup
          </Button>
        </div>
      </ModalBody>
    </Modal>
  );
};
