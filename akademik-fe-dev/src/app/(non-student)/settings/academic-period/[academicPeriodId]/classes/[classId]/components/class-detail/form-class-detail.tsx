"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { Controller, useForm } from "react-hook-form";
import { Button, Col, Input, Label, Row } from "reactstrap";

import { SaveIcon } from "@/components/icons/save";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { SelectComponent } from "@/components/ui/select";
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import {
  FormClassAcademicPeriodDetailSchema,
  FormClassAcademicPeriodDetailType,
} from "@/lib/validations/settings/academic-period/form-class";
import { useSearchCurriculumYear } from "@/services/api/data-referensi/curriculum-year/use-get-search-curriculum-year";
import { useGetUnsiaStudyProgram } from "@/services/api/data-referensi/study-program/use-get-unsia-study-program";
import { useGetSearchSubject } from "@/services/api/settings/subject/use-get-search-subject";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { updateClass } from "@/services/api/settings/academic-period/class/update-class";

export const FormClassDetail = ({
  data,
  classId,
  isDetail,
}: {
  data: ApiResponse<Class>;
  classId: string;
  isDetail?: boolean;
}) => {
  const { setModalConfirmationState } = useModalConfirmationContext();

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    watch,
    setValue,
  } = useForm({
    resolver: zodResolver(FormClassAcademicPeriodDetailSchema),
    defaultValues: {
      name: data?.data.name || "",
      capacity: data?.data?.capacity || 0,
      number_of_meeting: data?.data?.number_of_meeting || 0,
      program_study_id: {
        label: data.data.study_program_name,
        value: data?.data.study_program_id,
      },
      subject_id: {
        label: data.data.subject_name_en,
        value: data?.data.subject_id,
      },
      academic_periode_id: {
        label: data?.data?.academic_periode_id,
        value: data?.data?.academic_periode_id,
      },
      curriculum_year_id: {
        label: data?.data?.curriculum_year_name,
        value: data?.data?.curriculum_year_id,
      },
      code: data?.data?.code,
    },
  });

  const curriculumYearId = watch("curriculum_year_id")?.value;
  const studyProgramId = watch("program_study_id")?.value;

  const { data: curriculumYear, isLoading: isLoadingCurriculumYear } =
    useSearchCurriculumYear();

  const { data: studyProgram, isLoading: isLoadingStudyProgram } =
    useGetUnsiaStudyProgram();

  const { data: subjects, isLoading: isLoadingSubjects } = useGetSearchSubject({
    curriculum_year_id: curriculumYearId,
    study_program_id: studyProgramId,
  });

  // curriculum year options
  const curriculumYearOptions = curriculumYear?.data.map((item) => ({
    label: item.years,
    value: item.id,
  })) as OptionType[];

  // study program options
  const studyProgramOptions = studyProgram?.data?.map((item) => ({
    label: item.name,
    value: item.id,
  })) as OptionType[];

  // subject options
  const subjectOptions = subjects?.data?.map((item) => ({
    label: item.name_id,
    value: item.id,
  })) as OptionType[];

  const onSubmit = async (payload: FormClassAcademicPeriodDetailType) => {
    try {
      const response = await updateClass(classId, payload);

      if (response.error) {
        return setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "failed",
          message: response.message,
        }));
      }

      return setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        message: `Kelas berhasil di-ubah`,
        state: "success",
      }));
    } catch (error: any) {
      setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        state: "failed",
        message: error.toString(),
      }));
    }
  };
  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="d-flex flex-column gap-3"
      autoComplete="off"
    >
      <div className="d-flex justify-content-between align-items-center">
        <p className="fw-medium fs-5" style={{ color: "#3A3A3A" }}>
          Detail Kelas
        </p>
        {!isDetail ? (
          <Button color="success" disabled={isSubmitting}>
            <SaveIcon className="me-2" />
            Simpan
          </Button>
        ) : null}
      </div>
      <Row className="gap-1 gap-lg-0">
        <Col md={12} lg={6}>
          <Row className="gap-2">
            {/* Study Program */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="program_study_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Program Studi
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="program_study_id"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={studyProgramOptions as OptionType[]}
                        placeholder="Pilih Program Studi"
                        isError={!!errors.program_study_id}
                        id={"program_study_id"}
                        isLoading={isLoadingStudyProgram}
                        {...field}
                        onChange={(e) => {
                          setValue("subject_id", null);
                          field.onChange(e);
                        }}
                        isDisabled={isDetail}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.program_study_id} />
                </Col>
              </Row>
            </Col>
            {/* Curriculum Year */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="curriculum_year_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Tahun Kurikulum
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="curriculum_year_id"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={curriculumYearOptions as OptionType[]}
                        placeholder="Pilih Tahun Kurikulum"
                        isError={!!errors.curriculum_year_id}
                        id={"curriculum_year_id"}
                        isLoading={isLoadingCurriculumYear}
                        {...field}
                        onChange={(e) => {
                          setValue("subject_id", null);
                          field.onChange(e);
                        }}
                        isDisabled={isDetail}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.curriculum_year_id} />
                </Col>
              </Row>
            </Col>
            {/* Class Subjetct */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="subject_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Mata Kuliah
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="subject_id"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={subjectOptions as OptionType[]}
                        placeholder="Pilih Mata Kuliah"
                        isError={!!errors.subject_id}
                        isLoading={isLoadingSubjects}
                        id={"subject_id"}
                        isDisabled={
                          !curriculumYearId || !studyProgramId || isDetail
                        }
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.subject_id} />
                </Col>
              </Row>
            </Col>
          </Row>
        </Col>
        <Col md={12} lg={6}>
          <Row className="gap-2">
            {/* name */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label htmlFor="name" className="form-label mb-0 fw-medium">
                    Nama Kelas
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="name"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.name ? "border border-danger" : ""
                        }`}
                        id="name"
                        placeholder="Masukkan Nama Kelas"
                        {...field}
                        disabled={isDetail}
                        // onChange={(e) => {
                        //   const { stringValue } = handleInputNumberOnly(e);

                        //   field.onChange(stringValue);
                        // }}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.name} />
                </Col>
              </Row>
            </Col>
            {/* capacity */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="capacity"
                    className="form-label mb-0 fw-medium"
                  >
                    Kapasitas
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="capacity"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.capacity ? "border border-danger" : ""
                        }`}
                        id="capacity"
                        placeholder="Masukkan Kapasitas Kelas"
                        {...field}
                        onChange={(e) => {
                          const { numberValue } = handleInputNumberOnly(e);

                          field.onChange(numberValue);
                        }}
                        disabled={isDetail}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.capacity} />
                </Col>
              </Row>
            </Col>
            {/* number_of_meeting */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="number_of_meeting"
                    className="form-label mb-0 fw-medium"
                  >
                    Jumlah Pertemuan
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="number_of_meeting"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.number_of_meeting ? "border border-danger" : ""
                        }`}
                        id="number_of_meeting"
                        placeholder="Masukkan Jumlah Pertemuan Kelas"
                        {...field}
                        onChange={(e) => {
                          const { numberValue } = handleInputNumberOnly(e);

                          field.onChange(numberValue);
                        }}
                        disabled={isDetail}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.number_of_meeting} />
                </Col>
              </Row>
            </Col>
          </Row>
        </Col>
      </Row>
    </form>
  );
};
