"use client";

import { Button, Col, Input, Label, Row } from "reactstrap";
import { useCallback, useEffect, useState } from "react";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { EditIcon } from "@/components/icons/edit";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import {
  FormDocumentSchema,
  FormDocumentSchemaType,
} from "@/lib/validations/students/biodata/form-document-schema";
import { FormDescription } from "@/components/ui/form-description";
import { useDocumentStudent } from "@/services/api/students/biodata/document/use-get-document";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { useUpdateDocumentStudent } from "@/services/api/students/biodata/document/use-update-document";
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import { splitFileNameUploaded } from "@/lib/utils/split-filename-upload";
import { useGetFileStorage } from "@/services/api/file-storage/use-get-file-storage";

export const FormDocument = () => {
  const [isEdit, setIsEdit] = useState<boolean>(false);
  const { setModalConfirmationState } = useModalConfirmationContext();
  const { data: document, isLoading: isLoadingDocument } = useDocumentStudent();
  const { mutateAsync: updateDocumentStudent } = useUpdateDocumentStudent();
  const { getFileStorage, loading: loadingFile } = useGetFileStorage();

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    clearErrors,
    setValue,
  } = useForm<FormDocumentSchemaType>({
    resolver: zodResolver(FormDocumentSchema),
    defaultValues: {
      npwp: "",
      bpjs_healthcare: "",
      bpjs_employment: "",
    },
  });

  const handleSetFormValue = useCallback(() => {
    setValue("npwp", document?.data?.npwp || "");
    setValue("bpjs_employment", document?.data?.bpjs_employment || "");
    setValue("bpjs_healthcare", document?.data?.bpjs_healthcare || "");
  }, [document?.data, setValue]);

  const onSubmit = async (payload: FormDocumentSchemaType) => {
    try {
      if (
        !payload.npwp &&
        !payload.bpjs_employment &&
        !payload.bpjs_healthcare
      ) {
        return setIsEdit(false);
      }

      const response = await updateDocumentStudent(payload);

      if (response.error) {
        return setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "failed",
          message: response.message,
        }));
      }

      setIsEdit(false);
      return setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        message: "Data berhasil di-update",
        state: "success",
      }));
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "failed",
        message: error.toString(),
      }));
    }
  };

  useEffect(() => {
    if (document?.data) handleSetFormValue();
  }, [document?.data, handleSetFormValue]);

  if (document?.error) {
    return <h1>{document.message}</h1>;
  }

  return (
    <>
      <div className="d-flex border-2 border-bottom justify-content-between align-items-center">
        <h5 className="card-title py-3 mb-1" style={{ fontWeight: "500" }}>
          Dokumen
        </h5>
        {!isEdit ? (
          <button
            className="bg-transparent rounded px-3 d-flex gap-1 align-items-center justify-content-center text-primary"
            style={{ border: "1px solid #10487A", paddingBlock: "8px" }}
            onClick={() => setIsEdit(true)}
          >
            <EditIcon />
            <span>Edit</span>
          </button>
        ) : null}
      </div>
      <form onSubmit={handleSubmit(onSubmit)} className="my-2">
        <Row>
          {/* Left Column - NPWP and BPJS Kesehatan */}
          <Col md={6} className="pe-md-3">
            {/* NPWP Section */}
            <Row className="align-items-center gap-2 mb-4">
              <Col sm={12}>
                <Label
                  htmlFor="npwp"
                  className="form-label mb-0 fw-medium optional"
                >
                  NPWP
                </Label>
              </Col>
              <Col sm={12}>
                <Controller
                  name="npwp"
                  control={control}
                  render={({ field }) => (
                    <Input
                      className={`form-control form-control-icon ${
                        errors.npwp ? "border border-danger" : ""
                      }`}
                      id="npwp"
                      placeholder="Masukkan NPWP"
                      disabled={!isEdit || isLoadingDocument}
                      {...field}
                      onChange={(e) => {
                        const { stringValue } = handleInputNumberOnly(e);

                        field.onChange(stringValue);
                      }}
                    />
                  )}
                />
                <FormErrorMessage errors={errors.npwp} />
              </Col>

              <Col sm={12} className="mt-2">
                <Label
                  htmlFor="npwp_filepath"
                  className="form-label mb-0 fw-medium optional"
                >
                  File NPWP
                </Label>
              </Col>
              <Col sm={12}>
                <div className="d-flex gap-2">
                  <Controller
                    name="npwp_filepath"
                    control={control}
                    render={({ field: { onChange, value } }) => (
                      <div className="position-relative w-100">
                        <Input
                          type="text"
                          className={`form-control ${
                            errors.npwp_filepath ? "border border-danger" : ""
                          }`}
                          placeholder="Pilih File"
                          readOnly
                          value={
                            value instanceof File
                              ? value.name
                              : splitFileNameUploaded(
                                  document?.data?.npwp_filepath as string
                                ) || ""
                          }
                          disabled={!isEdit || isLoadingDocument}
                        />
                        {isEdit && (
                          <Input
                            type="file"
                            id="npwp_filepath"
                            className={`form-control form-control-icon ${
                              errors.npwp_filepath ? "border border-danger" : ""
                            }`}
                            onChange={(e) => {
                              const file = e.target.files?.[0];
                              if (file) onChange(file);
                            }}
                            hidden
                            accept=".png,.jpg,.jpeg"
                          />
                        )}
                      </div>
                    )}
                  />
                  {isEdit && (
                    <label
                      htmlFor="npwp_filepath"
                      className="btn d-flex align-items-center btn-light  mb-0"
                      style={{ whiteSpace: "nowrap" }}
                    >
                      Upload File
                    </label>
                  )}
                  {!isEdit && (
                    <button
                      className="btn d-flex align-items-center btn-light"
                      style={{ whiteSpace: "nowrap" }}
                      type="button"
                      onClick={
                        document?.data?.npwp_filepath
                          ? async () =>
                              await getFileStorage(
                                document?.data?.npwp_filepath as string
                              )
                          : () => null
                      }
                      disabled={loadingFile}
                    >
                      Lihat File
                    </button>
                  )}
                </div>
                {errors.npwp_filepath ? (
                  <FormErrorMessage errors={errors.npwp_filepath} />
                ) : (
                  <FormDescription message="File dengan format png dan jpg maksimal 2mb" />
                )}
              </Col>
            </Row>

            {/* BPJS Kesehatan Section */}
            <Row className="align-items-center gap-2 mb-4">
              <Col sm={12}>
                <Label
                  htmlFor="bpjs_healthcare"
                  className="form-label mb-0 fw-medium optional"
                >
                  Nomor BPJS Kesehatan
                </Label>
              </Col>
              <Col sm={12}>
                <Controller
                  name="bpjs_healthcare"
                  control={control}
                  render={({ field }) => (
                    <Input
                      className={`form-control form-control-icon ${
                        errors.bpjs_healthcare ? "border border-danger" : ""
                      }`}
                      id="bpjs_healthcare"
                      placeholder="Masukkan Nomor BPJS Kesehatan"
                      disabled={!isEdit || isLoadingDocument}
                      {...field}
                      onChange={(e) => {
                        const { stringValue } = handleInputNumberOnly(e);

                        field.onChange(stringValue);
                      }}
                    />
                  )}
                />
                <FormErrorMessage errors={errors.bpjs_healthcare} />
              </Col>

              <Col sm={12} className="mt-2">
                <Label
                  htmlFor="bpjs_healthcare_filepath"
                  className="form-label mb-0 fw-medium optional"
                >
                  File BPJS Kesehatan
                </Label>
              </Col>
              <Col sm={12}>
                <div className="d-flex gap-2">
                  <Controller
                    name="bpjs_healthcare_filepath"
                    control={control}
                    render={({ field: { onChange, value, ...field } }) => (
                      <div className="position-relative w-100">
                        <Input
                          type="text"
                          className={`form-control ${
                            errors.bpjs_healthcare_filepath
                              ? "border border-danger"
                              : ""
                          }`}
                          placeholder="Pilih File"
                          readOnly
                          value={
                            value instanceof File
                              ? value.name
                              : splitFileNameUploaded(
                                  document?.data
                                    ?.bpjs_healthcare_filepath as string
                                ) || ""
                          }
                          disabled={!isEdit || isLoadingDocument}
                          {...field}
                        />
                        {isEdit && (
                          <Input
                            type="file"
                            id="bpjs_healthcare_filepath"
                            className={`form-control form-control-icon ${
                              errors.bpjs_healthcare_filepath
                                ? "border border-danger"
                                : ""
                            }`}
                            onChange={(e) => {
                              const file = e.target.files?.[0];
                              if (file) onChange(file);
                            }}
                            hidden
                            accept=".png,.jpg,.jpeg"
                          />
                        )}
                      </div>
                    )}
                  />
                  {isEdit && (
                    <label
                      htmlFor="bpjs_healthcare_filepath"
                      className="btn btn-light d-flex align-items-center mb-0"
                      style={{ whiteSpace: "nowrap" }}
                    >
                      Upload File
                    </label>
                  )}
                  {!isEdit && (
                    <button
                      className="btn d-flex align-items-center btn-light"
                      style={{ whiteSpace: "nowrap" }}
                      type="button"
                      onClick={
                        document?.data?.bpjs_healthcare_filepath
                          ? async () =>
                              await getFileStorage(
                                document?.data
                                  ?.bpjs_healthcare_filepath as string
                              )
                          : () => null
                      }
                      disabled={loadingFile}
                    >
                      Lihat File
                    </button>
                  )}
                </div>
                {errors.bpjs_healthcare_filepath ? (
                  <FormErrorMessage errors={errors.bpjs_healthcare_filepath} />
                ) : (
                  <FormDescription message="File dengan format png dan jpg maksimal 2mb" />
                )}
              </Col>
            </Row>
          </Col>

          {/* Right Column - BPJS Ketenagakerjaan */}
          <Col md={6} className="ps-md-3">
            {/* BPJS Ketenagakerjaan Section */}
            <Row className="align-items-center gap-2 mb-4">
              <Col sm={12}>
                <Label
                  htmlFor="bpjs_employment"
                  className="form-label mb-0 fw-medium optional"
                >
                  Nomor BPJS Ketenagakerjaan
                </Label>
              </Col>
              <Col sm={12}>
                <Controller
                  name="bpjs_employment"
                  control={control}
                  render={({ field }) => (
                    <Input
                      className={`form-control form-control-icon ${
                        errors.bpjs_employment ? "border border-danger" : ""
                      }`}
                      id="bpjs_employment"
                      placeholder="Masukkan Nomor BPJS Ketenagakerjaan"
                      disabled={!isEdit || isLoadingDocument}
                      {...field}
                    />
                  )}
                />
                <FormErrorMessage errors={errors.bpjs_employment} />
              </Col>

              <Col sm={12} className="mt-2">
                <Label
                  htmlFor="bpjs_employment_filepath"
                  className="form-label mb-0 fw-medium optional"
                >
                  File BPJS Ketenagakerjaan
                </Label>
              </Col>
              <Col sm={12}>
                <div className="d-flex gap-2">
                  <Controller
                    name="bpjs_employment_filepath"
                    control={control}
                    render={({ field: { onChange, value, ...field } }) => (
                      <div className="position-relative w-100">
                        <Input
                          type="text"
                          className={`form-control ${
                            errors.bpjs_employment_filepath
                              ? "border border-danger"
                              : ""
                          }`}
                          placeholder="Pilih File"
                          readOnly
                          value={
                            value instanceof File
                              ? value.name
                              : splitFileNameUploaded(
                                  document?.data
                                    ?.bpjs_employment_filepath as string
                                ) || ""
                          }
                          disabled={!isEdit || isLoadingDocument}
                          {...field}
                        />
                        {isEdit && (
                          <Input
                            type="file"
                            id="bpjs_employment_filepath"
                            className={`form-control form-control-icon ${
                              errors.bpjs_employment_filepath
                                ? "border border-danger"
                                : ""
                            }`}
                            onChange={(e) => {
                              const file = e.target.files?.[0];
                              if (file) onChange(file);
                            }}
                            hidden
                            accept=".png,.jpg,.jpeg"
                          />
                        )}
                      </div>
                    )}
                  />
                  {isEdit && (
                    <label
                      htmlFor="bpjs_employment_filepath"
                      className="btn d-flex align-items-center btn-light  mb-0"
                      style={{ whiteSpace: "nowrap" }}
                    >
                      Upload File
                    </label>
                  )}
                  {!isEdit && (
                    <button
                      className="btn d-flex align-items-center btn-light"
                      style={{ whiteSpace: "nowrap" }}
                      type="button"
                      onClick={
                        document?.data?.bpjs_employment_filepath
                          ? async () =>
                              await getFileStorage(
                                document?.data
                                  ?.bpjs_employment_filepath as string
                              )
                          : () => null
                      }
                      disabled={loadingFile}
                    >
                      Lihat File
                    </button>
                  )}
                </div>

                {errors.bpjs_employment_filepath ? (
                  <FormErrorMessage errors={errors.bpjs_employment_filepath} />
                ) : (
                  <FormDescription message="File dengan format png dan jpg maksimal 2mb" />
                )}
              </Col>
            </Row>
          </Col>
        </Row>
        {isEdit && (
          <div className="d-flex justify-content-between mt-3 gap-3">
            <button
              onClick={() => {
                setIsEdit(!isEdit);
                clearErrors();
              }}
              className="bg-transparent text-primary rounded px-3"
              type="button"
              style={{ border: "1px solid #10487A" }}
              disabled={isSubmitting}
            >
              <span>Batal</span>
            </button>
            <Button
              color="primary"
              className="d-flex flex-grow-0 justify-content-center align-items-center"
              disabled={isSubmitting}
            >
              <span>Update</span>
            </Button>
          </div>
        )}
      </form>
    </>
  );
};
