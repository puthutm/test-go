"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useCallback, useEffect } from "react";
import { Controller, useForm } from "react-hook-form";
import {
  Button,
  Col,
  Input,
  Label,
  Modal,
  ModalBody,
  Row,
  Spinner,
} from "reactstrap";

import { FormErrorMessage } from "@/components/ui/form-error-message";
import { useModalContext } from "@/lib/hooks/use-modal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import {
  FormAcademicYearSchemaType,
  formAcademicYearSchema,
  formAcademicYearSchemaDefaultValues,
} from "@/lib/validations/settings/academic-year";
import { createAcademicYear } from "@/services/api/data-referensi/academic-year/create-academic-year";
import { updateAcademicYear } from "@/services/api/data-referensi/academic-year/update-academic-year";
import { useGetAcademicYearById } from "@/services/api/data-referensi/academic-year/use-get-academic-year-by-id";

export const ModalAcademicYear = () => {
  const { modalState, setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    reset,
    setValue,
  } = useForm<FormAcademicYearSchemaType>({
    resolver: zodResolver(formAcademicYearSchema),
    defaultValues: formAcademicYearSchemaDefaultValues,
  });

  const { data: academicYear, isLoading } = useGetAcademicYearById(
    modalState?.id as string
  );

  const handleToggleModal = () => {
    setModalState((prev) => ({
      ...prev,
      open: !prev.open,
      id: undefined,
    }));
  };

  const handleSetFormValue = useCallback(() => {
    if (academicYear) {
      setValue("name", academicYear?.data?.name);
      setValue("years", academicYear?.data?.years);
    }
  }, [academicYear, modalState.id]);
  const onSubmit = async (payload: FormAcademicYearSchemaType) => {
    try {
      const response =
        modalState.state === "add"
          ? await createAcademicYear(payload)
          : await updateAcademicYear(modalState.id as string, payload);

      if (response.error) {
        return setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "failed",
          message: response.message,
        }));
      }

      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        message: `${
          modalState.state === "add" ? "Tambah" : "Update"
        } data berhasil`,
        state: "success",
      }));

      return handleToggleModal();
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
    if (modalState.state !== "add" && modalState.open) {
      handleSetFormValue();
    } else {
      reset();
    }
  }, [modalState.state, modalState.open, academicYear]);

  return (
    <Modal isOpen={modalState.open} centered>
      <div className="d-flex justify-content-between align-items-center pt-3 px-3">
        <p className="fs-4 fw-semibold mb-0 text-black">
          {modalState.state === "add"
            ? "Tambah Tahun Ajaran"
            : modalState.state === "edit"
            ? "Ubah Tahun Ajaran"
            : "Detail Tahun Ajaran"}
        </p>
        <Button className="bg-white border-0 p-0" onClick={handleToggleModal}>
          <i className=" ri-close-fill text-black fs-4"></i>
        </Button>
      </div>
      <ModalBody>
        <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
          <Row className="gap-3">
            <Col sm={12}>
              <Label htmlFor="years" className="form-label mb-1">
                Tahun Ajaran
              </Label>

              <Controller
                name="years"
                control={control}
                render={({ field }) => (
                  <Input
                    className={`form-control form-control-icon ${
                      errors.years ? "border border-danger" : ""
                    }`}
                    id="years"
                    placeholder="Tahun Ajaran"
                    disabled={isLoading}
                    {...field}
                  />
                )}
              />
              <FormErrorMessage errors={errors.years} />
            </Col>
            <Col sm={12}>
              <Label htmlFor="name" className="form-label mb-1">
                Nama
              </Label>

              <Controller
                name="name"
                control={control}
                render={({ field }) => (
                  <Input
                    className={`form-control form-control-icon ${
                      errors.name ? "border border-danger" : ""
                    }`}
                    id="name"
                    placeholder="Nama"
                    disabled={isLoading}
                    {...field}
                  />
                )}
              />
              <FormErrorMessage errors={errors.name} />
            </Col>
          </Row>
          <div className="d-flex justify-content-end mt-3">
            <Button
              type="button"
              className={`${
                modalState.state === "detail" ? "btn-success" : "btn-light"
              } waves-effect waves-light me-2`}
              onClick={handleToggleModal}
            >
              Tutup
            </Button>
            {modalState.state !== "detail" && (
              <Button disabled={isSubmitting} color="primary">
                {isSubmitting ? (
                  <Spinner size={"sm"} />
                ) : modalState.state === "add" ? (
                  "Tambah"
                ) : (
                  "Ubah"
                )}
              </Button>
            )}
          </div>
        </form>
      </ModalBody>
    </Modal>
  );
};
