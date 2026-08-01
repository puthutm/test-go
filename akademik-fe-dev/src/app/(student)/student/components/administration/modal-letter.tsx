"use client";

import {
  Button,
  Modal,
  ModalHeader,
  ModalBody,
  Row,
  Col,
  Label,
  Input,
} from "reactstrap";

import { useModalContext } from "@/lib/hooks/use-modal";
import { Controller, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { FormMailSchema } from "@/lib/validations/students/administration/form-mail-schema";
import { SelectComponent } from "@/components/ui/select";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { FileUploadIcon } from "@/components/icons/file-upload";
import { FormDescription } from "@/components/ui/form-description";

export const ModalLetter = () => {
  const { modalState, setModalState } = useModalContext();

  const {
    formState: { errors },
    handleSubmit,
    control,
    clearErrors,
  } = useForm({
    resolver: zodResolver(FormMailSchema),
    defaultValues: {
      reason_mail: "",
    },
  });

  const onSubmit = async () => {
    console.log("Click");
  };

  return (
    <Modal isOpen={modalState.open} centered>
      <ModalHeader className="d-flex justify-content-between border-bottom px-0 mx-4 py-3">
        Ajukan Surat Baru
      </ModalHeader>
      <ModalBody className="py-2">
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="my-2"
          autoComplete="off"
        >
          <Row className="gap-2">
            {/* type letter */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="type_mail_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Tempat Lahir
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="type_mail_id"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={[] as OptionType[]}
                        // isLoading={isLoadingCities}
                        placeholder="Pilih Keperluan"
                        isError={!!errors.type_mail_id}
                        id={"type_mail_id"}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.type_mail_id} />
                </Col>
              </Row>
            </Col>
            {/* reason_mail */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="reason_mail"
                    className="form-label mb-0 fw-medium"
                  >
                    Keperluan Surat
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="reason_mail"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.reason_mail ? "border border-danger" : ""
                        }`}
                        id="reason_mail"
                        placeholder="Masukkan Keperluan Surat"
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.reason_mail} />
                </Col>
              </Row>
            </Col>
            {/* document */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="document"
                    className="form-label mb-0 fw-medium  optional"
                  >
                    Dokument Pendukung
                  </Label>
                </Col>
                <Col sm={12}>
                  <div className="d-flex gap-2">
                    <Controller
                      name="document"
                      control={control}
                      render={({ field }) => (
                        <div className="position-relative w-100">
                          <Input
                            type="text"
                            className={`form-control form-control-icon ${
                              errors.document ? "border border-danger" : ""
                            }`}
                            placeholder="Pilih File"
                          />
                          {modalState.state !== "detail" && (
                            <Input
                              type="file"
                              id="document"
                              accept=".pdf"
                              onChange={(e) => {
                                const file = e.target.files?.[0];
                                if (file) field.onChange(file);
                              }}
                              hidden
                            />
                          )}
                          {errors.document ? (
                            <FormErrorMessage errors={errors.document} />
                          ) : (
                            <FormDescription message="File dengan format .pdf maksimal 2mb" />
                          )}
                        </div>
                      )}
                    />
                    {modalState.state !== "detail" ? (
                      <label
                        htmlFor="document"
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
                      >
                        Lihat File
                      </button>
                    )}
                  </div>
                </Col>
              </Row>
            </Col>
            {/* type letter */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="send_mail"
                    className="form-label mb-0 fw-medium"
                  >
                    Kirim Surat Melalui
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="send_mail"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={[] as OptionType[]}
                        // isLoading={isLoadingCities}
                        placeholder="Pilih Kanal"
                        isError={!!errors.send_mail}
                        id={"send_mail"}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.send_mail} />
                </Col>
              </Row>
            </Col>
          </Row>
          <div className="d-flex justify-content-between mt-3">
            <Button
              color="outline"
              onClick={() => {
                setModalState((prev) => ({
                  ...prev,
                  open: false,
                }));
                clearErrors();
              }}
              style={{
                whiteSpace: "nowrap",
                border: "1px solid #10487A",
                backgroundColor: "transparent",
                color: "#10487A",
              }}
              type="button"
            >
              Batal
            </Button>
            <Button color="primary">Ajukan Surat Baru</Button>
          </div>
        </form>
      </ModalBody>
    </Modal>
  );
};
