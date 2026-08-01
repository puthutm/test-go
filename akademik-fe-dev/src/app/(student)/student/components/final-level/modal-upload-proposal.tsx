"use client";

import { Button, Col, Input, Label, Modal, ModalBody, Row } from "reactstrap";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect } from "react";
import { useQueryClient } from "@tanstack/react-query";

import {
  FormUploadFinalProjectProposalSchema,
  FormUploadPropFinalProjectosalSchemaType,
} from "@/lib/validations/students/final-project-proposal/form-upload-proposal-schema";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { FormDescription } from "@/components/ui/form-description";
import { FileUploadIcon } from "@/components/icons/file-upload";
import { useModalContext } from "@/lib/hooks/use-modal";
import { createFinalProjectProposal } from "@/services/api/students/final-project-proposal/create-final-project-proposal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { splitFileNameUploaded } from "@/lib/utils/split-filename-upload";
import { useGetFinalProjectByProposalId } from "@/services/api/students/final-project-proposal/use-get-final-project-by-proposal-id";
import { useGetFileStorage } from "@/services/api/file-storage/use-get-file-storage";

export const ModalUploadProposal = () => {
  const { modalState, setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();
  const { getFileStorage, loading: loadingFile } = useGetFileStorage();
  const queryClient = useQueryClient();

  const {
    data: proposal,
    isLoading,
    refetch: refetchProposal,
  } = useGetFinalProjectByProposalId(modalState?.id as string);

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    reset,
    setValue,
  } = useForm<FormUploadPropFinalProjectosalSchemaType>({
    resolver: zodResolver(FormUploadFinalProjectProposalSchema),
    defaultValues: {
      title_id: "",
      title_en: "",
      abstract: "",
      topic: "",
    },
  });

  const toggleModal = () => {
    reset();
    setModalState((prev) => ({
      ...prev,
      open: false,
      id: undefined,
    }));
  };

  const onSubmit = async (data: FormUploadPropFinalProjectosalSchemaType) => {
    const reqBody = new FormData();

    console.log(data, "<< payload");

    if (data.title_id) reqBody.append("title_id", data.title_id);
    if (data.title_en) reqBody.append("title_en", data.title_en);
    if (data.topic) reqBody.append("topic", data.topic);
    if (data.abstract) reqBody.append("abstract", data.abstract);
    if (data.file) reqBody.append("file", data.file);

    const response = await createFinalProjectProposal(reqBody);

    if (response?.error) {
      return setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "failed",
        message: response.message,
      }));
    }

    toggleModal();

    queryClient.refetchQueries({
      queryKey: ["final-project-proposal"],
    });

    return setModalConfirmationState((prev) => ({
      ...prev,
      open: true,
      message: "Data berhasil diunggah",
      state: "success",
    }));
  };

  useEffect(() => {
    if (modalState.open && modalState.id && modalState.state !== "add") {
      refetchProposal();
    }
  }, [modalState, refetchProposal]);

  useEffect(() => {
    if (proposal && proposal.data && modalState.state !== "add") {
      setValue("title_id", proposal.data.title_id);
      setValue("title_en", proposal.data.title_en);
      setValue("topic", proposal.data.topic);
      setValue("abstract", proposal.data.abstract);
    }
  }, [proposal, reset]);
  return (
    <Modal isOpen={modalState.open} centered>
      <div className="modal-title">
        <div className="border-bottom mx-3 pt-4 pb-2">
          <span className=" w-100 ">
            {modalState.state === "detail"
              ? "Detail proposal TA"
              : "Upload Proposal TA"}
          </span>
        </div>
      </div>
      <ModalBody className="pt-1">
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="mt-3"
          encType="multipart/form-data"
        >
          <Row className="gap-1">
            {/* Title Indonesia */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="title_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Judul Proposal (ID)
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="title_id"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.title_id ? "border border-danger" : ""
                        }`}
                        id="title_id"
                        placeholder="Masukkan Judul Proposal"
                        readOnly={modalState.state === "detail"}
                        disabled={isLoading}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.title_id} />
                </Col>
              </Row>
            </Col>
            {/* Title English */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="title_en"
                    className="form-label mb-0 fw-medium"
                  >
                    Judul Proposal (Eng)
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="title_en"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.title_en ? "border border-danger" : ""
                        }`}
                        id="title_en"
                        placeholder="Masukkan Judul Proposal"
                        readOnly={modalState.state === "detail"}
                        disabled={isLoading}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.title_en} />
                </Col>
              </Row>
            </Col>
            {/* Topic */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label htmlFor="topic" className="form-label mb-0 fw-medium">
                    Topik
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="topic"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.topic ? "border border-danger" : ""
                        }`}
                        id="topic"
                        placeholder="Masukkan Topik"
                        readOnly={modalState.state === "detail"}
                        disabled={isLoading}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.topic} />
                </Col>
              </Row>
            </Col>
            {/* Abstract */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="abstract"
                    className="form-label mb-0 fw-medium"
                  >
                    Abstrak
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="abstract"
                    control={control}
                    render={({ field }) => (
                      <textarea
                        className={`form-control form-control-icon ${
                          errors.abstract ? "border border-danger" : ""
                        }`}
                        id="abstract"
                        placeholder="Masukkan abstrak proposal"
                        {...field}
                        style={{ height: "100px" }}
                        readOnly={modalState.state === "detail"}
                        disabled={isLoading}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.abstract} />
                </Col>
              </Row>
            </Col>
            {/* proposal file */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label htmlFor="file" className="form-label mb-0 fw-medium">
                    File Proposal
                  </Label>
                </Col>
                <Col sm={12}>
                  {modalState.state === "detail" ? (
                    <div>
                      {proposal?.data?.file_path ? (
                        <button
                          onClick={() =>
                            getFileStorage(proposal?.data?.file_path)
                          }
                          type="button"
                          disabled={loadingFile}
                          className="bg-transparent border-0 p-0 fs-5"
                          style={{
                            textDecoration: "underline",
                            color: "#495057",
                          }}
                        >
                          {loadingFile ? "Downloading.." : "Lihat file"}
                        </button>
                      ) : (
                        "Tidak ada"
                      )}
                    </div>
                  ) : (
                    <div className="d-flex gap-2">
                      <Controller
                        name="file"
                        control={control}
                        render={({ field: { onChange, value, ...field } }) => (
                          <div className="position-relative w-100">
                            <Input
                              type="text"
                              className={`form-control form-control-icon ${
                                errors.file ? "border border-danger" : ""
                              }`}
                              value={
                                value instanceof File
                                  ? value.name
                                  : splitFileNameUploaded("") || ""
                              }
                              // value={value?.name}
                              placeholder="Pilih File"
                              readOnly
                              disabled={isLoading}
                              {...field}
                            />
                            <Input
                              type="file"
                              id="file"
                              accept=".pdf"
                              onChange={(e) => {
                                const file = e.target.files?.[0];
                                onChange(file);
                              }}
                              disabled={modalState.state === "detail"}
                              hidden
                            />
                            {errors.file ? (
                              <FormErrorMessage errors={errors.file} />
                            ) : (
                              <FormDescription message="File dengan format .pdf maksimal 10mb" />
                            )}
                          </div>
                        )}
                      />
                      <label
                        htmlFor="file"
                        className="btn d-flex align-items-center btn-light"
                        style={{ whiteSpace: "nowrap" }}
                      >
                        <FileUploadIcon className="me-1" /> Upload
                      </label>
                    </div>
                  )}
                </Col>
              </Row>
            </Col>
          </Row>
          <div className="d-flex justify-content-between mt-3 gap-3">
            {modalState.state === "detail" ? (
              <Button
                color="primary"
                onClick={toggleModal}
                type="button"
                style={{
                  border: "1px solid #10487A",
                }}
              >
                Tutup
              </Button>
            ) : (
              <>
                <button
                  onClick={toggleModal}
                  className="bg-transparent text-primary rounded px-3"
                  type="button"
                  style={{ border: "1px solid #10487A" }}
                >
                  Batal
                </button>
                <Button
                  color="primary"
                  className="d-flex flex-grow-0 justify-content-center align-items-center"
                  disabled={isSubmitting}
                >
                  Upload Proposal
                </Button>
              </>
            )}
          </div>
        </form>
      </ModalBody>
    </Modal>
  );
};
