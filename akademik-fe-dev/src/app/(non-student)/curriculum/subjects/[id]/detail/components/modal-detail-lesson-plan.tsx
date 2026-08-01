"use client";
import React, { Dispatch, SetStateAction } from "react";
// import third pary component
import { Modal, ModalBody, Row, Button, Col } from "reactstrap";

// import component
import { CloseIcon } from "@/components/icons/close";
// import { SelectComponent } from '@/components/ui/select'
// import { FormErrorMessage } from '@/components/ui/form-error-message'

import { IModalDetailLessonPlan } from "./_sections/section-lesson-plan";

function ModalDetailLessonPlan({
  showModal,
  setShowModal,
}: {
  showModal: IModalDetailLessonPlan;
  setShowModal: Dispatch<SetStateAction<IModalDetailLessonPlan>>;
}) {
  // event handle close
  const handleCloseModal = () => {
    setShowModal(() => ({
      status: false,
      title: "Detail Rencana Pembelajaran",
      data: "",
    }));
  };

  return (
    <Modal
      isOpen={showModal.status}
      centered
      size="md"
      className="position-relative p-0"
      style={{ border: "0" }}
    >
      {/*//! modal header */}
      <section className="px-4">
        <section className="position-relative py-3 d-flex align-items-center justify-content-end gap-2 border-bottom border-3">
          <h2
            style={{ fontSize: 15, color: "#3A3A3A" }}
            className="m-0 p-0 fw-semibold flex-grow-1 w-100"
          >
            {showModal.title}
          </h2>

          <Button
            className="p-0"
            color={"transparent"}
            onClick={handleCloseModal}
          >
            <CloseIcon width="25" hanging={"25"} />
          </Button>
        </section>
      </section>

      {/*//! modal Body */}
      <ModalBody className="p-4">
        <section className="d-flex row-gap-2 flex-column">
          {/*//! col pertemuan */}
          <Row className="gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#909090" }}
              >
                Pertemuan Ke-{" "}
              </h3>
            </Col>
            <Col sm={12}>
              <p className="m-0 p-0 fw-medium" style={{ color: "#3A3A3A" }}>
                1
              </p>
            </Col>
          </Row>

          {/*//! col sub CPMK */}
          <Row className="gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#909090" }}
              >
                Sub-CMPK
              </h3>
            </Col>
            <Col sm={12}>
              <p className="m-0 p-0 fw-medium" style={{ color: "#3A3A3A" }}>
                Pengantar Interaksi Manusia dan Komputer
              </p>
            </Col>
          </Row>

          {/*//! col Metode Pembelajaran Luring */}
          <Row className="gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#909090" }}
              >
                indikator penilaian
              </h3>
            </Col>
            <Col sm={12}>
              <p className="m-0 p-0 fw-medium" style={{ color: "#3A3A3A" }}>
                Pengantar Interaksi Manusia dan Komputer
              </p>
            </Col>
          </Row>

          {/*//! col Metode Pembelajaran Luring */}
          <Row className="gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#909090" }}
              >
                Metode Pembelajaran Luring
              </h3>
            </Col>
            <Col sm={12}>
              <p className="m-0 p-0 fw-medium" style={{ color: "#3A3A3A" }}>
                text
              </p>
            </Col>
          </Row>

          {/*//! col Metode Pembelajaran Daring */}
          <Row className="gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#909090" }}
              >
                Metode Pembelajaran Daring
              </h3>
            </Col>
            <Col sm={12}>
              <p className="m-0 p-0 fw-medium" style={{ color: "#3A3A3A" }}>
                text
              </p>
            </Col>
          </Row>

          {/*//! col Materi Pembelajaran  - span */}
          <Row className="gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#909090" }}
              >
                Materi Pembelajaran -{" "}
                <span
                  className="font-italic fw-normal"
                  style={{ fontStyle: "italic" }}
                >
                  Digunakan untuk rencana perkuliahan per sesi
                </span>
              </h3>
            </Col>
            <Col sm={12}>
              <p className="m-0 p-0 fw-medium" style={{ color: "#3A3A3A" }}>
                text
              </p>
            </Col>
          </Row>

          {/*//! col Materi Pembelajaran (EN) */}
          <Row className="gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#909090" }}
              >
                Materi Pembelajaran (EN)
              </h3>
            </Col>
            <Col sm={12}>
              <p className="m-0 p-0 fw-medium" style={{ color: "#3A3A3A" }}>
                text
              </p>
            </Col>
          </Row>

          {/*//! col Materi Pembelajaran (IND) */}
          <Row className="gap-1">
            <Col sm={12}>
              <h3
                className="m-0 p-0 fw-semibold fs-6"
                style={{ color: "#909090" }}
              >
                Materi Pembelajaran (IND)
              </h3>
            </Col>
            <Col sm={12}>
              <p className="m-0 p-0 fw-medium" style={{ color: "#3A3A3A" }}>
                text
              </p>
            </Col>
          </Row>
        </section>
      </ModalBody>
    </Modal>
  );
}

export default ModalDetailLessonPlan;
