"use client";

import { Button, Col, Input, Label, Row } from "reactstrap";
import { useCallback, useEffect, useState } from "react";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { FormDescription } from "@/components/ui/form-description";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import { EditIcon } from "@/components/icons/edit";
import {
  FormOriginalEducationSchema,
  FormOriginalEducationSchemaType,
} from "@/lib/validations/students/biodata/form-original-education-schema";
import { FileUploadIcon } from "@/components/icons/file-upload";
import { splitFileNameUploaded } from "@/lib/utils/split-filename-upload";
import { useOriginalEducationStudent } from "@/services/api/students/biodata/original-education/use-get-original-education";
import { useUpdateOriginalEducationStudent } from "@/services/api/students/biodata/original-education/use-update-original-education";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { useGetFileStorage } from "@/services/api/file-storage/use-get-file-storage";

export const FormOriginalEducation = () => {
  const [isEdit, setIsEdit] = useState<boolean>(false);
  const { setModalConfirmationState } = useModalConfirmationContext();

  const {
    data: dataOriginalEducationStudent,
    isLoading: isLoadingOriginalEducationStudent,
  } = useOriginalEducationStudent();
  const { mutateAsync: updateOriginalEducationStudent } =
    useUpdateOriginalEducationStudent();
  const { getFileStorage, loading: loadingFile } = useGetFileStorage();

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    clearErrors,
    setValue,
  } = useForm<FormOriginalEducationSchemaType>({
    resolver: zodResolver(FormOriginalEducationSchema),
    defaultValues: {
      institution_name: "",
      school_major: "",
      nisn: "",
      national_exam_score: "",
      certificate_number: "",
      certificate_filepath: undefined,
      transcripts_filepath: undefined,
    },
  });

  const handleSetFormValue = useCallback(() => {
    setValue(
      "institution_name",
      dataOriginalEducationStudent?.data.institution_name || ""
    );
    setValue(
      "school_major",
      dataOriginalEducationStudent?.data.school_major || ""
    );
    setValue("nisn", dataOriginalEducationStudent?.data.nisn || "");
    setValue(
      "certificate_number",
      dataOriginalEducationStudent?.data.certificate_number || ""
    );
    setValue(
      "national_exam_score",
      dataOriginalEducationStudent?.data.national_exam_score?.toString() || ""
    );
  }, [dataOriginalEducationStudent?.data, setValue]);

  const onSubmit = async (payload: FormOriginalEducationSchemaType) => {
    try {
      const response = await updateOriginalEducationStudent(payload);

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
    if (dataOriginalEducationStudent?.data) handleSetFormValue();
  }, [dataOriginalEducationStudent?.data, handleSetFormValue]);

  if (dataOriginalEducationStudent?.error) {
    return <h1>{dataOriginalEducationStudent.message}</h1>;
  }

  return (
    <>
      <div className="d-flex border-2 border-bottom justify-content-between align-items-center">
        <h5 className="card-title py-3 mb-1" style={{ fontWeight: "500" }}>
          Pendidikan Asal
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
        <Row className="gap-1 gap-lg-0">
          {/* left section */}
          <Col md={12} lg={6}>
            <Row className="gap-2">
              {/* institution_name */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="institution_name"
                      className="form-label mb-0 fw-medium"
                    >
                      Sekolah / Perguruan Tinggi
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="institution_name"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.institution_name
                              ? "border border-danger"
                              : ""
                          }`}
                          id="institution_name"
                          placeholder="Masukkan Sekolah / Perguruan Tinggi"
                          disabled={
                            !isEdit || isLoadingOriginalEducationStudent
                          }
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.institution_name} />
                  </Col>
                </Row>
              </Col>
              {/* school_major */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="school_major"
                      className="form-label mb-0 fw-medium"
                    >
                      Jurusan / Program Studi
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="school_major"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.school_major ? "border border-danger" : ""
                          }`}
                          id="school_major"
                          placeholder="Masukkan Jurusan / Program Studi"
                          disabled={
                            !isEdit || isLoadingOriginalEducationStudent
                          }
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.school_major} />
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
                    <Label htmlFor="nisn" className="form-label mb-0 fw-medium">
                      NIM / NISN
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
                          id="nisn"
                          placeholder="Masukkan NIM / NISN"
                          disabled={
                            !isEdit || isLoadingOriginalEducationStudent
                          }
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);

                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.nisn} />
                  </Col>
                </Row>
              </Col>
              {/* national_exam_score */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="national_exam_score"
                      className="form-label mb-0 fw-medium"
                    >
                      IPK / Nilai Rata-Rata
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="national_exam_score"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.national_exam_score
                              ? "border border-danger"
                              : ""
                          }`}
                          id="national_exam_score"
                          placeholder="Masukkan IPK / Nilai Rata-Rata"
                          disabled={
                            !isEdit || isLoadingOriginalEducationStudent
                          }
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);

                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.national_exam_score} />
                  </Col>
                </Row>
              </Col>
            </Row>
          </Col>
          {/* right section */}
          <Col md={12} lg={6}>
            <Row className="gap-2">
              {/* certificate_number */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="certificate_number"
                      className="form-label mb-0 fw-medium"
                    >
                      No. Ijazah
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="certificate_number"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.certificate_number
                              ? "border border-danger"
                              : ""
                          }`}
                          id="certificate_number"
                          placeholder="Masukkan No. Ijazah"
                          disabled={
                            !isEdit || isLoadingOriginalEducationStudent
                          }
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);

                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.certificate_number} />
                  </Col>
                </Row>
              </Col>
              {/* certificate file */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="certificate_filepath"
                      className="form-label mb-0 fw-medium"
                    >
                      Ijazah
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <div className="d-flex gap-2">
                      <Controller
                        name="certificate_filepath"
                        control={control}
                        render={({ field: { onChange, value, ...field } }) => (
                          <div className="position-relative w-100">
                            <Input
                              type="text"
                              className={`form-control form-control-icon ${
                                errors.certificate_filepath
                                  ? "border border-danger"
                                  : ""
                              }`}
                              value={
                                value instanceof File
                                  ? value.name
                                  : splitFileNameUploaded(
                                      dataOriginalEducationStudent?.data
                                        ?.certificate_filepath as string
                                    ) || ""
                              }
                              disabled={
                                !isEdit || isLoadingOriginalEducationStudent
                              }
                              placeholder="Pilih File"
                              {...field}
                            />
                            {isEdit && (
                              <Input
                                type="file"
                                id="certificate_filepath"
                                accept=".pdf"
                                onChange={(e) => {
                                  const file = e.target.files?.[0];
                                  if (file) onChange(file);
                                }}
                                hidden
                              />
                            )}
                            {errors.certificate_filepath ? (
                              <FormErrorMessage
                                errors={errors.certificate_filepath}
                              />
                            ) : (
                              <FormDescription message="File dengan format .pdf maksimal 2mb" />
                            )}
                          </div>
                        )}
                      />
                      {isEdit ? (
                        <label
                          htmlFor="certificate_filepath"
                          className="btn d-flex align-items-center btn-light mb-0"
                          style={{ whiteSpace: "nowrap" }}
                        >
                          <FileUploadIcon className="me-1" /> Upload
                        </label>
                      ) : (
                        <button
                          className="btn d-flex align-items-center btn-light"
                          style={{ whiteSpace: "nowrap" }}
                          type="button"
                          onClick={
                            dataOriginalEducationStudent?.data
                              ?.certificate_filepath
                              ? async () =>
                                  await getFileStorage(
                                    dataOriginalEducationStudent?.data
                                      ?.certificate_filepath as string
                                  )
                              : () => null
                          }
                          disabled={loadingFile}
                        >
                          Lihat File
                        </button>
                      )}
                    </div>
                  </Col>
                </Row>
              </Col>
              {/* transcript */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="transcripts_filepath"
                      className="form-label mb-0 fw-medium"
                    >
                      Transkrip Nilai
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <div className="d-flex gap-2">
                      <Controller
                        name="transcripts_filepath"
                        control={control}
                        render={({ field: { onChange, value, ...field } }) => (
                          <div className="position-relative w-100">
                            <Input
                              type="text"
                              className={`form-control form-control-icon ${
                                errors.transcripts_filepath
                                  ? "border border-danger"
                                  : ""
                              }`}
                              value={
                                value instanceof File
                                  ? value.name
                                  : splitFileNameUploaded(
                                      dataOriginalEducationStudent?.data
                                        ?.transcripts_filepath as string
                                    ) || ""
                              }
                              disabled={
                                !isEdit || isLoadingOriginalEducationStudent
                              }
                              placeholder="Pilih File"
                              {...field}
                            />
                            {isEdit && (
                              <Input
                                type="file"
                                id="transcripts_filepath"
                                accept=".pdf"
                                onChange={(e) => {
                                  const file = e.target.files?.[0];
                                  if (file) onChange(file);
                                }}
                                hidden
                              />
                            )}
                            {errors.transcripts_filepath ? (
                              <FormErrorMessage
                                errors={errors.transcripts_filepath}
                              />
                            ) : (
                              <FormDescription message="File dengan format .pdf maksimal 2mb" />
                            )}
                          </div>
                        )}
                      />
                      {isEdit ? (
                        <label
                          htmlFor="transcripts_filepath"
                          className="btn d-flex align-items-center btn-light mb-0"
                          style={{ whiteSpace: "nowrap" }}
                        >
                          <FileUploadIcon className="me-1" /> Upload
                        </label>
                      ) : (
                        <button
                          className="btn d-flex align-items-center btn-light"
                          style={{ whiteSpace: "nowrap" }}
                          type="button"
                          onClick={
                            dataOriginalEducationStudent?.data
                              ?.transcripts_filepath
                              ? async () =>
                                  await getFileStorage(
                                    dataOriginalEducationStudent?.data
                                      ?.transcripts_filepath as string
                                  )
                              : () => null
                          }
                          disabled={loadingFile}
                        >
                          Lihat File
                        </button>
                      )}
                    </div>
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
