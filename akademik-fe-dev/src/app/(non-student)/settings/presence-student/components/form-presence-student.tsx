"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import React, { useEffect, useMemo } from "react";
import { Controller, useForm } from "react-hook-form";
import { Button, Col, Input, Label, Progress, Row } from "reactstrap";

import { useGetUnsiaStudyProgram } from "@/services/api/data-referensi/study-program/use-get-unsia-study-program";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { classMerge } from "@/lib/utils/class-merge";
import { useGetSearchAcademicPeriod } from "@/services/api/data-referensi/academic-period/use-get-search-academic-period";
import { SelectComponent } from "@/components/ui/select";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import {
  formPresenceSchema,
  FormPresenceSchemaType,
} from "@/lib/validations/settings/presence-student/create-presence-student-validation";
import { useRouter } from "next/navigation";
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import { createPresenceStudent } from "@/services/api/settings/presence/students/create-presence-student";

export const FormPresenceStudent = () => {
  const router = useRouter();

  const { setModalConfirmationState } = useModalConfirmationContext();

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    watch,
    setValue,
  } = useForm<FormPresenceSchemaType>({
    resolver: zodResolver(formPresenceSchema),
    mode: "onChange", // <— live validation
    reValidateMode: "onChange",
    defaultValues: {
      use_comment: false,
      use_document_material: false,
      use_open_session: false,
      use_quiz: false,
      use_task: false,
      use_uas: false,
      use_uts: false,
      use_video: false,
      open_session_percentage: 0,
      task_percentage: 0,
      video_percentage: 0,
      comment_percentage: 0,
      quiz_percentage: 0,
      document_material_percentage: 0,
      uts_percentage: 0,
      uas_percentage: 0,
    },
  });

  const useOpenSession = watch("use_open_session");
  const useTask = watch("use_task");
  const useComment = watch("use_comment");
  const useQuiz = watch("use_quiz");
  const useVideo = watch("use_video");
  const useDocumentMaterial = watch("use_document_material");
  const useUts = watch("use_uts");
  const useUas = watch("use_uas");

  const openSessionPercentage = watch("open_session_percentage");
  const taskPercentage = watch("task_percentage");
  const videoPercentage = watch("video_percentage");
  const commentPercentage = watch("comment_percentage");
  const quizPercentage = watch("quiz_percentage");
  const documentMaterialPercentage = watch("document_material_percentage");

  const totalPercentage = useMemo(() => {
    const total =
      Number(openSessionPercentage) +
      Number(taskPercentage) +
      Number(videoPercentage) +
      Number(commentPercentage) +
      Number(quizPercentage) +
      Number(documentMaterialPercentage);
    return total;
  }, [
    openSessionPercentage,
    taskPercentage,
    videoPercentage,
    commentPercentage,
    quizPercentage,
    documentMaterialPercentage,
  ]);

  const { data: academicPeriod, isLoading: isLoadingAcademicPeriod } =
    useGetSearchAcademicPeriod();

  const academicPeriodOptions = academicPeriod?.data?.map((opt) => ({
    label: opt.fullname,
    value: opt.id,
  }));

  const { data: studyProgram, isLoading: isLoadingStudyProgram } =
    useGetUnsiaStudyProgram();

  const studyProgramOptions = studyProgram?.data?.map((opt) => ({
    label: opt.name,
    value: opt.id,
  }));

  const studentPresences = [
    {
      title: "Open Sesi",
      key: "open_session_percentage",
      toggle: "use_open_session",
    },
    {
      title: "Materi",
      key: "document_material_percentage",
      toggle: "use_document_material",
    },
    {
      title: "Quis",
      key: "quiz_percentage",
      toggle: "use_quiz",
    },
    {
      title: "Tugas",
      key: "task_percentage",
      toggle: "use_task",
    },
    {
      title: "Video",
      key: "video_percentage",
      toggle: "use_video",
    },
    {
      title: "Komentar",
      key: "comment_percentage",
      toggle: "use_comment",
    },
    {
      title: "UTS",
      key: "uts_percentage",
      toggle: "use_uts",
    },
    {
      title: "UAS",
      key: "uas_percentage",
      toggle: "use_uas",
    },
  ];

  const onSubmit = async (payload: FormPresenceSchemaType) => {
    try {
      const response = await createPresenceStudent(payload);

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
        message: `Pengaturan presensi berhasil ditambah`,
        state: "success",
      }));

      router.push("/settings/presence-student");

      return response;

      // return handleToggleModal();
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "failed",
        message: error?.toString() || "Something went wrong",
      }));
    }
  };

  useEffect(() => {
    if (!useComment) setValue("comment_percentage", 0);
    if (!useDocumentMaterial) setValue("document_material_percentage", 0);
    if (!useComment) setValue("comment_percentage", 0);
    if (!useQuiz) setValue("quiz_percentage", 0);
    if (!useOpenSession) setValue("open_session_percentage", 0);
    if (!useVideo) setValue("video_percentage", 0);
    if (!useTask) setValue("task_percentage", 0);
    if (!useUts) {
      setValue("uts_percentage", 0);
    } else {
      setValue("uts_percentage", 100);
    }
    if (!useUas) {
      setValue("uas_percentage", 0);
    } else {
      setValue("uas_percentage", 100);
    }
  }, [
    useComment,
    useDocumentMaterial,
    useComment,
    useQuiz,
    useOpenSession,
    useVideo,
    useTask,
    useUts,
    useUas,
  ]);
  return (
    <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
      <Row className="gap-3">
        {/* acedemic period */}
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
                    options={academicPeriodOptions as OptionType[]}
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
          <Row className="gap-2">
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
                    placeholder="Pilih Mulai Berlaku"
                    isError={!!errors.study_program_id}
                    id={"study_program_id"}
                    {...field}
                    isDisabled={isLoadingStudyProgram}
                  />
                )}
              />
              <FormErrorMessage errors={errors.study_program_id} />
            </Col>
          </Row>
        </Col>
        <Col sm={12}>
          <button className="border-0 rounded-top text-center py-2 px-4 fw-semibold bg-primary text-white">
            Komponen Presensi
          </button>
          <div className="pb-3" style={{ borderBottom: "2px solid #DEE5EC" }} />
        </Col>
        {/* component presence percentage */}
        <Col sm={12}>
          <Row>
            {/* input percentage */}
            <Col sm={12} lg={8}>
              <div className="d-flex flex-column gap-3 border rounded-3 p-3">
                {studentPresences?.map((component) => (
                  <div
                    className={`d-flex  border rounded-2 py-2 px-3 ${
                      totalPercentage > 100 ? "border-danger" : ""
                    } ${
                      component.key === "uts_percentage" ||
                      component.key === "uas_percentage"
                        ? "justify-content-between"
                        : "flex-column "
                    }`}
                    key={component.key}
                  >
                    <p>{component.title}</p>
                    <div className="d-flex align-items-center gap-4">
                      <div
                        style={{ width: "70%" }}
                        className={`${
                          component.key === "uts_percentage" ||
                          component.key === "uas_percentage"
                            ? "d-none"
                            : "d-block"
                        }`}
                      >
                        <Progress value={watch(component.key as any)} />
                      </div>
                      <Controller
                        name={component.key as any}
                        control={control}
                        render={({ field }) => (
                          <input
                            type="text"
                            {...field}
                            className={`form-input p-2 rounded-3 ${
                              component.key === "uts_percentage" ||
                              component.key === "uas_percentage"
                                ? "d-none"
                                : "d-block"
                            }`}
                            style={{
                              border: "1px solid #DEE5EC",
                              width: "15%",
                            }}
                            onChange={(e) => {
                              const { numberValue } = handleInputNumberOnly(e);

                              return field.onChange(numberValue);
                            }}
                            disabled={Boolean(!watch(component.toggle as any))}
                          />
                        )}
                      />
                      <span
                        className={`fs-5 ${
                          component.key === "uts_percentage" ||
                          component.key === "uas_percentage"
                            ? "d-none"
                            : "d-block"
                        }`}
                      >
                        %
                      </span>
                      <div className=" form-check form-switch ms-1">
                        <Controller
                          name={component.toggle as any}
                          control={control}
                          render={({ field }) => (
                            <Input
                              type="checkbox"
                              {...field}
                              className="form-check-input"
                              role="switch"
                              value={field.value ? "true" : "false"}
                              checked={field.value === true}
                              onChange={(e) => field.onChange(e.target.checked)}
                            />
                          )}
                        />
                      </div>
                    </div>
                  </div>
                ))}
              </div>
              <div className="d-flex justify-content-end flex gap-2 mt-4">
                <Button
                  type="button"
                  className={classMerge(
                    "btn-light waves-effect waves-light me-2 "
                  )}
                  onClick={() => router.push("/settings/presence-student")}
                >
                  Kembali
                </Button>
                <Button color="primary" disabled={isSubmitting}>
                  Simpan
                </Button>
              </div>
            </Col>
            {/* total percentage and the component */}
            <Col sm={12} lg={4}>
              <div className="d-flex flex-column gap-3">
                {/* total percentage */}
                <div className="d-flex flex-column gap-1 px-3">
                  <p className="fs-5">Total Persentasi</p>
                  <div className="d-flex align-items-center gap-3">
                    <div className="w-100">
                      <Progress value={totalPercentage} />
                    </div>
                    <span className="fs-5">{totalPercentage}</span>
                    <span className="fs-5">%</span>
                  </div>
                  {totalPercentage > 100 ? (
                    <span className="text-danger fst-italic">
                      Persentasi tidak boleh lebih dari 100%
                    </span>
                  ) : null}
                </div>
                {/* presence component */}
                <div className="d-flex flex-column gap-3 rounded-3 border p-3">
                  <p className="fs-5">Komponen</p>
                  <span className="text-center">Tidak ada komponen</span>
                </div>
              </div>
            </Col>
          </Row>
        </Col>
      </Row>
    </form>
  );
};
