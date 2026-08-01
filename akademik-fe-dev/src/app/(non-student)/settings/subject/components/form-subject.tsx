"use client";

import { Button, Col, Input, Label, Row } from "reactstrap";
import { useCallback, useEffect, useState } from "react";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useDebouncedCallback } from "use-debounce";

import { SelectComponent } from "@/components/ui/select";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import {
  FormSubjectSchema,
  FormSubjectSchemaType,
} from "@/lib/validations/settings/subject/form-subject";
import { useSearchCurriculumYear } from "@/services/api/data-referensi/curriculum-year/use-get-search-curriculum-year";
import { useSearchCourseType } from "@/services/api/data-referensi/course-type/use-get-search-course-type";
import { useSearchFieldOfStudy } from "@/services/api/data-referensi/field-of-study/use-get-search-field-of-study";
import { useSearchLecturer } from "@/services/api/sdm/biodata/use-get-search-lecture";
import { createSubject } from "@/services/api/settings/subject/create-subject";
import { useGetUnsiaStudyProgram } from "@/services/api/data-referensi/study-program/use-get-unsia-study-program";
import { useSearchCourseGroup } from "@/services/api/data-referensi/course-group/use-get-search-course-group";
import { editSubject } from "@/services/api/settings/subject/edit-subject";

type FormSubjectProps = {
  isEdit?: boolean;
  isDetail?: boolean;
  data?: Subject;
};

export const FormSubject = ({ isEdit, isDetail, data }: FormSubjectProps) => {
  const params = useParams();
  const router = useRouter();
  const { setModalConfirmationState } = useModalConfirmationContext();
  const [search, setSearch] = useState<string | null>(null);
  const [lecturerState, setLecturerState] = useState<
    | "supporting_lecturer_id"
    | "developer_rps_lecturer_id"
    | "subject_coordinator_lecturer_id"
    | ""
  >("");

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    setValue,
  } = useForm({
    resolver: zodResolver(FormSubjectSchema),
    defaultValues: {
      code: "",
      name_id: "",
      name_en: "",
      face_to_face_sks: 0,
      practicum_sks: 0,
      field_practice_sks: 0,
      simulation_sks: 0,
    },
  });

  const { data: curriculumYear, isLoading: isLoadingCurriculumYear } =
    useSearchCurriculumYear();

  const { data: courseType, isLoading: isLoadingCourseType } =
    useSearchCourseType();

  const { data: courseGroup, isLoading: isLoadingCourseGroup } =
    useSearchCourseGroup();

  const { data: programStudy, isLoading: isLoadingProgramStudy } =
    useGetUnsiaStudyProgram();

  const { data: fieldOfStudy, isLoading: isLoadingFieldOfStudy } =
    useSearchFieldOfStudy();

  const { data: lecturer, isLoading: isLoadingLecturer } = useSearchLecturer({
    page: 1,
    limit: 200,
    study_program_id: "",
    search: search,
  });

  // curriculum year options
  const curriculumYearOptions = curriculumYear?.data?.map((data) => ({
    label: data.years,
    value: data.id,
  }));

  // course type options
  const courseTypeOptions = courseType?.data?.map((data) => ({
    label: data.name,
    value: data.id,
  }));

  // course group options
  const courseGroupOptions = courseGroup?.data?.map((data) => ({
    label: data.name,
    value: data.id,
  }));

  // program study options
  const programStudyOptions = programStudy?.data?.map((data) => {
    return {
      label: data.name,
      value: data.id,
    };
  });

  // field Of Study options
  const fieldOfStudyptions = fieldOfStudy?.data?.map((data) => ({
    label: data.name,
    value: data.id,
  }));

  // field Of Study options
  const lecturerOptions = lecturer?.data?.data?.map((data) => ({
    label: data.name_of_user,
    value: data.id,
  }));

  const handleSearchLecturer = useDebouncedCallback(
    (
      value: string,
      state:
        | "supporting_lecturer_id"
        | "developer_rps_lecturer_id"
        | "subject_coordinator_lecturer_id"
    ) => {
      if (value) {
        setSearch(value);
        setLecturerState(state);
      } else {
        setSearch("");
        setLecturerState("");
      }
    },
    500
  );

  const handleSetFormValue = useCallback(() => {
    setValue("curriculum_year_id", {
      label: data?.curriculum_year_name as string,
      value: data?.curriculum_year_id as string,
    });
    setValue("code", data?.code as string);
    setValue("name_id", data?.name_id as string);
    setValue("name_en", data?.name_en as string);
    setValue("course_type_id", {
      label: data?.course_type_name as string,
      value: data?.course_type_id as string,
    });
    setValue("course_group_id", {
      label: data?.course_group_name as string,
      value: data?.course_group_id as string,
    });
    setValue("face_to_face_sks", data?.face_to_face_sks as number);
    setValue("practicum_sks", data?.practicum_sks as number);
    setValue("field_practice_sks", data?.field_practice_sks as number);
    setValue("simulation_sks", data?.simulation_sks as number);
    setValue("study_program_id", {
      label: data?.study_program_name as string,
      value: data?.study_program_id as string,
    });
    setValue("field_of_studies_id", {
      label: data?.field_study_name as string,
      value: data?.field_of_studies_id as string,
    });
    setValue(
      "supporting_lecturer_id",
      data?.supporting_lecturers?.map((lecture) => ({
        label: `${lecture.lecturer_front_title.replace("-", "")} ${
          lecture.lecturer_name
        }, ${lecture.lecturer_back_title}`,
        value: lecture.id,
      })) as OptionType[]
    );
    setValue(
      "developer_rps_lecturer_id",
      data?.developer_rps_lecturers?.map((lecture) => ({
        label: `${lecture.lecturer_front_title} ${lecture.lecturer_name}, ${lecture.lecturer_back_title}`,
        value: lecture.id,
      })) as OptionType[]
    );
    setValue(
      "subject_coordinator_lecturer_id",
      data?.subject_coordinator_lecturers?.map((lecture) => ({
        label: `${lecture.lecturer_front_title} ${lecture.lecturer_name}, ${lecture.lecturer_back_title}`,
        value: lecture.id,
      })) as OptionType[]
    );
    setValue("is_mku", data?.is_mku as boolean);
    setValue("is_sap", data?.is_sap as boolean);
    setValue("is_silabus", data?.is_silabus as boolean);
    setValue("is_teaching_material", data?.is_teaching_material as boolean);
    setValue("is_diktat", data?.is_diktat as boolean);
  }, [data, isDetail, isEdit]);

  const onSubmit = async (data: FormSubjectSchemaType) => {
    try {
      const response = isEdit
        ? await editSubject(params?.subjectId as string, data)
        : await createSubject(data);

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
        message: "Mata kuliah berhasil ditambah",
        state: "success",
      }));

      return router.push("/settings/subject");
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
    if (data) {
      handleSetFormValue();
    }
  }, [data]);

  return (
    <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
      <Row>
        <Col sm={12} lg={6}>
          <Row className="gap-3">
            {/* curriculum year */}
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
                        isLoading={isLoadingCurriculumYear}
                        placeholder="Pilih Tahun Kurikulum"
                        isError={!!errors.curriculum_year_id}
                        id={"curriculum_year_id"}
                        isDisabled={isDetail}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.curriculum_year_id} />
                </Col>
              </Row>
            </Col>
            {/* code subject */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label htmlFor="code" className="form-label mb-0 fw-medium">
                    Kode Mata Kuliah
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="code"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.code ? "border border-danger" : ""
                        }`}
                        id="code"
                        placeholder="Masukkan Kode Mata Kuliah"
                        disabled={isDetail}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.code} />
                </Col>
              </Row>
            </Col>
            {/* name_id subject */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="name_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Nama Mata Kuliah (IND)
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="name_id"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.name_id ? "border border-danger" : ""
                        }`}
                        id="name_id"
                        placeholder="Masukkan Nama Mata Kuliah (IND)"
                        disabled={isDetail}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.name_id} />
                </Col>
              </Row>
            </Col>
            {/* name_en subject */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="name_en"
                    className="form-label mb-0 fw-medium"
                  >
                    Nama Mata Kuliah (EN)
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="name_en"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.name_en ? "border border-danger" : ""
                        }`}
                        id="name_en"
                        placeholder="Masukkan Nama Mata Kuliah (EN)"
                        disabled={isDetail}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.name_en} />
                </Col>
              </Row>
            </Col>
            {/* course type */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="course_type_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Jenis Mata Kuliah
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="course_type_id"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={courseTypeOptions as OptionType[]}
                        isLoading={isLoadingCourseType}
                        placeholder="Pilih Jenis Mata Kuliah"
                        isError={!!errors.course_type_id}
                        id={"course_type_id"}
                        isDisabled={isDetail}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.course_type_id} />
                </Col>
              </Row>
            </Col>
            {/* course group */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="course_group_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Kelompok Mata Kuliah
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="course_group_id"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={courseGroupOptions as OptionType[]}
                        isLoading={isLoadingCourseGroup}
                        placeholder="Pilih Kelompok Mata Kuliah"
                        isError={!!errors.course_group_id}
                        id={"course_group_id"}
                        isDisabled={isDetail}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.course_group_id} />
                </Col>
              </Row>
            </Col>
            {/* face_to_face_sks*/}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="face_to_face_sks"
                    className="form-label mb-0 fw-medium"
                  >
                    SKS Tatap Muka
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="face_to_face_sks"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.face_to_face_sks ? "border border-danger" : ""
                        }`}
                        id="face_to_face_sks"
                        placeholder="Masukkan SKS Tatap Muka"
                        disabled={isDetail}
                        {...field}
                        onChange={(e) => {
                          const { numberValue } = handleInputNumberOnly(e);

                          field.onChange(numberValue);
                        }}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.face_to_face_sks} />
                </Col>
              </Row>
            </Col>
            {/* practicum_sks*/}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="practicum_sks"
                    className="form-label mb-0 fw-medium optional"
                  >
                    SKS Praktikum
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="practicum_sks"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.practicum_sks ? "border border-danger" : ""
                        }`}
                        id="practicum_sks"
                        placeholder="Masukkan SKS Praktikum"
                        disabled={isDetail}
                        {...field}
                        onChange={(e) => {
                          const { numberValue } = handleInputNumberOnly(e);

                          field.onChange(numberValue);
                        }}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.practicum_sks} />
                </Col>
              </Row>
            </Col>
            {/* field_practice_sks*/}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="field_practice_sks"
                    className="form-label mb-0 fw-medium optional"
                  >
                    SKS Praktik Lapangan
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="field_practice_sks"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.field_practice_sks
                            ? "border border-danger"
                            : ""
                        }`}
                        id="field_practice_sks"
                        placeholder="Masukkan SKS Praktik Lapangan"
                        disabled={isDetail}
                        {...field}
                        onChange={(e) => {
                          const { numberValue } = handleInputNumberOnly(e);

                          field.onChange(numberValue);
                        }}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.field_practice_sks} />
                </Col>
              </Row>
            </Col>
            {/* simulation_sks*/}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="simulation_sks"
                    className="form-label mb-0 fw-medium optional"
                  >
                    SKS Simulasi
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="simulation_sks"
                    control={control}
                    render={({ field }) => (
                      <Input
                        className={`form-control form-control-icon ${
                          errors.simulation_sks ? "border border-danger" : ""
                        }`}
                        id="simulation_sks"
                        placeholder="Masukkan SKS Simulasi"
                        disabled={isDetail}
                        {...field}
                        onChange={(e) => {
                          const { numberValue } = handleInputNumberOnly(e);

                          field.onChange(numberValue);
                        }}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.simulation_sks} />
                </Col>
              </Row>
            </Col>
          </Row>
        </Col>
        <Col sm={12} lg={6}>
          <Row className="gap-3">
            {/* program study */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="study_program_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Unit Pengampu
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
                            label: "Universitas Siber Asia",
                            value: "0",
                          },
                          ...(programStudy
                            ? (programStudyOptions as OptionType[])
                            : []),
                        ]}
                        isLoading={isLoadingProgramStudy}
                        placeholder="Pilih Unit Pengampu"
                        isError={!!errors.study_program_id}
                        id={"program_study_id"}
                        isDisabled={isDetail}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.study_program_id} />
                </Col>
              </Row>
            </Col>
            {/* field_of_studies */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="field_of_studies_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Rumpun Mata Kuliah
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="field_of_studies_id"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={fieldOfStudyptions as OptionType[]}
                        isLoading={isLoadingFieldOfStudy}
                        placeholder="Pilih Rumpun Mata Kuliah"
                        isError={!!errors.field_of_studies_id}
                        id={"field_of_studies_id"}
                        isDisabled={isDetail}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.field_of_studies_id} />
                </Col>
              </Row>
            </Col>
            {/* supporting_lecturer_id */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="supporting_lecturer_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Dosen Pengampu
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="supporting_lecturer_id"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={lecturerOptions as OptionType[]}
                        isLoading={isLoadingLecturer}
                        placeholder="Cari Dosen Pengampu"
                        isError={!!errors.supporting_lecturer_id}
                        id={"supporting_lecturer_id"}
                        isDisabled={isDetail}
                        {...field}
                        onInputChange={(e) =>
                          handleSearchLecturer(e, "supporting_lecturer_id")
                        }
                        menuIsOpen={
                          !!lecturerOptions?.length &&
                          lecturerState === "supporting_lecturer_id"
                        }
                        hideIndicator
                        isMulti
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.supporting_lecturer_id} />
                </Col>
              </Row>
            </Col>
            {/* developer_rps_lecturer_id */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="developer_rps_lecturer_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Dosen Pengembang RPS
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="developer_rps_lecturer_id"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={lecturerOptions as OptionType[]}
                        isLoading={isLoadingLecturer}
                        placeholder="Cari Dosen Pengembang RPS"
                        isError={!!errors.developer_rps_lecturer_id}
                        id={"developer_rps_lecturer_id"}
                        isDisabled={isDetail}
                        {...field}
                        onInputChange={(e) =>
                          handleSearchLecturer(e, "developer_rps_lecturer_id")
                        }
                        menuIsOpen={
                          !!lecturerOptions?.length &&
                          lecturerState === "developer_rps_lecturer_id"
                        }
                        hideIndicator
                        isMulti
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.developer_rps_lecturer_id} />
                </Col>
              </Row>
            </Col>
            {/* subject_coordinator_lecturer_id */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label
                    htmlFor="subject_coordinator_lecturer_id"
                    className="form-label mb-0 fw-medium"
                  >
                    Koord. Pengampu MK
                  </Label>
                </Col>
                <Col sm={12}>
                  <Controller
                    name="subject_coordinator_lecturer_id"
                    control={control}
                    render={({ field }) => (
                      <SelectComponent
                        options={lecturerOptions as OptionType[]}
                        isLoading={isLoadingLecturer}
                        placeholder="Cari Koord. Pengampu MK"
                        isError={!!errors.subject_coordinator_lecturer_id}
                        id={"subject_coordinator_lecturer_id"}
                        isDisabled={isDetail}
                        {...field}
                        onInputChange={(e) =>
                          handleSearchLecturer(
                            e,
                            "subject_coordinator_lecturer_id"
                          )
                        }
                        menuIsOpen={
                          !!lecturerOptions?.length &&
                          lecturerState === "subject_coordinator_lecturer_id"
                        }
                        hideIndicator
                        isMulti
                      />
                    )}
                  />
                  <FormErrorMessage
                    errors={errors.subject_coordinator_lecturer_id}
                  />
                </Col>
              </Row>
            </Col>
            {/* switch toggle */}
            <Col sm={12}>
              <Row className="gap-3">
                {/* is mku */}
                <Col sm={5}>
                  <label className="form-check-label" htmlFor="is_mku">
                    Apakah Termasuk MKU
                  </label>
                  <div className="form-check form-switch mt-2">
                    <Controller
                      name="is_mku"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className="form-check-input"
                          type="checkbox"
                          role="switch"
                          id="is_mku"
                          style={{ width: "40px", height: "20px" }}
                          {...field}
                          value={field.value ? "true" : "false"}
                          checked={field.value === true}
                          disabled={isDetail}
                        />
                      )}
                    />
                  </div>
                </Col>
                {/* is sap */}
                <Col sm={5}>
                  <label className="form-check-label" htmlFor="is_sap">
                    SAP
                  </label>
                  <div className="form-check form-switch mt-2">
                    <Controller
                      name="is_sap"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className="form-check-input"
                          type="checkbox"
                          role="switch"
                          id="is_sap"
                          style={{ width: "40px", height: "20px" }}
                          {...field}
                          value={field.value ? "true" : "false"}
                          checked={field.value === true}
                          disabled={isDetail}
                        />
                      )}
                    />
                  </div>
                </Col>
                {/* is silabus */}
                <Col sm={5}>
                  <label className="form-check-label" htmlFor="is_silabus">
                    Silabus
                  </label>
                  <div className="form-check form-switch mt-2">
                    <Controller
                      name="is_silabus"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className="form-check-input"
                          type="checkbox"
                          role="switch"
                          id="is_silabus"
                          style={{ width: "40px", height: "20px" }}
                          {...field}
                          value={field.value ? "true" : "false"}
                          checked={field.value === true}
                          disabled={isDetail}
                        />
                      )}
                    />
                  </div>
                </Col>
                {/* is material */}
                <Col sm={5}>
                  <label
                    className="form-check-label"
                    htmlFor="is_teaching_material"
                  >
                    Bahan Ajar
                  </label>
                  <div className="form-check form-switch mt-2">
                    <Controller
                      name="is_teaching_material"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className="form-check-input"
                          type="checkbox"
                          role="switch"
                          id="is_teaching_material"
                          style={{ width: "40px", height: "20px" }}
                          {...field}
                          value={field.value ? "true" : "false"}
                          checked={field.value === true}
                          disabled={isDetail}
                        />
                      )}
                    />
                  </div>
                </Col>
                {/* is dictat */}
                <Col sm={5}>
                  <label className="form-check-label" htmlFor="is_diktat">
                    Diktat
                  </label>
                  <div className="form-check form-switch mt-2">
                    <Controller
                      name="is_diktat"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className="form-check-input"
                          type="checkbox"
                          role="switch"
                          id="is_diktat"
                          style={{ width: "40px", height: "20px" }}
                          {...field}
                          value={field.value ? "true" : "false"}
                          checked={field.value === true}
                          disabled={isDetail}
                        />
                      )}
                    />
                  </div>
                </Col>
              </Row>
            </Col>
          </Row>
        </Col>
      </Row>
      <div className="d-flex justify-content-between mt-3">
        <Link href={"/settings/subject"}>
          <button
            type="button"
            className={
              "waves-effect waves-light btn-outline text-primary px-3 py-2"
            }
            disabled={isSubmitting}
          >
            Kembali
          </button>
        </Link>
        {!isDetail ? (
          <Button disabled={isSubmitting} color="primary">
            {isEdit ? "Ubah" : "Simpan"}
          </Button>
        ) : null}
      </div>
    </form>
  );
};
