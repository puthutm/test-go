"use client";
import React, { useEffect } from "react";
import { useQueryClient } from "@tanstack/react-query";
import Link from "next/link";
import {
  Card,
  CardBody,
  CardHeader,
  Button,
  Row,
  Input,
  Col,
  Label,
  Spinner,
} from "reactstrap";
import { UploadIcon } from "@/components/icons/upload";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { FormDescription } from "@/components/ui/form-description";
import { signOut } from "next-auth/react";
import { editClassScheduleSubDetail } from "@/services/api/academic/lecturer/class-schedule/detail-class/class-schedule/update-class-schedule-sub-detail";
import {
  FormSubDetailClassSchedule,
  FormSubDetailClassScheduleType,
} from "@/lib/validations/academic/settings/college-class/form-sub-detail-class-schedule";
import { useForm, SubmitHandler, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useGetFileStorage } from "@/services/api/file-storage/use-get-file-storage";
import { splitFileNameUploaded } from "@/lib/utils/split-filename-upload";
import { ModalSuccessConfirmation } from "@/components/ui/modal-success-confirmation";

import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

function ClientFormEditSubClassSchedule({
  dataDetailClassSchedule,
}: {
  dataDetailClassSchedule: ApiResponse<IClassScheduleSubDetail>; // Replace 'any' with the actual type of dataDetailClassSchedule
}) {
  const queryclient = useQueryClient();
  // use form
  const {
    handleSubmit,
    setValue,
    reset,
    formState: { errors, isSubmitting },
    control,
  } = useForm<FormSubDetailClassScheduleType>({
    resolver: zodResolver(FormSubDetailClassSchedule),
    mode: "onChange",
    defaultValues: {
      material_plan: "",
      material_realization: "",
    },
  });
  // state show modal confirmation
  const { setModalConfirmationState } = useModalConfirmationContext();
  //! get file storage
  const { getFileStorage, loading } = useGetFileStorage();

  useEffect(() => {
    if (dataDetailClassSchedule.data) {
      // Set default values for the form fields based on dataDetailClassSchedule
      setValue(
        "material_attachment_file",
        dataDetailClassSchedule?.data?.material_attachment_file_path ?? ""
      );
      setValue(
        "attendance_document_file",
        dataDetailClassSchedule?.data?.attendance_document_file_path ?? ""
      );
      setValue(
        "journal_document_file",
        dataDetailClassSchedule?.data?.journal_document_file_path ?? ""
      );
      setValue(
        "material_plan",
        dataDetailClassSchedule.data.material_plan || ""
      );
      setValue(
        "material_realization",
        dataDetailClassSchedule.data.material_realization || ""
      );
    }
  }, [dataDetailClassSchedule]);

  //! event submit
  const onSubmit: SubmitHandler<FormSubDetailClassScheduleType> = async (
    data,
    event
  ) => {
    event?.preventDefault();
    try {
      const formData = new FormData();
      formData.append("material_plan", data.material_plan);
      formData.append("material_realization", data.material_realization);
      if (data.material_attachment_file instanceof File) {
        formData.append(
          "material_attachment_file",
          data.material_attachment_file
        );
      }
      if (data.attendance_document_file instanceof File) {
        formData.append(
          "attendance_document_file",
          data.attendance_document_file
        );
      }
      if (data.journal_document_file instanceof File) {
        formData.append("journal_document_file", data.journal_document_file);
      }

      const response = await editClassScheduleSubDetail(
        dataDetailClassSchedule.data.class_id,
        dataDetailClassSchedule.data.id,
        formData
      );

      if (response.error) {
        setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          id: null,
          state: "failed",
          message: response.message || "Gagal mengubah data",
        }));
        return;
      }

      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        id: null,
        state: "success",
        message: response.message || "Berhasil mengubah data",
      }));
      queryclient.invalidateQueries({
        queryKey: ["get-detail-class-schedule-sub-detail"],
      });
    } catch (error: any) {
      throw new Error(error.message || "gagal mengupdate data");
    } finally {
      reset();
    }
    // Handle form submission logic here
  };
  if (dataDetailClassSchedule.status === 401) {
    signOut();
  }
  // You can use the dataDetailClassSchedule to render your component
  return (
    <section className="position-relative">
      <ModalSuccessConfirmation />

      {/*//! card ipload materi */}
      <Card className=" p-3 rounded-3 bg-white">
        {/*//! card header */}
        <CardHeader className="p-0 m-0 pb-2 border-b-2">
          {/*//! TITLE */}
          <section className="d-flex align-items-center gap-1 ">
            <h2 className="m-0 p-0 fs-5 fw-semibold flex-grow-1 ">Materi</h2>
          </section>
        </CardHeader>

        {/*//! card body */}
        <CardBody className="position-relative px-0">
          <form onSubmit={handleSubmit(onSubmit)}>
            <Row className="row-gap-3">
              {/*//! col lampiran materi */}
              <Col sm={12} style={{ position: "unset" }}>
                <label
                  className="form-label  mb-1 "
                  style={{ color: "#3A3A3A" }}
                >
                  Lampiran Materi
                </label>
                <Controller
                  name="material_attachment_file"
                  control={control}
                  render={({ field }) => (
                    <div className="d-flex gap-2">
                      <label
                        htmlFor="material_attachment_file"
                        className={`form-control form-control-icon mb-0 ${
                          errors.material_attachment_file
                            ? "border border-danger"
                            : ""
                        }`}
                      >
                        {field.value instanceof File
                          ? field.value.name
                          : splitFileNameUploaded(
                              dataDetailClassSchedule?.data
                                ?.material_attachment_file_path ?? ""
                            ) || "Belum ada berkas yang dipilih"}
                      </label>
                      <Input
                        type="file"
                        id="material_attachment_file"
                        accept="application/pdf"
                        onChange={(e) => {
                          const file = e.target.files?.[0];
                          field.onChange(file);
                        }}
                        hidden
                      />
                      {dataDetailClassSchedule?.data
                        ?.material_attachment_file_path ? (
                        <Button
                          color="light"
                          className="flex-shrink-0 rounded"
                          type="button"
                          onClick={() =>
                            getFileStorage(
                              dataDetailClassSchedule?.data
                                ?.material_attachment_file_path ?? ""
                            )
                          }
                          disabled={loading}
                        >
                          {loading ? "Mengunduh" : "Lihat File"}
                        </Button>
                      ) : (
                        <label
                          htmlFor="material_attachment_file"
                          className={`btn d-flex align-items-center btn-light mb-0 ${
                            isSubmitting || loading ? "pe-none" : ""
                          }`}
                          style={{ whiteSpace: "nowrap" }}
                        >
                          <UploadIcon /> Upload File
                        </label>
                      )}
                    </div>
                  )}
                />
                {errors.material_attachment_file ? (
                  <FormErrorMessage errors={errors.material_attachment_file} />
                ) : (
                  <FormDescription message="File dengan format .pdf maksimal 10mb" />
                )}
              </Col>

              {/*//! col dokumen presensi */}
              <Col sm={12} style={{ position: "unset" }}>
                <label
                  className="form-label  mb-1 "
                  style={{ color: "#3A3A3A" }}
                >
                  Dokumen Presensi
                </label>
                <Controller
                  name="attendance_document_file"
                  control={control}
                  render={({ field }) => (
                    <div className="d-flex gap-2">
                      <label
                        htmlFor="attendance_document_file"
                        className={`form-control form-control-icon mb-0 ${
                          errors.attendance_document_file
                            ? "border border-danger"
                            : ""
                        }`}
                      >
                        {field.value instanceof File
                          ? field.value.name
                          : splitFileNameUploaded(
                              dataDetailClassSchedule?.data
                                ?.attendance_document_file_path ?? ""
                            ) || "Belum ada berkas yang dipilih"}
                      </label>
                      <Input
                        type="file"
                        id="attendance_document_file"
                        accept="application/pdf"
                        onChange={(e) => {
                          const file = e.target.files?.[0];
                          field.onChange(file);
                        }}
                        hidden
                      />
                      {dataDetailClassSchedule?.data
                        ?.attendance_document_file_path ? (
                        <Button
                          color="light"
                          className="flex-shrink-0 rounded"
                          type="button"
                          onClick={() =>
                            getFileStorage(
                              dataDetailClassSchedule?.data
                                ?.attendance_document_file_path ?? ""
                            )
                          }
                          disabled={loading}
                        >
                          {loading ? "Mengunduh" : "Lihat File"}
                        </Button>
                      ) : (
                        <label
                          htmlFor="attendance_document_file"
                          className={`btn d-flex align-items-center btn-light mb-0 ${
                            isSubmitting || loading ? "pe-none" : ""
                          }`}
                          style={{ whiteSpace: "nowrap" }}
                        >
                          <UploadIcon /> Upload File
                        </label>
                      )}
                    </div>
                  )}
                />
                {errors.attendance_document_file ? (
                  <FormErrorMessage errors={errors.attendance_document_file} />
                ) : (
                  <FormDescription message="File dengan format .pdf maksimal 10mb" />
                )}
              </Col>

              {/*//! col dokument Jurnal */}
              <Col sm={12} style={{ position: "unset" }}>
                <label
                  className="form-label  mb-1 "
                  style={{ color: "#3A3A3A" }}
                >
                  Dokumen Jurnal
                </label>
                <Controller
                  name="journal_document_file"
                  control={control}
                  render={({ field }) => (
                    <div className="d-flex gap-2">
                      <label
                        htmlFor="journal_document_file"
                        className={`form-control form-control-icon mb-0 ${
                          errors.journal_document_file
                            ? "border border-danger"
                            : ""
                        }`}
                      >
                        {field.value instanceof File
                          ? field.value.name
                          : splitFileNameUploaded(
                              dataDetailClassSchedule?.data
                                ?.journal_document_file_path ?? ""
                            ) || "Belum ada berkas yang dipilih"}
                      </label>
                      <Input
                        type="file"
                        id="journal_document_file"
                        accept="application/pdf"
                        onChange={(e) => {
                          const file = e.target.files?.[0];
                          field.onChange(file);
                        }}
                        hidden
                      />
                      {dataDetailClassSchedule?.data
                        ?.journal_document_file_path ? (
                        <Button
                          color="light"
                          className="flex-shrink-0 rounded"
                          type="button"
                          onClick={() =>
                            getFileStorage(
                              dataDetailClassSchedule?.data
                                ?.journal_document_file_path ?? ""
                            )
                          }
                          disabled={loading}
                        >
                          {loading ? "Mengunduh" : "Lihat File"}
                        </Button>
                      ) : (
                        <label
                          htmlFor="journal_document_file"
                          className={`btn d-flex align-items-center btn-light mb-0 ${
                            isSubmitting || loading ? "pe-none" : ""
                          }`}
                          style={{ whiteSpace: "nowrap" }}
                        >
                          <UploadIcon /> Upload File
                        </label>
                      )}
                    </div>
                  )}
                />
                {errors.journal_document_file ? (
                  <FormErrorMessage errors={errors.journal_document_file} />
                ) : (
                  <FormDescription message="File dengan format .pdf maksimal 10mb" />
                )}
              </Col>

              {/*//! col materi plan */}
              <Col sm={12}>
                <Label htmlFor="material_plan" className="form-label mb-1 ">
                  Rencana Materi
                </Label>
                <Controller
                  control={control}
                  name="material_plan"
                  render={({ field }) => (
                    <Input
                      {...field}
                      className={`${
                        errors.material_plan && "border border-danger"
                      } form-control`}
                      id="material_plan"
                      placeholder="Masukkan rencana materi"
                      type="textarea"
                      rows={3}
                    />
                  )}
                />
                <FormErrorMessage errors={errors.material_plan} />
              </Col>

              {/*//! col materi realization */}
              <Col sm={12}>
                <Label
                  htmlFor="material_realization"
                  className="form-label mb-1 "
                >
                  Realisasi Materi
                </Label>
                <Controller
                  control={control}
                  name="material_realization"
                  render={({ field }) => (
                    <Input
                      {...field}
                      className={`${
                        errors.material_realization && "border border-danger"
                      } form-control`}
                      id="material_realization"
                      placeholder="Masukkan realisasi materi"
                      type="textarea"
                      rows={3}
                    />
                  )}
                />
                <FormErrorMessage errors={errors.material_realization} />
              </Col>
            </Row>
          </form>
        </CardBody>
      </Card>

      {/*//! button back */}
      <section className="position-relative mt-4 d-flex justify-content-between gap-2">
        <Link
          href={`/academic/college-class/${dataDetailClassSchedule.data.class_id}/detail/${dataDetailClassSchedule.data.id}/detail-class-schedule`}
          className="p-2 rounded-2 "
          style={{
            color: "#10487A",
            fontSize: "13px",
            border: "1px solid #10487A",
          }}
        >
          Kembali
        </Link>

        <Button
          onClick={handleSubmit(onSubmit)}
          className="btn  d-flex align-items-center gap-2 "
          color="primary"
          disabled={isSubmitting || loading}
        >
          {isSubmitting ? <Spinner size={"sm"} /> : "Update"}
        </Button>
      </section>
    </section>
  );
}

export default ClientFormEditSubClassSchedule;
