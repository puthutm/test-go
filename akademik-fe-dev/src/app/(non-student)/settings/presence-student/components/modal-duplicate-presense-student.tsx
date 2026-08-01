"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import React, { useEffect } from "react";
import { Controller, useForm } from "react-hook-form";
import { Button, Col, Label, Modal, ModalBody, Row, Spinner } from "reactstrap";

import { useModalContext } from "@/lib/hooks/use-modal";
import { classMerge } from "@/lib/utils/class-merge";
import { useGetSearchAcademicPeriod } from "@/services/api/data-referensi/academic-period/use-get-search-academic-period";
import {
  formDuplicatePresenceSchema,
  FormDuplicatePresenceSchemaType,
} from "@/lib/validations/settings/presence-student/duplicate-presence-student-validatoin";
import { useGetUnsiaStudyProgram } from "@/services/api/data-referensi/study-program/use-get-unsia-study-program";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { SelectComponent } from "@/components/ui/select";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { duplicatePresenceStudent } from "@/services/api/settings/presence/students/duplicate-presence-student";

export const ModalDuplicatePresenceStudent = ({
  studyProgramId,
}: {
  studyProgramId: string;
}) => {
  const { modalState, setModalState } = useModalContext();

  const { setModalConfirmationState } = useModalConfirmationContext();

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    setValue,
    reset,
  } = useForm<FormDuplicatePresenceSchemaType>({
    resolver: zodResolver(formDuplicatePresenceSchema),
  });

  const { data: academicPeriod, isLoading: isLoadingAcademicPeriod } =
    useGetSearchAcademicPeriod();

  const academicPeriodOptions = academicPeriod?.data?.map((opt) => ({
    label: opt.fullname,
    value: opt.id,
  }));

  const academicPeriodOptionsFiltered = academicPeriodOptions?.filter(
    (period) => period.value !== modalState?.id
  );

  const { data: studyProgram, isLoading: isLoadingStudyProgram } =
    useGetUnsiaStudyProgram();

  const studyProgramOptions = studyProgram?.data?.map((opt) => ({
    label: opt.name,
    value: opt.id,
  }));

  const handleToggleModal = () => {
    setModalState((prev) => ({
      ...prev,
      open: !prev.open,
    }));
    reset();
  };

  const onSubmit = async (payload: FormDuplicatePresenceSchemaType) => {
    try {
      const response = await duplicatePresenceStudent(payload);

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

  // set period
  useEffect(() => {
    if (modalState.open && academicPeriod?.data) {
      const findPeriod = academicPeriodOptions?.find(
        (period) => period.value === modalState.id
      );
      setValue("academic_periode_id_target", {
        label: findPeriod?.label as string,
        value: findPeriod?.value as string,
      });
    }
  }, [modalState.open, academicPeriod, setValue]);

  // set studyprogram
  useEffect(() => {
    if (modalState.open && studyProgramId && studyProgram?.data) {
      const findStudyProgram = studyProgramOptions?.find(
        (data) => data.value === studyProgramId
      );

      if (studyProgramId !== "0")
        return setValue("study_program_id", {
          label: findStudyProgram?.label as string,
          value: findStudyProgram?.value as string,
        });

      setValue("study_program_id", {
        label: "Semua Program Studi",
        value: "0",
      });
    }
  }, [modalState.open, studyProgramId, studyProgram?.data]);

  return (
    <Modal isOpen={modalState.open} centered>
      <div className="d-flex justify-content-between align-items-center pt-3 px-3">
        <p className="fs-4 fw-semibold mb-0 text-black">Duplikasi Presensi</p>
        <Button className="bg-white border-0 p-0" onClick={handleToggleModal}>
          <i className=" ri-close-fill text-black fs-4"></i>
        </Button>
      </div>
      <ModalBody>
        <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
          <Row className="gap-3">
            {/* period target */}
            <Col sm={12}>
              <Row className="align-items-center gap-2">
                <Col sm={12}>
                  <Label
                    htmlFor="academic_periode_id_target"
                    className="form-label mb-0 fw-medium"
                  >
                    Periode Terpilih
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="academic_periode_id_target"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={academicPeriodOptions as OptionType[]}
                        isLoading={isLoadingAcademicPeriod}
                        placeholder="Pilih Mulai Berlaku"
                        isError={!!errors.academic_periode_id_target}
                        id={"academic_periode_id_target"}
                        {...field}
                        isDisabled
                      />
                    )}
                  />
                  <FormErrorMessage
                    errors={errors.academic_periode_id_target}
                  />
                </Col>
              </Row>
            </Col>
            {/* period */}
            <Col sm={12}>
              <Row className="align-items-center gap-2">
                <Col sm={12}>
                  <Label
                    htmlFor="academic_periode_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Periode
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="academic_periode_id"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={academicPeriodOptionsFiltered as OptionType[]}
                        isLoading={isLoadingAcademicPeriod}
                        placeholder="Pilih periode"
                        isError={!!errors.academic_periode_id}
                        id={"academic_periode_id"}
                        {...field}
                        isDisabled={isLoadingAcademicPeriod}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.academic_periode_id} />
                </Col>
              </Row>
            </Col>
            {/* study program */}
            <Col sm={12}>
              <Row className="align-items-center gap-2">
                <Col sm={12}>
                  <Label
                    htmlFor="study_program_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Program Studi
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="study_program_id"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={[
                          {
                            label: "Semua Program Studi",
                            value: "0",
                          },
                          ...((studyProgramOptions as OptionType[]) ?? []),
                        ]}
                        isLoading={isLoadingStudyProgram}
                        placeholder="Pilih Program Studi"
                        isError={!!errors.study_program_id}
                        id={"study_program_id"}
                        {...field}
                        isDisabled
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.study_program_id} />
                </Col>
              </Row>
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
            <Button disabled={isSubmitting} color="primary">
              {isSubmitting ? <Spinner size={"sm"} /> : "Duplikasi"}
            </Button>
          </div>
        </form>
      </ModalBody>
    </Modal>
  );
};
