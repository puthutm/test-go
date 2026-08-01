"use client";
import { useModalContext } from "@/lib/hooks/use-modal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import {
  GradeCompositionFormType,
  GradeCompositionSchema,
} from "@/lib/validations/settings/grade-composition/grade-composition";

import { useGetSearchValueElement } from "@/services/api/reference/value-element/get-search";
import { createGradeComposition } from "@/services/api/settings/grade-composition/create-grade-composition";
import { useGetDetailGradeComposition } from "@/services/api/settings/grade-composition/use-get-detail-grade-composition";
import { editGradeComposition } from "@/services/api/settings/grade-composition/update-grade-composition";

import { zodResolver } from "@hookform/resolvers/zod";
import React, { useEffect } from "react";

import { SelectComponent } from "@/components/ui/select";
import { CloseIcon } from "@/components/icons/close";
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
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import { useGetSearchAcademicPeriod } from "@/services/api/data-referensi/academic-period/use-get-search-academic-period";

const ModalGradeComposition = () => {
  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    setValue,
    reset,
    clearErrors,
  } = useForm<GradeCompositionFormType>({
    resolver: zodResolver(GradeCompositionSchema),
    defaultValues: {
      percentage: "",
    },
    mode: "onChange",
  });

  const { data: valueElementOptions } = useGetSearchValueElement({
    page: 1,
    page_size: 1000,
  });

  const { modalState, setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const {
    data: detailGradeCompositionData,
    isLoading: isGradeCompositionLoading,
  } = useGetDetailGradeComposition(modalState?.id as string | null);

  //! maping option element
  const elemntOption = valueElementOptions?.data?.map((item) => ({
    value: item.id,
    label: item.name,
  }));

  // academic period options
  const { data: academicPeriod, isLoading: isLoadingAcademicPeriod } =
    useGetSearchAcademicPeriod();

  const academicPeriodOptions = academicPeriod?.data.map((item) => ({
    label: item.fullname,
    value: item.id,
  }));

  const handleToogleModal = () => {
    setModalState((prev) => ({
      ...prev,
      id: null,
      open: false,
      state: "add",
    }));
    reset();
    clearErrors();
  };

  const onSubmit = async (data: GradeCompositionFormType) => {
    try {
      if (modalState.state === "add") {
        const response = await createGradeComposition(data);
        if (response.error) {
          throw new Error(response.message);
        }

        setModalConfirmationState({
          open: true,
          message: "berhasil menambahkan data komposisi nilai",
          state: "success",
        });
        handleToogleModal();
      } else {
        const response = await editGradeComposition(
          modalState.id as string,
          data
        );
        if (response.error) {
          throw new Error(response.message);
        }

        setModalConfirmationState({
          open: true,
          message: "berhasil mengubah data komposisi nilai",
          state: "success",
        });
        handleToogleModal();
      }
    } catch (err: any) {
      setModalConfirmationState({
        open: true,
        message: err.message,
        state: "failed",
      });
    }
  };

  useEffect(() => {
    if (modalState.state === "edit" && detailGradeCompositionData) {
      setValue("value_element_id", {
        value: detailGradeCompositionData?.data?.value_element_id as string,
        label: detailGradeCompositionData?.data?.value_element_name ?? "",
      });
      setValue("academic_periode_id", {
        value: detailGradeCompositionData?.data?.academic_periode_id as string,
        label: detailGradeCompositionData?.data?.academic_periode_name ?? "",
      });
      setValue(
        "percentage",
        String(detailGradeCompositionData?.data?.percentage)
      );
      setValue(
        "is_passing_requirement",
        detailGradeCompositionData?.data?.is_passing_requirement as boolean
      );
    }
  }, [modalState, detailGradeCompositionData]);

  return (
    <Modal
      isOpen={modalState.open && modalState.state !== "duplicate"}
      centered
    >
      <section className="px-3 py-2  d-flex border-bottom justify-items-center">
        <p
          className="fs-4 fw-semibold p-0 mb-0 flex-grow-1"
          // style={{ color: "#909090" }}
        >
          {modalState.state === "detail" && "Detail Komposisi Nilai"}
          {modalState.state === "edit" && "Ubah Komposisi Nilai"}
          {modalState.state === "add" && "Tambah Komposisi Nilai"}
        </p>
        <Button className="bg-white border-0 p-0" onClick={handleToogleModal}>
          <CloseIcon color="#909090" />
        </Button>
      </section>

      <ModalBody>
        <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
          <Row className="gap-2">
            {/*//! academic period id */}
            <Col sm={12}>
              <Label htmlFor="academic_periode_id" className="form-label mb-1 ">
                Periode Akademik
              </Label>
              <Controller
                name="academic_periode_id"
                control={control}
                render={({ field }) => (
                  <SelectComponent
                    {...field}
                    isDisabled={isLoadingAcademicPeriod}
                    options={academicPeriodOptions as OptionType[]}
                    isLoading={isLoadingAcademicPeriod}
                    placeholder="Pilih Periode Akademik"
                    isError={!!errors.academic_periode_id}
                    id={"academic_periode_id"}
                    // onChange={(value) => {
                    //   setValue("academic_periode_id", value.value);
                    // }}
                  />
                )}
              />
              {errors?.academic_periode_id && (
                <div className="text-danger fs-6">
                  {errors?.academic_periode_id.message}
                </div>
              )}
            </Col>
            {/*//! element id */}
            <Col sm={12}>
              <Label htmlFor="value_element_id" className="form-label mb-1 ">
                Element
              </Label>
              <Controller
                name="value_element_id"
                control={control}
                render={({ field }) => (
                  <SelectComponent
                    {...field}
                    isDisabled={isGradeCompositionLoading}
                    options={elemntOption as OptionType[]}
                    isLoading={isGradeCompositionLoading}
                    placeholder="Pilih Element"
                    isError={!!errors.value_element_id}
                    id={"value_element_id"}
                    // onChange={(value) => {
                    //   setValue("value_element_id", value.value);
                    // }}
                  />
                )}
              />
              {errors?.value_element_id && (
                <div className="text-danger fs-6">
                  {errors?.value_element_id.message}
                </div>
              )}
            </Col>

            {/*//! percentage */}
            <Col sm={12}>
              <Label htmlFor="percentage" className="form-label mb-1 ">
                Persentase
              </Label>
              <Controller
                control={control}
                name="percentage"
                render={({ field }) => (
                  <div className="">
                    <Input
                      {...field}
                      className={`${
                        errors.percentage ? "border border-danger" : ""
                      }`}
                      id="percentage"
                      placeholder="Masukkan Persen "
                      type="text"
                      onChange={(e) => {
                        const { numberValue } = handleInputNumberOnly(e);
                        field.onChange(String(numberValue));
                      }}
                      disabled={
                        isGradeCompositionLoading ||
                        modalState.state === "detail"
                      }
                    />
                  </div>
                )}
              />
              {errors?.percentage && (
                <div className="text-danger fs-6">
                  {errors?.percentage.message}
                </div>
              )}
            </Col>

            {/*//! Syarat Lulus Mata Kuliah */}
            <Col sm={12}>
              <Label
                htmlFor="is_passing_requirement"
                className="form-label mb-1 "
              >
                Syarat Lulus Mata Kuliah
              </Label>
              <Controller
                control={control}
                name="is_passing_requirement"
                render={({ field }) => (
                  <div className="form-check form-switch ">
                    <Input
                      role="switch"
                      type="switch"
                      style={{
                        width: 40,
                        height: 20,
                      }}
                      disabled={
                        modalState.state === "detail" ||
                        isSubmitting ||
                        isGradeCompositionLoading
                          ? true
                          : false
                      }
                      checked={field.value}
                      onChange={(e) => {
                        field.onChange(e.target.checked);
                      }}
                    />
                  </div>
                )}
              />
              {errors?.is_passing_requirement && (
                <div className="text-danger fs-6">
                  {errors?.is_passing_requirement.message}
                </div>
              )}
            </Col>
          </Row>

          <div className="d-flex justify-content-end mt-3 gap-2 ">
            <Button
              type="button"
              className="waves-effect waves-light"
              color="light"
              onClick={handleToogleModal}
              style={{ width: "80px" }}
            >
              Tutup
            </Button>

            {modalState.state !== "detail" && (
              <Button
                className="btn waves-effect waves-light"
                disabled={isSubmitting || isGradeCompositionLoading}
                style={{ width: "80px" }}
                color="primary"
              >
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

export default ModalGradeComposition;
