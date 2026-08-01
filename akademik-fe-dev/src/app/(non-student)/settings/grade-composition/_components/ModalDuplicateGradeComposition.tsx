"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import React from "react";
import { Controller, useForm } from "react-hook-form";
import {
  Button,
  Col,
  Label,
  Modal,
  ModalBody,
  Row,
  Spinner,
  Input,
} from "reactstrap";

import { useModalContext } from "@/lib/hooks/use-modal";
import { useGetSearchAcademicPeriod } from "@/services/api/data-referensi/academic-period/use-get-search-academic-period";
import {
  DuplicateGradeCompositionFormType,
  DuplicateGradeCompositionSchema,
} from "@/lib/validations/settings/grade-composition/duplicate-grade-composition";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { SelectComponent } from "@/components/ui/select";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { duplicateGradeComposition } from "@/services/api/settings/grade-composition/duplicate-grade-composition";
import { CloseIcon } from "@/components/icons/close";

export const ModalDuplicateGradeComposition = () => {
  const { modalState, setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    reset,
    clearErrors,
    watch,
  } = useForm<DuplicateGradeCompositionFormType>({
    resolver: zodResolver(DuplicateGradeCompositionSchema),
    defaultValues: {
      isOverWrite: false,
    },
  });

  const sourcePeriod = watch("academicPeriodIdSource");

  const { data: academicPeriod, isLoading: isLoadingAcademicPeriod } =
    useGetSearchAcademicPeriod();

  const academicPeriodOptions = academicPeriod?.data?.map((opt) => ({
    label: opt.fullname,
    value: opt.id,
  }));

  const academicPeriodOptionsFiltered = academicPeriodOptions?.filter(
    (opt) => opt.value !== sourcePeriod?.value
  );

  const handleToggleModal = () => {
    setModalState((prev) => ({
      ...prev,
      open: false,
      state: "add",
      id: null,
    }));
    reset();
    clearErrors();
  };

  const onSubmit = async (payload: DuplicateGradeCompositionFormType) => {
    try {
      const response = await duplicateGradeComposition({
        academicPeriodIdSource: payload.academicPeriodIdSource.value,
        academicPeriodIdTarget: payload.academicPeriodIdTarget.value,
        isOverWrite: payload.isOverWrite,
      });

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
        message: "Duplikasi data berhasil",
        state: "success",
      }));

      return handleToggleModal();
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "failed",
        message: error?.toString() || "Something went wrong",
      }));
    }
  };

  return (
    <Modal
      isOpen={modalState.open && modalState.state === "duplicate"}
      centered
    >
      <section className="px-3 py-2 d-flex border-bottom justify-content-between align-items-center">
        <p className="fs-4 fw-semibold p-0 mb-0">Duplikasi Komposisi Nilai</p>
        <Button className="bg-white border-0 p-0" onClick={handleToggleModal}>
          <CloseIcon color="#909090" />
        </Button>
      </section>

      <ModalBody>
        <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
          <Row className="gap-2">
            {/* Source Academic Period */}
            <Col sm={12}>
              <Label
                htmlFor="academicPeriodIdSource"
                className="form-label mb-1 fw-medium"
              >
                Periode Akademik Asal
              </Label>
              <Controller
                name="academicPeriodIdSource"
                control={control}
                render={({ field }) => (
                  <SelectComponent
                    {...field}
                    options={academicPeriodOptions as OptionType[]}
                    isLoading={isLoadingAcademicPeriod}
                    placeholder="Pilih Periode Asal"
                    isError={!!errors.academicPeriodIdSource}
                    id={"academicPeriodIdSource"}
                  />
                )}
              />
              <FormErrorMessage errors={errors.academicPeriodIdSource} />
            </Col>

            {/* Target Academic Period */}
            <Col sm={12}>
              <Label
                htmlFor="academicPeriodIdTarget"
                className="form-label mb-1 fw-medium"
              >
                Periode Akademik Tujuan
              </Label>
              <Controller
                name="academicPeriodIdTarget"
                control={control}
                render={({ field }) => (
                  <SelectComponent
                    {...field}
                    options={academicPeriodOptionsFiltered as OptionType[]}
                    isLoading={isLoadingAcademicPeriod}
                    isDisabled={!sourcePeriod || isLoadingAcademicPeriod}
                    placeholder="Pilih Periode Tujuan"
                    isError={!!errors.academicPeriodIdTarget}
                    id={"academicPeriodIdTarget"}
                  />
                )}
              />
              <FormErrorMessage errors={errors.academicPeriodIdTarget} />
            </Col>

            {/* Overwrite Checkbox */}
            <Col sm={12}>
              <Label
                htmlFor="isOverWrite"
                className="form-label mb-1 fw-medium"
              >
                Overwrite jika data sudah ada
              </Label>
              <Controller
                control={control}
                name="isOverWrite"
                render={({ field }) => (
                  <div className="form-check form-switch">
                    <Input
                      role="switch"
                      type="switch"
                      style={{
                        width: 40,
                        height: 20,
                      }}
                      checked={field.value}
                      onChange={(e) => {
                        field.onChange(e.target.checked);
                      }}
                    />
                  </div>
                )}
              />
            </Col>
          </Row>

          <div className="d-flex justify-content-end mt-3 gap-2">
            <Button
              type="button"
              className="waves-effect waves-light"
              color="light"
              onClick={handleToggleModal}
              style={{ width: "80px" }}
            >
              Tutup
            </Button>
            <Button
              disabled={isSubmitting}
              color="primary"
              className="waves-effect waves-light"
              style={{ minWidth: "100px" }}
            >
              {isSubmitting ? <Spinner size={"sm"} /> : "Duplikasi"}
            </Button>
          </div>
        </form>
      </ModalBody>
    </Modal>
  );
};
