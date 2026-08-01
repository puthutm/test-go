"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import React, { useEffect, useMemo } from "react";
import { Controller, useForm } from "react-hook-form";
import { Button, Col, Input, Progress, Row } from "reactstrap";

import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import {
  formPresenceComponentSchema,
  FormPresenceComponentSchemaType,
} from "@/lib/validations/settings/presence-student/create-presence-student-for-lecturer-validation";
import { createOrUpdatePresenceComponentStudent } from "@/services/api/settings/presence/lecturer/create-or-update-presence-component";
import { useParams, useSearchParams } from "next/navigation";

interface Props {
  presenceComponent: ApiResponse<PresenceComponent> | undefined;
}

const FormPresenceComponent = ({ presenceComponent }: Props) => {
  const { setModalConfirmationState } = useModalConfirmationContext();

  const searchParams = useSearchParams();
  const params = useParams();

  const studyProgram = searchParams.get("studyProgram");
  const academicPeriod = searchParams.get("period");
  const subjectId = params.subjectId;

  const component = presenceComponent?.data;

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    watch,
    setValue,
  } = useForm<FormPresenceComponentSchemaType>({
    resolver: zodResolver(formPresenceComponentSchema),
    mode: "onChange",
    reValidateMode: "onChange",
    defaultValues: {
      study_program_id: studyProgram ?? "",
      use_comment: component?.use_comment ?? false,
      use_document_material: component?.use_document_material ?? false,
      use_open_session: component?.use_open_session ?? false,
      use_quiz: component?.use_quiz ?? false,
      use_task: component?.use_task ?? false,
      use_uas: component?.use_uas ?? false,
      use_uts: component?.use_uts ?? false,
      use_video: component?.use_video ?? false,
      open_session_percentage: component?.open_session_percentage ?? 0,
      task_percentage: component?.task_percentage ?? 0,
      video_percentage: component?.video_percentage ?? 0,
      comment_percentage: component?.comment_percentage ?? 0,
      quiz_percentage: component?.quiz_percentage ?? 0,
      document_material_percentage:
        component?.document_material_percentage ?? 0,
      uts_percentage: component?.uts_percentage ?? 0,
      uas_percentage: component?.uas_percentage ?? 0,
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
  const utsPercentage = watch("uts_percentage");
  const uasPercentage = watch("uas_percentage");

  const totalPercentage = useMemo(() => {
    const openSession = useOpenSession ? openSessionPercentage : 0;
    const documentMaterial = useDocumentMaterial
      ? documentMaterialPercentage
      : 0;
    const quiz = useQuiz ? quizPercentage : 0;
    const task = useTask ? taskPercentage : 0;
    const video = useVideo ? videoPercentage : 0;
    const comment = useComment ? commentPercentage : 0;

    const total =
      openSession + documentMaterial + quiz + task + video + comment;

    return total;
  }, [
    openSessionPercentage,
    taskPercentage,
    videoPercentage,
    commentPercentage,
    quizPercentage,
    documentMaterialPercentage,
    utsPercentage,
    uasPercentage,
    useOpenSession,
    useDocumentMaterial,
    useComment,
    useQuiz,
    useUas,
    useUts,
    useVideo,
    useTask,
  ]);

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

  const onSubmit = async (payload: FormPresenceComponentSchemaType) => {
    try {
      const response = await createOrUpdatePresenceComponentStudent({
        academicPeriodId: academicPeriod as string,
        subjectId: subjectId as string,
        payload,
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
        message: `Pengaturan presensi berhasil disimpan`,
        state: "success",
      }));
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
    if (searchParams) {
      setValue("study_program_id", studyProgram as string);
    }
  }, [searchParams]);

  useEffect(() => {
    if (!useComment) {
      setValue("comment_percentage", 0);
    }

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
      <Row>
        <Col sm={12}>
          <Row>
            {/* input percentage */}
            <Col sm={12} lg={12}>
              <div className="d-flex flex-column gap-3 ">
                {studentPresences?.map((component) => (
                  <div
                    className={`d-flex border rounded-2 py-2 px-3 ${
                      totalPercentage > 100 ||
                      errors[component.key as keyof typeof errors]
                        ? "border-danger"
                        : ""
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
                        style={{ width: "80%" }}
                        className={`${
                          component.key === "uts_percentage" ||
                          component.key === "uas_percentage"
                            ? "d-none"
                            : "d-block"
                        }`}
                      >
                        <Progress value={watch(component.key as any)} />
                      </div>
                      <div className="d-flex flex-column">
                        <Controller
                          name={component.key as any}
                          control={control}
                          render={({ field }) => (
                            <input
                              type="text"
                              {...field}
                              className={`form-input p-2 rounded-3 ${
                                errors[component.key as keyof typeof errors]
                                  ? "border-danger"
                                  : ""
                              } ${
                                component.key === "uts_percentage" ||
                                component.key === "uas_percentage"
                                  ? "d-none"
                                  : "d-block"
                              }`}
                              style={{
                                border: errors[
                                  component.key as keyof typeof errors
                                ]
                                  ? "1px solid #dc3545"
                                  : "1px solid #DEE5EC",
                                width: "80px",
                              }}
                              onChange={(e) => {
                                const { numberValue } =
                                  handleInputNumberOnly(e);

                                return field.onChange(numberValue);
                              }}
                              tabIndex={1}
                              disabled={Boolean(
                                !watch(component.toggle as any)
                              )}
                            />
                          )}
                        />
                      </div>
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
                    {errors[component.key as keyof typeof errors] && (
                      <small className="text-danger mt-1">
                        {errors[component.key as keyof typeof errors]?.message}
                      </small>
                    )}
                  </div>
                ))}
              </div>
            </Col>
            {/* total percentage and the component */}
            <Col sm={12}>
              <div className="d-flex flex-column gap-3 rounded-3 border mt-3 py-2">
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
              </div>
              <div className="d-flex justify-content-end flex gap-2 mt-4">
                <Button
                  color="primary"
                  disabled={
                    isSubmitting ||
                    totalPercentage > 100 ||
                    totalPercentage !== 100
                  }
                  type="submit"
                >
                  {isSubmitting ? "Menyimpan..." : "Simpan"}
                </Button>
              </div>
            </Col>
          </Row>
        </Col>
      </Row>
    </form>
  );
};

export default FormPresenceComponent;
