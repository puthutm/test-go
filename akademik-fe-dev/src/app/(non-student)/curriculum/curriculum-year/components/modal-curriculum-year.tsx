"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import React, { useCallback, useEffect } from "react";
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

import { useModalContext } from "@/lib/hooks/use-modal";
import { classMerge } from "@/lib/utils/class-merge";
import {
  CurriculumYearFormType,
  CurriculumYearSchema,
} from "@/lib/validations/academic/settings/curriculum-year";
import { CalendarTodayIcon } from "@/components/icons/calendar-today";
import { DatePicker } from "@/components/ui/date-picker";
import { SelectComponent } from "@/components/ui/select";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { useGetSearchAcademicPeriod } from "@/services/api/data-referensi/academic-period/use-get-search-academic-period";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { createCurriculumYear } from "@/services/api/data-referensi/curriculum-year/create-curriculum-year";
import { updateCurriculumYear } from "@/services/api/data-referensi/curriculum-year/update-curriculum-year";
import { useGetCurriculumYearById } from "@/services/api/data-referensi/curriculum-year/use-get-curriculum-year-by-id";

const ModalCurriculumYear = () => {
  const { modalState, setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const { data: curriculumYear, isLoading: isLoadingCurriculumYear } =
    useGetCurriculumYearById(modalState?.id as string);

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    watch,
    setValue,
    reset,
    resetField,
  } = useForm<CurriculumYearFormType>({
    resolver: zodResolver(CurriculumYearSchema),
    defaultValues: {
      years: "",
      description: "",
      start_date: [],
      end_date: [],
    },
  });

  const { data: academicPeriod, isLoading: isLoadingAcademicPeriod } =
    useGetSearchAcademicPeriod();

  const academicPeriodOptions = academicPeriod?.data.map((item) => ({
    label: item.fullname,
    value: item.id,
  }));

  const handleToggleModal = () => {
    setModalState((prev) => ({
      ...prev,
      open: !prev.open,
    }));
    reset();
  };

  const startDate = watch("start_date");

  const onSubmit = async (data: CurriculumYearFormType) => {
    try {
      const response =
        modalState.state === "add"
          ? await createCurriculumYear(data)
          : await updateCurriculumYear(modalState.id as string, data);

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
        message: error?.toString() || "Something went wrong",
      }));
    }
  };

  const handleSetFormValue = useCallback(() => {
    setValue("years", curriculumYear?.data?.years as string);
    setValue("starts", {
      label: curriculumYear?.data?.academic_periode_name as string,
      value: curriculumYear?.data?.starts as string,
    });
    setValue("start_date", [
      new Date(curriculumYear?.data?.start_date as string),
    ]);
    setValue("end_date", [new Date(curriculumYear?.data?.end_date as string)]);
    if (curriculumYear?.data?.description) {
      setValue("description", curriculumYear?.data?.description as string);
    }
  }, [modalState, curriculumYear, setValue]);

  useEffect(() => {
    if (curriculumYear?.data) {
      handleSetFormValue();
    }
  }, [curriculumYear?.data, handleSetFormValue]);

  useEffect(() => {
    if (modalState.state === "add") {
      reset();
      resetField("starts", undefined);
    }
  }, [modalState]);

  return (
    <Modal isOpen={modalState.open} centered>
      <div className="d-flex justify-content-between align-items-center pt-3 px-3">
        <p className="fs-4 fw-semibold mb-0 text-black">
          {modalState.state === "add"
            ? "Tambah Tahun Kurikulum"
            : modalState.state === "edit"
            ? "Ubah Tahun Kurikulum"
            : "Detail Tahun Kurikulum"}
        </p>
        <Button className="bg-white border-0 p-0" onClick={handleToggleModal}>
          <i className=" ri-close-fill text-black fs-4"></i>
        </Button>
      </div>

      <ModalBody>
        <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
          <Row className="gap-3">
            {/* years */}
            <Col sm={12}>
              <Label htmlFor="years" className="form-label mb-1">
                Nama Tahun Kurikulum
              </Label>
              <Col sm={12}>
                <Controller
                  name="years"
                  control={control}
                  render={({ field }) => (
                    <Input
                      className={`form-control form-control-icon ${
                        errors.years ? "border border-danger" : ""
                      }`}
                      id="years"
                      placeholder="Masukkan tahun kurikulum"
                      disabled={
                        modalState.state === "detail" || isLoadingCurriculumYear
                      }
                      {...field}
                    />
                  )}
                />
                <FormErrorMessage errors={errors.years} />
              </Col>
            </Col>
            {/* starts */}
            <Col sm={12}>
              <Row className="align-items-center gap-2">
                <Col sm={12}>
                  <Label htmlFor="starts" className="form-label mb-0 fw-medium">
                    Mulai Berlaku
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="starts"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={academicPeriodOptions as OptionType[]}
                        isLoading={isLoadingAcademicPeriod}
                        placeholder="Pilih Mulai Berlaku"
                        isError={!!errors.starts}
                        id={"starts"}
                        {...field}
                        isDisabled={
                          modalState.state === "detail" ||
                          isLoadingCurriculumYear
                        }
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.starts} />
                </Col>
              </Row>
            </Col>
            {/* start date */}
            <Col sm={12}>
              <Row className="align-items-center gap-2">
                <Col sm={12}>
                  <Label
                    htmlFor="number_of_study_program"
                    className="form-label mb-0 fw-medium"
                  >
                    Tanggal Mulai
                  </Label>
                </Col>
                <Col sm={12}>
                  <div className="form-icon">
                    <Controller
                      name="start_date"
                      control={control}
                      render={({ field }) => {
                        return (
                          <DatePicker
                            onChange={(e) => field.onChange(e)}
                            value={field.value}
                            className={`p-0 ${
                              errors.start_date ? "border border-danger" : ""
                            }`}
                            classNameFlatpickr={`form-control form-control-icon disabled-input`}
                            options={{
                              mode: "single",
                              dateFormat: "d F Y",
                            }}
                            disabled={
                              modalState.state === "detail" ||
                              isLoadingCurriculumYear
                            }
                          />
                        );
                      }}
                    />
                    <i style={{ left: "15px" }}>
                      <CalendarTodayIcon color="#878A99" />
                    </i>
                  </div>
                  <FormErrorMessage errors={errors.start_date} />
                </Col>
              </Row>
            </Col>
            {/* end date */}
            <Col sm={12}>
              <Row className="align-items-center gap-2">
                <Col sm={12}>
                  <Label
                    htmlFor="number_of_study_program"
                    className="form-label mb-0 fw-medium"
                  >
                    Tanggal Selesai
                  </Label>
                </Col>
                <Col sm={12}>
                  <div className="form-icon">
                    <Controller
                      name="end_date"
                      control={control}
                      render={({ field }) => {
                        return (
                          <DatePicker
                            onChange={(e) => field.onChange(e)}
                            value={field.value}
                            className={`p-0 ${
                              errors.end_date ? "border border-danger" : ""
                            }`}
                            classNameFlatpickr={`form-control form-control-icon disabled-input`}
                            options={{
                              mode: "single",
                              dateFormat: "d F Y",
                              minDate: startDate[0],
                            }}
                            disabled={
                              modalState.state === "detail" ||
                              isLoadingCurriculumYear
                            }
                          />
                        );
                      }}
                    />
                    <i style={{ left: "15px" }}>
                      <CalendarTodayIcon color="#878A99" />
                    </i>
                  </div>
                  <FormErrorMessage errors={errors.end_date} />
                </Col>
              </Row>
            </Col>
            {/* description */}
            <Col sm={12}>
              <Label htmlFor="description" className="form-label mb-1 optional">
                Deskripsi
              </Label>
              <Controller
                name="description"
                control={control}
                render={({ field }) => (
                  <Input
                    className={`form-control form-control-icon ${
                      errors.description ? "border border-danger" : ""
                    }`}
                    id="description"
                    placeholder="Masukkan deskripsi"
                    type="textarea"
                    disabled={
                      modalState.state === "detail" || isLoadingCurriculumYear
                    }
                    {...field}
                  />
                )}
              />
              <FormErrorMessage errors={errors.description} />
            </Col>
          </Row>

          <div className="d-flex justify-content-end mt-3">
            <Button
              type="button"
              className={classMerge(
                modalState.state === "detail" ? "btn-success" : "btn-light",
                " waves-effect waves-light me-2"
              )}
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

export default ModalCurriculumYear;
