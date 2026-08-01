"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useCallback, useEffect, useState } from "react";
import { Controller, useForm } from "react-hook-form";
import { Button, Col, Input, Label, Row } from "reactstrap";

import { EditIcon } from "@/components/icons/edit";
import { FormDescription } from "@/components/ui/form-description";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import {
  FormCompletenessSchema,
  FormCompletenessSchemaType,
} from "@/lib/validations/students/biodata/form-completness-schema";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { useUpdateCompleteness } from "@/services/api/students/biodata/completeness/use-update-completeness";
import { useCompletenessStudent } from "@/services/api/students/biodata/completeness/use-get-completeness";
import { useGetFileStorage } from "@/services/api/file-storage/use-get-file-storage";
import { splitFileNameUploaded } from "@/lib/utils/split-filename-upload";

export const FormCompleteness = () => {
  const [isEdit, setIsEdit] = useState<boolean>(false);

  const { setModalConfirmationState } = useModalConfirmationContext();
  const { mutateAsync: updateCompletenessStudent } = useUpdateCompleteness();
  const { data: dataCompleteness } = useCompletenessStudent();
  const { getFileStorage, loading: loadingFile } = useGetFileStorage();

  const {
    formState: { errors, isSubmitting },
    control,
    clearErrors,
    setValue,
    handleSubmit,
  } = useForm<FormCompletenessSchemaType>({
    resolver: zodResolver(FormCompletenessSchema),
    defaultValues: {
      nis: "",
      nisn: "",
      no_passport: "",
      google_scholar: "",
      sinta_id: "",
      scopus_id: "",
    },
  });

  const handleSetFormValue = useCallback(() => {
    setValue("google_scholar", dataCompleteness?.data.google_scholar || "");
    setValue("no_passport", dataCompleteness?.data.no_passport || "");
    setValue("scopus_id", dataCompleteness?.data.scopus_id || "");
    setValue("sinta_id", dataCompleteness?.data.sinta_id || "");
  }, [dataCompleteness?.data, setValue]);

  const onSubmit = async (payload: FormCompletenessSchemaType) => {
    try {
      const response = await updateCompletenessStudent(payload);

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
    if (dataCompleteness?.data) handleSetFormValue();
  }, [dataCompleteness?.data, handleSetFormValue]);

  if (dataCompleteness?.error) {
    return <h1>{dataCompleteness?.message}</h1>;
  }
  return (
    <>
      <div className="d-flex border-2 border-bottom justify-content-between align-items-center">
        <h5 className="card-title py-3 mb-1" style={{ fontWeight: "500" }}>
          Kelengkapan
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
      <form
        className="my-2"
        onSubmit={handleSubmit(onSubmit)}
        autoComplete="off"
      >
        <Row className="gap-1 gap-lg-0">
          {/* bagian kiri */}
          <Col md={12} lg={6}>
            <Row className="gap-2">
              {/* NIS */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label htmlFor="nim" className="form-label mb-0 fw-medium">
                      NIS
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="nis"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.nis ? "border border-danger" : ""
                          }`}
                          id="nim"
                          placeholder="Text"
                          disabled
                          {...field}
                          onChange={(e) => {
                            const { numberValue } = handleInputNumberOnly(e);

                            field.onChange(numberValue);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.nis} />
                  </Col>
                </Row>
              </Col>
              {/* nisn */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label htmlFor="name" className="form-label mb-0 fw-medium">
                      NISN
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="nisn"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.nisn ? "border border-danger" : ""
                          }`}
                          id="name"
                          placeholder="Text"
                          disabled
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.nisn} />
                  </Col>
                </Row>
              </Col>
              {/* no paspor */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="back_title"
                      className="form-label mb-0 fw-medium optional"
                    >
                      No Paspor
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="no_passport"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.no_passport ? "border border-danger" : ""
                          }`}
                          id="back_title"
                          placeholder="Text"
                          disabled={!isEdit}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.no_passport} />
                  </Col>
                </Row>
              </Col>
              {/* google scholar */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="back_title"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Google Scholar
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="google_scholar"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.google_scholar ? "border border-danger" : ""
                          }`}
                          id="back_title"
                          placeholder="Text"
                          disabled={!isEdit}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.google_scholar} />
                  </Col>
                </Row>
              </Col>
            </Row>
          </Col>
          {/* right section */}
          <Col md={16} lg={6}>
            <Row className="gap-2">
              {/* id sinta */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="back_title"
                      className="form-label mb-0 fw-medium optional"
                    >
                      ID Sinta
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="sinta_id"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.sinta_id ? "border border-danger" : ""
                          }`}
                          id="back_title"
                          placeholder="Text"
                          disabled={!isEdit}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.sinta_id} />
                  </Col>
                </Row>
              </Col>
              {/* id scopus */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="back_title"
                      className="form-label mb-0 fw-medium optional"
                    >
                      ID Scopus
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="scopus_id"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.scopus_id ? "border border-danger" : ""
                          }`}
                          id="back_title"
                          placeholder="Text"
                          disabled={!isEdit}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.scopus_id} />
                  </Col>
                </Row>
              </Col>
              {/* file tanda tangan */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="signature_path_file"
                      className="form-label mb-0 fw-medium optional"
                    >
                      File Tanda Tangan
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <div className="d-flex gap-2">
                      <Controller
                        name="signature_path_file"
                        control={control}
                        render={({ field: { onChange, value, ...field } }) => (
                          <div className="position-relative w-100">
                            <Input
                              type="text"
                              className={`form-control ${
                                errors.signature_path_file
                                  ? "border border-danger"
                                  : ""
                              }`}
                              placeholder="Pilih File"
                              readOnly
                              value={
                                value instanceof File
                                  ? value.name
                                  : splitFileNameUploaded(
                                      dataCompleteness?.data
                                        ?.signature_path_file as string
                                    ) || ""
                              }
                              disabled={!isEdit}
                              {...field}
                            />
                            {isEdit && (
                              <Input
                                type="file"
                                id="signature_path_file"
                                className={`form-control form-control-icon ${
                                  errors.signature_path_file
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
                          htmlFor="signature_path_file"
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
                            dataCompleteness?.data?.signature_path_file
                              ? async () =>
                                  await getFileStorage(
                                    dataCompleteness?.data
                                      ?.signature_path_file as string
                                  )
                              : () => null
                          }
                          disabled={loadingFile}
                        >
                          Lihat File
                        </button>
                      )}
                    </div>
                    {errors.signature_path_file ? (
                      <FormErrorMessage errors={errors.signature_path_file} />
                    ) : (
                      <FormDescription message="File dengan format png dan jpg maksimal 2mb" />
                    )}
                  </Col>
                </Row>
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
