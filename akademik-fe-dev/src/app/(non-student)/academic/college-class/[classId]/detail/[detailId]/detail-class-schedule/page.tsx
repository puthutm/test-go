"use client";

import React, { useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { ModalDeleteConfirmation } from "@/components/ui/modal-delete-confirmation";
import { signOut } from "next-auth/react";

// import third party component
import {
  Card,
  CardBody,
  CardHeader,
  Button,
  Row,
  Col,
  Label,
} from "reactstrap";
import { DeleteIcon } from "@/components/icons/delete";
import { EditIcon } from "@/components/icons/edit";

import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { useGetDetailClassScheduleSubDetail } from "@/services/api/academic/lecturer/class-schedule/detail-class/class-schedule/use-get-detail-class-schedule-sub-detail";
import { deleteClassScheduleSubDetail } from "@/services/api/academic/lecturer/class-schedule/detail-class/class-schedule/delete-class-schedule-sub-detail";
import { useGetFileStorage } from "@/services/api/file-storage/use-get-file-storage";
import { formatDate, getHourAndMinute } from "@/lib/utils/format-date";

function PageClassScheduleDetail() {
  const queryclient = useQueryClient();
  const params = useParams();

  const [isLoadingDelete, setIsLoadingDelete] = useState(false);
  // state show modal confirmation
  const { modalConfirmationState, setModalConfirmationState } =
    useModalConfirmationContext();

  const router = useRouter();

  //! get data
  const {
    data: dataDetailClassScheduleSubDetail,
    isLoading: isLoadingDetailClassScheduleSubDetail,
    isError: isErrorDetailClassScheduleSubDetail,
  } = useGetDetailClassScheduleSubDetail(
    params.classId as string,
    params.detailId as string
  );

  //! get file storage
  const { getFileStorage, loading } = useGetFileStorage();

  const handleDelete = async (classId: string, class_schedule_id: string) => {
    setIsLoadingDelete(true);
    try {
      const response = await deleteClassScheduleSubDetail(
        classId,
        class_schedule_id
      );
      if (response.error) {
        setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          id: null,
          state: "failed",
          message: response.message || "Gagal menghapus data",
        }));

        return;
      }

      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "success",
        message: "Data berhasil dihapus",
      }));

      router.replace(
        `/academic/college-class/${params.classId}/detail?tab=schedule-college`
      );
      queryclient.invalidateQueries({
        queryKey: ["get-class-schedule-sub-detail"],
      });
    } catch (err: any) {
      throw new Error(err.message);
    } finally {
      setIsLoadingDelete(false);
    }
  };
  if (dataDetailClassScheduleSubDetail?.status === 401) {
    signOut();
  }
  return (
    <section className="position-relative">
      {/* modal delete confirm */}
      <ModalDeleteConfirmation
        isLoading={isLoadingDelete}
        onDelete={async () => {
          await handleDelete(
            params.classId as string,
            modalConfirmationState.id as string
          );
        }}
      />
      {/*//! DETAIL INFORMATIONAL */}
      <Card className="py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0 pb-2 border-b-2">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1 ">
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 ">
              Detail Jadwal
            </h2>
            <Button
              className="btn d-flex align-items-center gap-2 "
              color="transparent"
              disabled={
                isErrorDetailClassScheduleSubDetail ||
                isLoadingDetailClassScheduleSubDetail ||
                isLoadingDelete
              }
              onClick={() => {
                setModalConfirmationState(() => ({
                  open: true,
                  state: "confirm",
                  message: "hapus data jadwal kuliah",
                  id: dataDetailClassScheduleSubDetail?.data.id,
                }));
              }}
              style={{ color: "#d9534f", border: "1px solid #d9534f" }}
            >
              <DeleteIcon color="#d9534f" />
              Hapus
            </Button>
            <Link
              href={`/academic/college-class/${params.classId}/detail/${params.detailId}/detail-class-schedule/edit`}
              className="btn  d-flex align-items-center gap-2 "
              color="transparent"
              style={{ color: "#10487A", border: "1px solid #10487A" }}
            >
              <EditIcon color="#10487A" />
              Ubah
            </Link>
          </section>
        </CardHeader>
        {/*//! card body */}
        <CardBody className="position-relative px-0">
          <Row className="row-gap-3">
            {/*//! session */}
            <Col sm={6}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    Sesi
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : (
                    <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                      {dataDetailClassScheduleSubDetail?.data.session ?? "-"}
                    </p>
                  )}
                </Col>
              </Row>
            </Col>
            {/*//! metode Pembelajaran */}
            <Col sm={6}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    Metode Pembelajaran
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : (
                    <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                      {"-"}
                    </p>
                  )}
                </Col>
              </Row>
            </Col>
            {/*//! Tanggal Jadwal */}
            <Col sm={6}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    Tanggal Jadwal
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : (
                    <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                      {formatDate(
                        dataDetailClassScheduleSubDetail?.data.date ?? "-"
                      )}
                    </p>
                  )}
                </Col>
              </Row>
            </Col>
            {/*//! Keterangan Ruang Kuliah */}
            <Col sm={6}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    Keterangan Ruang Kuliah
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : (
                    <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                      {"-"}
                    </p>
                  )}
                </Col>
              </Row>
            </Col>
            {/*//! Waktu Mulai */}
            <Col sm={6}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    Waktu Mulai
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : (
                    <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                      {getHourAndMinute(
                        dataDetailClassScheduleSubDetail?.data.start_time ?? "-"
                      )}
                    </p>
                  )}
                </Col>
              </Row>
            </Col>
            {/*//! URL Kuliah Online */}
            <Col sm={6}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    URL Kuliah Online
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : (
                    <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                      {"-"}
                    </p>
                  )}
                </Col>
              </Row>
            </Col>
            {/*//! Waktu Selesai */}
            <Col sm={6}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    Waktu Selesai
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : (
                    <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                      {getHourAndMinute(
                        dataDetailClassScheduleSubDetail?.data.end_time ?? "-"
                      )}
                    </p>
                  )}
                </Col>
              </Row>
            </Col>
            {/*//! SKS */}
            <Col sm={6}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    SKS
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : (
                    <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                      {"-"}
                    </p>
                  )}
                </Col>
              </Row>
            </Col>
            {/*//! Jenis Pertemuan */}
            <Col sm={6}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    Jenis Pertemuan
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : (
                    <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                      {dataDetailClassScheduleSubDetail?.data.type_of_meeting ??
                        "-"}
                    </p>
                  )}
                </Col>
              </Row>
            </Col>
            {/*//! Status */}
            <Col sm={6}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    Status
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : (
                    <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                      {"-"}
                    </p>
                  )}
                </Col>
              </Row>
            </Col>
          </Row>
        </CardBody>
      </Card>

      {/*//! INFORMATIONAL MATERI */}
      <Card className="mt-4 py-3 px-4 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0 pb-2 border-b-2">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1 ">
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 ">Materi</h2>
          </section>
        </CardHeader>

        {/*//! card body */}
        <CardBody className="position-relative px-0">
          <Row className="row-gap-3">
            {/*//! Lampiran Materi */}
            <Col sm={12}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    Lampiran Materi
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : dataDetailClassScheduleSubDetail?.data
                      .material_attachment_file_path != null ? (
                    <Button
                      type="button"
                      disabled={loading}
                      className="btn create-btn d-flex align-items-center p-0 text-light"
                      style={{ background: "#44A7FF" }}
                      // onClick={handleShow}
                      onClick={() => {
                        getFileStorage(
                          dataDetailClassScheduleSubDetail?.data
                            ?.material_attachment_file_path ?? ""
                        );
                      }}
                      color="transparent"
                      data-bs-target="#api-key-modal"
                    >
                      <Label
                        htmlFor="material_attachment_file_path"
                        className="py-1 px-3 m-0 pointer fs-6"
                        style={{ cursor: "pointer" }}
                      >
                        Lampiran Materi.pdf
                      </Label>
                    </Button>
                  ) : (
                    "-"
                  )}
                </Col>
              </Row>
            </Col>

            {/*//! Dokumen Presensi */}
            <Col sm={12}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    Dokumen Presensi
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : dataDetailClassScheduleSubDetail?.data
                      .attendance_document_file_path != null ? (
                    <Button
                      type="button"
                      disabled={loading}
                      className="btn create-btn d-flex align-items-center p-0 text-light"
                      style={{ background: "#44A7FF" }}
                      onClick={() => {
                        getFileStorage(
                          dataDetailClassScheduleSubDetail?.data
                            ?.attendance_document_file_path ?? ""
                        );
                      }}
                      color="transparent"
                      data-bs-target="#api-key-modal"
                    >
                      <Label
                        htmlFor="material_attachment_file_path"
                        className="py-1 px-3 m-0 pointer fs-6"
                        style={{ cursor: "pointer" }}
                      >
                        Dokumen Presensi.pdf
                      </Label>
                    </Button>
                  ) : (
                    "-"
                  )}
                </Col>
              </Row>
            </Col>

            {/*//! Dokument Jurnal */}
            <Col sm={12}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    Dokumen Jurnal
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : dataDetailClassScheduleSubDetail?.data
                      .journal_document_file_path != null ? (
                    <Button
                      type="button"
                      disabled={loading}
                      className="btn create-btn d-flex align-items-center p-0 text-light"
                      style={{ background: "#44A7FF" }}
                      onClick={() => {
                        getFileStorage(
                          dataDetailClassScheduleSubDetail?.data
                            ?.journal_document_file_path ?? ""
                        );
                      }}
                      color="transparent"
                      data-bs-target="#api-key-modal"
                    >
                      <Label
                        htmlFor="material_attachment_file_path"
                        className="py-1 px-3 m-0 pointer fs-6"
                        style={{ cursor: "pointer" }}
                      >
                        Dokumen Jurnal.pdf
                      </Label>
                    </Button>
                  ) : (
                    "-"
                  )}
                </Col>
              </Row>
            </Col>

            {/*//! Materi Pembelajaran */}
            <Col sm={12}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    Materi Pembelajaran
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : (
                    <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                      {dataDetailClassScheduleSubDetail?.data.material_plan ??
                        "-"}
                    </p>
                  )}
                </Col>
              </Row>
            </Col>

            {/*//!  Realisasi Materi */}
            <Col sm={12}>
              <Row className="row-gap-1">
                <Col sm={12}>
                  <h3
                    className="m-0 p-0 fw-semibold fs-6"
                    style={{ color: "#3A3A3A" }}
                  >
                    Realisasi Materi
                  </h3>
                </Col>
                <Col sm={12}>
                  {isLoadingDetailClassScheduleSubDetail ? (
                    <span className="placeholder w-100 " />
                  ) : (
                    <p className="m-0 p-0 " style={{ color: "#3A3A3A" }}>
                      {dataDetailClassScheduleSubDetail?.data
                        .material_realization ?? "-"}
                    </p>
                  )}
                </Col>
              </Row>
            </Col>
          </Row>
        </CardBody>
      </Card>

      {/*//! button back */}
      <section className="position-relative mt-4">
        <Link
          href={`/academic/college-class/${params.classId}/detail?tab=schedule-college`}
          className="p-2 rounded-2 "
          style={{
            color: "#10487A",
            fontSize: "13px",
            border: "1px solid #10487A",
          }}
        >
          Kembali
        </Link>
      </section>
    </section>
  );
}

export default PageClassScheduleDetail;
