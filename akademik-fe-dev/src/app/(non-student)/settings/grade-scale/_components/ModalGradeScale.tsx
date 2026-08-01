"use client";

import { useModalContext } from "@/lib/hooks/use-modal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { CloseIcon } from "@/components/icons/close";
import { SelectComponent } from "@/components/ui/select";

import {
  GradeScaleFormType,
  GradeScaleSchema,
} from "@/lib/validations/settings/grade-scale/form-grade-scale";

import { useGrade } from "@/services/api/data-referensi/grade/use-grade";
import { useGetUnsiaStudyProgram } from "@/services/api/data-referensi/study-program/use-get-unsia-study-program";
import { createGradeScale } from "@/services/api/settings/grade-scale/create-grade-scale";
import { useGetDetailGradeScale } from "@/services/api/settings/grade-scale/use-get-detail-grade-scale";
import { editGradeScale } from "@/services/api/settings/grade-scale/update-grade-scale";
import { zodResolver } from "@hookform/resolvers/zod";

import React, { useEffect } from "react";
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

const ModalGradeScale = () => {
  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    setValue,
    reset,
    clearErrors,
    watch,
  } = useForm<GradeScaleFormType>({
    resolver: zodResolver(GradeScaleSchema),
    defaultValues: {
      weight_value: "",
      lower_value: "",
      upper_value: "",
      description: "",
    },
    mode: "onChange",
  });

  const gradeId = watch("grade_id");

  const { modalState, setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const { data: detailGradeScaleData, isLoading: isGradeScaleLoading } =
    useGetDetailGradeScale(modalState?.id as string | null);

  const { data: programStudy, isLoading: isLoadingProgramStudy } =
    useGetUnsiaStudyProgram();

  const { data: grade, isLoading: isLoadingGrade } = useGrade(modalState.open);

  //! mpaing option study program
  const mapingStudyProgramData =
    programStudy?.data?.map((el: UnsiaStudyProgram) => {
      return {
        label: el.name,
        value: el.id,
      };
    }) ?? [];

  //! maping option grade
  const mapingGrade =
    grade?.data?.map((el: GradeOptions) => {
      return {
        label: el.name,
        value: el.id,
      };
    }) ?? [];

  const handleToogleModal = () => {
    setModalState((prev) => ({
      ...prev,
      id: null,
      open: false,
    }));
    reset();
    clearErrors();
  };

  const onSubmit = async (data: GradeScaleFormType) => {
    try {
      if (modalState.state === "add") {
        const response = await createGradeScale(data);
        if (response.error) {
          throw new Error(response.message);
        }

        setModalConfirmationState({
          open: true,
          message: "berhasil menambahkan data skala nilai",
          state: "success",
        });
        handleToogleModal();
      } else {
        const response = await editGradeScale(modalState.id as string, data);
        if (response.error) {
          throw new Error(response.message);
        }

        setModalConfirmationState({
          open: true,
          message: "berhasil mengubah data skala nilai",
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
    if (modalState.state === "edit" && detailGradeScaleData) {
      setValue("study_program_id", {
        label: detailGradeScaleData?.data?.study_program_name as string,
        value: detailGradeScaleData?.data?.study_program_id as string,
      });
      setValue("grade_id", {
        label: detailGradeScaleData?.data?.grade_name as string,
        value: detailGradeScaleData?.data?.grade_id as string,
      });
      setValue(
        "weight_value",
        String(detailGradeScaleData?.data?.weight_value as number)
      );
      setValue(
        "lower_value",
        String(detailGradeScaleData?.data?.lower_value as number)
      );
      setValue(
        "upper_value",
        String(detailGradeScaleData?.data?.upper_value as number)
      );
      setValue(
        "description",
        detailGradeScaleData?.data?.description as string
      );
    }
  }, [modalState, detailGradeScaleData]);

  useEffect(() => {
    if (gradeId?.value) {
      const findGrade = grade?.data?.find((val) => val.id === gradeId.value);
      setValue("lower_value", findGrade?.lower_limit as string);
      setValue("upper_value", findGrade?.upper_limit as string);
    }
  }, [gradeId?.value]);

  return (
    <Modal isOpen={modalState.open} centered>
      <section className="px-3 py-2  d-flex border-bottom justify-items-center">
        <p
          className="fs-4 fw-semibold p-0 mb-0 flex-grow-1"
          style={{ color: "#909090" }}
        >
          {modalState.state === "detail" && "Detail Skala Nilai"}
          {modalState.state === "edit" && "Ubah Skala Nilai"}
          {modalState.state === "add" && "Tambah Skala Nilai"}
        </p>
        <Button className="bg-white border-0 p-0" onClick={handleToogleModal}>
          <CloseIcon color="#909090" />
        </Button>
      </section>

      <ModalBody>
        <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
          <Row className="gap-2">
            {/* study program */}
            <Col sm={12}>
              <Label htmlFor="study_program_id" className="form-label mb-1 ">
                Program Studi
              </Label>
              <Controller
                name="study_program_id"
                control={control}
                render={({ field }) => (
                  <SelectComponent
                    {...field}
                    isDisabled={isGradeScaleLoading}
                    options={mapingStudyProgramData as OptionType[]}
                    isLoading={isLoadingProgramStudy}
                    placeholder="Pilih Study Program"
                    isError={!!errors.study_program_id}
                    id={"study_program_id"}
                    // onChange={(value) => {
                    //   setValue("study_program_id", value.value);
                    // }}
                  />
                )}
              />
              {errors?.study_program_id && (
                <div className="text-danger fs-6">
                  {errors?.study_program_id.message}
                </div>
              )}
            </Col>

            {/* Grade */}
            <Col sm={12}>
              <Label htmlFor="study_program_id" className="form-label mb-1">
                Nilai
              </Label>
              <Controller
                name="grade_id"
                control={control}
                render={({ field }) => (
                  <SelectComponent
                    {...field}
                    isDisabled={isGradeScaleLoading}
                    options={mapingGrade as OptionType[]}
                    isLoading={isLoadingGrade}
                    placeholder="Pilih Nilai"
                    isError={!!errors.grade_id}
                    id={"grade_id"}
                  />
                )}
              />
              {errors?.grade_id && (
                <div className="text-danger fs-6">
                  {errors?.grade_id.message}
                </div>
              )}
            </Col>

            {/* Bobot Nilai (weigth_value) */}
            <Col sm={12}>
              <Label htmlFor="weigth_value" className="form-label mb-1 ">
                Bobot Nilai
              </Label>
              <Controller
                control={control}
                name="weight_value"
                render={({ field }) => (
                  <div className="">
                    <Input
                      {...field}
                      className={`${
                        errors.weight_value ? "border border-danger" : ""
                      }`}
                      id="weight_value"
                      placeholder="Masukkan bobot nilai"
                      type="text"
                      onChange={(e) => {
                        const { numberValue } = handleInputNumberOnly(e);
                        field.onChange(String(numberValue));
                      }}
                      disabled={
                        isGradeScaleLoading || modalState.state === "detail"
                      }
                    />
                  </div>
                )}
              />
              {errors?.weight_value && (
                <div className="text-danger fs-6">
                  {errors?.weight_value.message}
                </div>
              )}
            </Col>

            {/* Nilai Bawah (lower_value) */}
            <Col sm={12}>
              <Label htmlFor="lower_value" className="form-label mb-1 ">
                Nilai Bawah
              </Label>
              <Controller
                control={control}
                name="lower_value"
                render={({ field }) => (
                  <div className="">
                    <Input
                      {...field}
                      className={`${
                        errors.lower_value && "border border-danger"
                      }`}
                      id="lower_value"
                      placeholder="Masukkan nilai bawah"
                      type="text"
                      onChange={(e) => {
                        const { numberValue } = handleInputNumberOnly(e);
                        field.onChange(String(numberValue));
                      }}
                      disabled={
                        isGradeScaleLoading || modalState.state === "detail"
                      }
                    />
                  </div>
                )}
              />
              {errors?.lower_value && (
                <div className="text-danger fs-6">
                  {errors?.lower_value.message}
                </div>
              )}
            </Col>

            {/* Nilai Atas (upper_value) */}
            <Col sm={12}>
              <Label htmlFor="upper_value" className="form-label mb-1 ">
                Nilai Atas
              </Label>
              <Controller
                control={control}
                name="upper_value"
                render={({ field }) => (
                  <div className="">
                    <Input
                      {...field}
                      className={`${
                        errors.upper_value && "border border-danger"
                      }`}
                      id="upper_value"
                      placeholder="Masukkan nilai atas"
                      type="text"
                      onChange={(e) => {
                        const { numberValue } = handleInputNumberOnly(e);
                        field.onChange(String(numberValue));
                      }}
                      disabled={
                        isGradeScaleLoading || modalState.state === "detail"
                      }
                    />
                  </div>
                )}
              />
              {errors?.upper_value && (
                <div className="text-danger fs-6">
                  {errors?.upper_value.message}
                </div>
              )}
            </Col>

            {/* Deskripsi */}
            <Col sm={12}>
              <Label htmlFor="description" className="form-label mb-1">
                Deskripsi
              </Label>
              <Controller
                control={control}
                name="description"
                render={({ field }) => (
                  <div className="">
                    <Input
                      {...field}
                      className={`${
                        errors.description && "border border-danger"
                      }`}
                      id="description"
                      placeholder="Masukkan deskripsi"
                      type="textarea"
                      disabled={
                        isGradeScaleLoading || modalState.state === "detail"
                      }
                    />
                  </div>
                )}
              />
              {errors?.description && (
                <div className="text-danger">{errors?.description.message}</div>
              )}
            </Col>
          </Row>

          <div className="d-flex justify-content-end mt-3 gap-2">
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
                disabled={isSubmitting || isGradeScaleLoading}
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

export default ModalGradeScale;
