"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useQueryClient } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import { Controller, useForm } from "react-hook-form";
import { Button, Col, Input, Label, Modal, ModalBody, Row } from "reactstrap";

import { FormErrorMessage } from "@/components/ui/form-error-message";
import { SelectComponent } from "@/components/ui/select";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { curriculumStudyProgramForAcademicSchema } from "@/lib/validations/curriculum/form-curriculum-study-program";
import { createCurriculumStudyProgramForProgramHead } from "@/services/api/curriculum/curriculum-study-program/program-head/create-curriculum-study-program-for-program-head";
import { useSearchCurriculumYear } from "@/services/api/data-referensi/curriculum-year/use-get-search-curriculum-year";
import { useGetSearchGrade } from "@/services/api/data-referensi/grade/use-get-search-grade";
import { useGetSearchSemesterNumber } from "@/services/api/data-referensi/semester-number/use-get-search-semester-number";
import { useGetSearchSubjectProgramHead } from "@/services/api/settings/subject/use-get-search-subject-program-head";
import { useGetSearchSubject } from "@/services/api/settings/subject/use-get-search-subject";
import { AKADEMIK, KAPRODI } from "@/lib/constants/role";
import { createCurriculumStudyProgramForAcademic } from "@/services/api/curriculum/curriculum-study-program/academic/create-curriculum-study-program-for-academic";
import { useModalContext } from "@/lib/hooks/use-modal";
import { AddIcon } from "@/components/icons/add";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import { useGetSearchSubjectPrerequisitesForAcademic } from "@/services/api/curriculum/curriculum-study-program/academic/use-get-search-subject-prerequisites-for-academic";
import { useSearchFielStudyConcentration } from "@/services/api/data-referensi/field-study-concentration/use-get-search-field-study-concentration";
import { DeleteCurriculumStudyProgram } from "./delete-curriculum-study-program-confirmation";
import { useGetSearchSubjectPrerequisitesForProgramHead } from "@/services/api/curriculum/curriculum-study-program/program-head/use-get-search-subject-prerequisites-for-program-head";
import { useGetDetailCurriculumStudyProgramForAcademic } from "@/services/api/curriculum/curriculum-study-program/academic/use-get-detail-curriculum-study-program-for-academic";
import { useGetDetailCurriculumStudyProgramForProgramHead } from "@/services/api/curriculum/curriculum-study-program/program-head/use-get-detail-curriculum-study-program-for-program-head";
import { updateCurriculumStudyProgramForAcademic } from "@/services/api/curriculum/curriculum-study-program/academic/update-curriculum-study-program-for-academic";
import { updateCurriculumStudyProgramForProgramHead } from "@/services/api/curriculum/curriculum-study-program/program-head/update-curriculum-study-program-for-program-head";

export const FormCurriculumStudyProgram = ({
  studyProgramId,
  role,
}: {
  studyProgramId: string;
  role: string;
}) => {
  const [curriculumState, setCurriculumState] = useState<
    OptionType | undefined
  >(undefined);
  const pathname = usePathname();
  const router = useRouter();
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);

  const { modalState, setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();
  const {
    control,
    formState: { errors, isSubmitting },
    handleSubmit,
    watch,
    reset,
    setValue,
  } = useForm({
    resolver: zodResolver(curriculumStudyProgramForAcademicSchema),
    defaultValues: {
      curriculum_year_id: "",
      semester_number_id: "",
      limit_grade_id: "",
      is_mandatory: false,
      study_program_id: studyProgramId,
    },
  });

  const queryClient = useQueryClient();

  const curriculumYearParam = searchParams.get("curriculum_year");
  const curriculumYearId = watch("curriculum_year_id");
  const semesterNumberId = watch("semester_number_id");

  const { data: curriculumYear, isLoading: isLoadingCurriculumYear } =
    useSearchCurriculumYear();

  // curriculum year options
  const curriculumYearOptions = curriculumYear?.data?.map((curriculum) => ({
    label: curriculum.years,
    value: curriculum.id,
  }));

  // get semester
  const {
    data: semester,
    refetch: refetchSemester,
    isLoading: isLoadingSemester,
  } = useGetSearchSemesterNumber({ page: 1 });

  // get grade
  const {
    data: grade,
    refetch: refetchGrade,
    isLoading: isLoadingGrade,
  } = useGetSearchGrade({ page: 1 });

  // get subject program head
  const {
    data: subjectsProgramHead,
    isLoading: isLoadingSubjectsProgramHead,
    refetch: refetchSubjectProgramHead,
  } = useGetSearchSubjectProgramHead({
    curriculum_year_id: curriculumYearId,
  });

  // get subject academic
  const {
    data: subjects,
    isLoading: isLoadingSubject,
    refetch: refetchSubject,
  } = useGetSearchSubject({
    curriculum_year_id: curriculumYearId,
    study_program_id: studyProgramId,
  });

  // subject options based on role
  const subjectOptions =
    role === AKADEMIK
      ? subjects?.data?.map((item) => ({
          label: item.name_id,
          value: item.id,
        }))
      : subjectsProgramHead?.data?.map((item) => ({
          label: item.name_id,
          value: item.id,
        }));

  // get subject prerequisites academic
  const {
    data: subjectPrerequisites,
    isLoading: isLoadingSubjectPrerequisites,
    refetch: refetchSubjectPrerequisites,
  } = useGetSearchSubjectPrerequisitesForAcademic({
    curriculumYearId,
    semesterNumberId,
    studyProgramId,
  });

  // get subject prerequisites program head
  const {
    data: subjectPrerequisitesProgramHead,
    isLoading: isLoadingSubjectPrerequisitesProgramHead,
    refetch: refetchSubjectPrerequisitesProgramHead,
  } = useGetSearchSubjectPrerequisitesForProgramHead({
    curriculumYearId,
    semesterNumberId,
  });

  // get data by id for academic
  const {
    data: curriculumStudyProgramAcademic,
    isLoading: isLoadingCurriculumStudyProgramAcademic,
  } = useGetDetailCurriculumStudyProgramForAcademic({
    curriculumStudyProgramId: modalState.id as string,
    role,
  });

  // get data by id for program head
  const {
    data: curriculumStudyProgramProgramHead,
    isLoading: isLoadingCurriculumStudyProgramHead,
  } = useGetDetailCurriculumStudyProgramForProgramHead({
    curriculumStudyProgramId: modalState.id as string,
    role,
  });

  // subject prerequisites options condition based on role
  const subjectPrerequisitesOptions =
    role === AKADEMIK
      ? subjectPrerequisites?.data.map((subject) => ({
          label: subject.subject_name_id,
          value: subject.id,
        }))
      : subjectPrerequisitesProgramHead?.data?.map((subject) => ({
          label: subject.subject_name_id,
          value: subject.id,
        }));

  // field study concentration
  const {
    data: fieldStudyConcentration,
    isLoading: isLoadingFieldStudyConcentration,
    refetch: refetchFieldStudyConcentration,
  } = useSearchFielStudyConcentration();

  // field study concentration options
  const fieldStudyConcentrationOtions = fieldStudyConcentration?.data?.map(
    (data) => ({
      label: `${data.code} - ${data.name}`,
      value: data.id,
    })
  );

  // handle curriculum year filter
  const handleCurriculumYearFilter = (option: OptionType) => {
    if (option) {
      params.set("curriculum_year", option.value);
    } else {
      params.delete("curriculum_year");
    }

    router.replace(`${pathname}?${params.toString()}`);
  };

  // on submit based on role
  const onSubmitCreate = async (payload: any) => {
    try {
      // check role session
      const response =
        role === AKADEMIK
          ? await createCurriculumStudyProgramForAcademic({
              ...payload,
              study_program_id: studyProgramId,
            }) // this for academic
          : await createCurriculumStudyProgramForProgramHead(payload); //this for program head aka kaprodi

      if (response.error) {
        return setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "failed",
          message: response.message,
        }));
      }

      // refetch data when success
      queryClient.refetchQueries({
        queryKey: [
          "curriculum-study-program-for-academic",
          semesterNumberId,
          studyProgramId,
          curriculumYearId,
          AKADEMIK,
        ],
      });

      // refetch data when success
      queryClient.refetchQueries({
        queryKey: [
          "curriculum-study-program-for-program-head",
          semesterNumberId,
          curriculumYearId,
        ],
      });

      setModalState((prev) => ({
        ...prev,
        open: false,
        id: undefined,
      }));

      return setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        message: `Data berhasil ditambah`,
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

  // on submit edit based on role
  const onSubmitEdit = async (payload: any) => {
    try {
      // check role session
      const response =
        role === AKADEMIK
          ? await updateCurriculumStudyProgramForAcademic({
              id: modalState.id as string,
              payload: {
                ...payload,
                study_program_id: studyProgramId,
              },
            }) // this for academic
          : await updateCurriculumStudyProgramForProgramHead({
              id: modalState.id as string,
              payload,
            }); //this for program head aka kaprodi

      if (response.error) {
        return setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "failed",
          message: response.message,
        }));
      }

      // refetch data when success
      queryClient.refetchQueries({
        queryKey: [
          "curriculum-study-program-for-academic",
          semesterNumberId,
          studyProgramId,
          curriculumYearId,
          AKADEMIK,
        ],
      });

      // refetch data when success
      queryClient.refetchQueries({
        queryKey: [
          "curriculum-study-program-for-program-head",
          semesterNumberId,
          curriculumYearId,
        ],
      });

      queryClient.resetQueries({
        queryKey: ["detail-curriculum-study-program", modalState.id],
      });

      setModalState((prev) => ({
        ...prev,
        open: false,
        id: undefined,
      }));

      return setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        message: `Data berhasil diubah`,
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

  useEffect(() => {
    // search not work
    if (curriculumYearId && role === KAPRODI) {
      refetchSubjectProgramHead();
    }
  }, [curriculumYearId, role]);

  useEffect(() => {
    if (curriculumYearId && role === AKADEMIK && studyProgramId) {
      refetchSubject();
    }
  }, [curriculumYearId, role, studyProgramId]);

  useEffect(() => {
    reset();
  }, [studyProgramId]);

  // set value for academic base on data by id
  useEffect(() => {
    if (curriculumYearParam && modalState.state === "add") {
      reset();

      setValue("curriculum_year_id", curriculumYearParam);
    } else if (
      modalState.open &&
      curriculumStudyProgramAcademic &&
      modalState.state === "edit"
    ) {
      setValue(
        "curriculum_year_id",
        curriculumStudyProgramAcademic?.data.curriculum_year_id
      );
      setValue("subject_id", {
        label: curriculumStudyProgramAcademic?.data?.subject_name_id,
        value: curriculumStudyProgramAcademic?.data?.subject_id,
      });
      setValue(
        "limit_grade_id",
        curriculumStudyProgramAcademic?.data?.limit_grade_id
      );
      if (curriculumStudyProgramAcademic?.data.subject_prerequisites.length) {
        // const findSubjectPrerequisite = subjectPrerequisitesOptions?.find(
        //   (opt) =>
        //     curriculumStudyProgramAcademic?.data.subject_prerequisites?.some(
        //       (p) => p.subject_id === opt.value
        //     )
        // );

        // console.log(findSubjectPrerequisite);

        setValue(
          "subject_prerequisites",
          curriculumStudyProgramAcademic?.data?.subject_prerequisites.map(
            (data) => ({
              label: data.subject_name_id,
              value: data.id,
            })
          )
        );
      } else {
        setValue("subject_prerequisites", undefined);
      }
      if (curriculumStudyProgramAcademic.data.field_study_concentration_id)
        setValue("field_study_concentration_id", {
          label:
            curriculumStudyProgramAcademic?.data
              ?.field_study_concentration_name,
          value:
            curriculumStudyProgramAcademic.data.field_study_concentration_id,
        });
      setValue(
        "semester_number_id",
        curriculumStudyProgramAcademic.data.semester_number_id
      );
      setValue(
        "is_mandatory",
        curriculumStudyProgramAcademic.data.is_mandatory
      );
    }
  }, [
    searchParams,
    modalState.open,
    modalState.state,
    curriculumStudyProgramAcademic,
    setValue,
  ]);

  // set value for program head base on data by id
  useEffect(() => {
    if (curriculumYearParam && modalState.state === "add") {
      setValue("curriculum_year_id", curriculumYearParam);
    } else if (
      modalState.open &&
      curriculumStudyProgramProgramHead &&
      modalState.state === "edit"
    ) {
      setValue(
        "curriculum_year_id",
        curriculumStudyProgramProgramHead?.data.curriculum_year_id
      );
      setValue("subject_id", {
        label: curriculumStudyProgramProgramHead?.data?.subject_name_id,
        value: curriculumStudyProgramProgramHead?.data?.subject_id,
      });
      setValue(
        "limit_grade_id",
        curriculumStudyProgramProgramHead?.data?.limit_grade_id
      );
      if (
        curriculumStudyProgramProgramHead?.data.subject_prerequisites.length
      ) {
        setValue(
          "subject_prerequisites",
          curriculumStudyProgramProgramHead?.data?.subject_prerequisites.map(
            (data) => ({
              label: data.subject_name_id,
              value: data.id,
            })
          )
        );
      } else {
        setValue("subject_prerequisites", undefined);
      }
      if (curriculumStudyProgramProgramHead.data.field_study_concentration_id)
        setValue("field_study_concentration_id", {
          label:
            curriculumStudyProgramProgramHead?.data
              ?.field_study_concentration_name,
          value:
            curriculumStudyProgramProgramHead.data.field_study_concentration_id,
        });
      setValue(
        "semester_number_id",
        curriculumStudyProgramProgramHead.data.semester_number_id
      );
      setValue(
        "is_mandatory",
        curriculumStudyProgramProgramHead.data.is_mandatory
      );
    }
  }, [
    searchParams,
    modalState.open,
    modalState.state,
    curriculumStudyProgramProgramHead,
    setValue,
  ]);

  // initial fetch for options
  useEffect(() => {
    refetchSemester();
    refetchGrade();
    refetchFieldStudyConcentration();
  }, []);

  // set curriculum year filter based on param
  useEffect(() => {
    if (searchParams && curriculumYearParam) {
      const findCurriculum = curriculumYearOptions?.find(
        (data) => data.value === curriculumYearParam
      );
      setCurriculumState(findCurriculum);
    } else {
      setCurriculumState(undefined);
    }
  }, [searchParams, curriculumYear]);

  // refetch subject prerequisites
  useEffect(() => {
    if (curriculumYearParam && semesterNumberId) {
      if (role === AKADEMIK) {
        refetchSubjectPrerequisites();
      } else {
        refetchSubjectPrerequisitesProgramHead();
      }
    }
  }, [role, curriculumYearParam, semesterNumberId]);

  return (
    <>
      <DeleteCurriculumStudyProgram role={role} />
      <div className="row justify-content-between align-items-center">
        <Col sm={3}>
          <SelectComponent
            id={"curriculum"}
            options={curriculumYearOptions as OptionType[]}
            isClearable
            onChange={(value) => handleCurriculumYearFilter(value)}
            placeholder="Filter Tahun Kurikulum"
            value={curriculumState}
          />
        </Col>
        <Col sm={1} className="w-auto">
          {curriculumYearParam ? (
            <button
              className="btn-outline text-primary py-2 px-3"
              onClick={() => {
                setModalState((prev) => ({
                  ...prev,
                  open: true,
                  state: "add",
                }));
                reset();
              }}
            >
              <AddIcon />
              Tambah
            </button>
          ) : null}
        </Col>
      </div>
      <Modal isOpen={modalState.open} centered>
        <div className="d-flex justify-content-between align-items-center pt-3 px-3 border-bottom pb-2">
          <p className="fs-4 fw-semibold mb-0 text-black">Kurikulum Prodi</p>
        </div>
        <ModalBody>
          <form
            onSubmit={handleSubmit(
              modalState.state === "add" ? onSubmitCreate : onSubmitEdit
            )}
            autoComplete="off"
            autoCapitalize="none"
          >
            <Row className="mb-3 gap-3">
              {/* curriculum_year_id */}
              <Col sm={12}>
                <Label
                  htmlFor="curriculum_year_id"
                  className="form-label fw-medium"
                >
                  Tahun Kurikulum
                </Label>
                <Controller
                  control={control}
                  name="curriculum_year_id"
                  render={({ field }) => (
                    <select
                      {...field}
                      className={`form-select ${
                        errors.curriculum_year_id ? "border border-danger" : ""
                      }`}
                      id="curriculum_year_id"
                      disabled={
                        isLoadingCurriculumYear ||
                        isLoadingCurriculumStudyProgramAcademic ||
                        isLoadingCurriculumStudyProgramHead
                      }
                    >
                      <option value="" disabled>
                        Kurikulum
                      </option>
                      {curriculumYear?.data?.map((item) => (
                        <option key={item.id} value={item.id}>
                          {item.years}
                        </option>
                      ))}
                    </select>
                  )}
                />
                <FormErrorMessage errors={errors.curriculum_year_id} />
              </Col>
              {/* semester_number_id */}
              <Col sm={12}>
                <Label
                  htmlFor="semester_number_id"
                  className="form-label fw-medium"
                >
                  Semester
                </Label>
                <Controller
                  control={control}
                  name="semester_number_id"
                  render={({ field }) => (
                    <select
                      {...field}
                      className={`form-select ${
                        errors.semester_number_id ? "border border-danger" : ""
                      }`}
                      id="semester_number_id"
                      disabled={
                        isLoadingSemester ||
                        isLoadingCurriculumStudyProgramAcademic ||
                        isLoadingCurriculumStudyProgramHead
                      }
                    >
                      <option value="" disabled>
                        Semester
                      </option>
                      {semester?.data?.map((item) => (
                        <option key={item.id} value={item.id}>
                          {item.semester_number}
                        </option>
                      ))}
                    </select>
                  )}
                />
                <FormErrorMessage errors={errors.semester_number_id} />
              </Col>
              {/* subject_id */}
              <Col sm={12}>
                <Label htmlFor="subject_id" className="form-label fw-medium">
                  Mata Kuliah
                </Label>
                <Controller
                  name="subject_id"
                  control={control}
                  render={({ field }) => (
                    <SelectComponent
                      options={subjectOptions as OptionType[]}
                      isLoading={
                        isLoadingSubjectsProgramHead ||
                        isLoadingSubject ||
                        isLoadingCurriculumStudyProgramAcademic ||
                        isLoadingCurriculumStudyProgramHead
                      }
                      placeholder="Pilih Mata Kuliah"
                      isError={!!errors.subject_id}
                      id={"subject_id"}
                      {...field}
                      isDisabled={
                        !curriculumYearId ||
                        isLoadingCurriculumStudyProgramAcademic ||
                        isLoadingCurriculumStudyProgramHead
                      }
                    />
                  )}
                />
                <FormErrorMessage errors={errors.subject_id} />
              </Col>
              {/* limit_grade_id */}
              <Col sm={12}>
                <Label
                  htmlFor="limit_grade_id"
                  className="form-label fw-medium"
                >
                  Min Nilai
                </Label>
                <Controller
                  control={control}
                  name="limit_grade_id"
                  render={({ field }) => (
                    <select
                      {...field}
                      className={`form-select ${
                        errors.limit_grade_id ? "border border-danger" : ""
                      }`}
                      id="limit_grade_id"
                      disabled={
                        isLoadingGrade ||
                        isLoadingCurriculumStudyProgramAcademic ||
                        isLoadingCurriculumStudyProgramHead
                      }
                    >
                      <option value="" disabled>
                        Min Nilai
                      </option>
                      {grade?.data?.map((item) => (
                        <option key={item.id} value={item.id}>
                          {item.code} - {item.name}
                        </option>
                      ))}
                    </select>
                  )}
                />
                <FormErrorMessage errors={errors.limit_grade_id} />
              </Col>
              {/* subject_prerequisites */}
              <Col sm={12}>
                <Label
                  htmlFor="subject_prerequisites"
                  className="form-label fw-medium optional"
                >
                  Prasyarat
                </Label>
                <Controller
                  name="subject_prerequisites"
                  control={control}
                  render={({ field }) => (
                    <SelectComponent
                      options={subjectPrerequisitesOptions as OptionType[]}
                      isLoading={
                        isLoadingSubjectPrerequisites ||
                        isLoadingSubjectPrerequisitesProgramHead ||
                        isLoadingCurriculumStudyProgramAcademic ||
                        isLoadingCurriculumStudyProgramHead
                      }
                      placeholder="Pilih Mata Kuliah Prasyarat"
                      isError={!!errors.subject_prerequisites}
                      id={"subject_prerequisites"}
                      {...field}
                      isDisabled={
                        !curriculumYearId ||
                        !semesterNumberId ||
                        isLoadingSubjectPrerequisites ||
                        isLoadingSubjectPrerequisitesProgramHead ||
                        isLoadingCurriculumStudyProgramAcademic ||
                        isLoadingCurriculumStudyProgramHead
                      }
                      isMulti
                    />
                  )}
                />
                <FormErrorMessage errors={errors.subject_prerequisites} />
              </Col>
              {/* field_study_concentration_id */}
              <Col sm={12}>
                <Label
                  htmlFor="field_study_concentration_id"
                  className="form-label fw-medium optional"
                >
                  Konsentrasi Bidang Studi
                </Label>
                <Controller
                  name="field_study_concentration_id"
                  control={control}
                  render={({ field }) => (
                    <SelectComponent
                      options={fieldStudyConcentrationOtions as OptionType[]}
                      isLoading={
                        isLoadingFieldStudyConcentration ||
                        isLoadingCurriculumStudyProgramAcademic ||
                        isLoadingCurriculumStudyProgramHead
                      }
                      placeholder="Pilih Konsentrasi Bidang Studi"
                      isError={!!errors.field_study_concentration_id}
                      id={"field_study_concentration_id"}
                      {...field}
                      isDisabled={
                        isLoadingFieldStudyConcentration ||
                        isLoadingCurriculumStudyProgramAcademic ||
                        isLoadingCurriculumStudyProgramHead
                      }
                    />
                  )}
                />
                <FormErrorMessage
                  errors={errors.field_study_concentration_id}
                />
              </Col>
              <Col sm={12} className="w-auto">
                <div className="d-flex gap-2">
                  <div className="d-flex flex-column">
                    <label
                      className="form-check-label mb-1"
                      style={{ color: "#3A3A3A", fontWeight: 300 }}
                    >
                      Pilihan Tambahan
                    </label>
                    <div className="form-check">
                      <Controller
                        name="is_mandatory"
                        control={control}
                        render={({ field }) => (
                          <Input
                            className="form-check-input"
                            type="checkbox"
                            id="is_mandatory"
                            {...field}
                            value={field.value ? "true" : "false"}
                            checked={field.value === true}
                            disabled={
                              isLoadingCurriculumStudyProgramAcademic ||
                              isLoadingCurriculumStudyProgramHead
                            }
                          />
                        )}
                      />
                      <label
                        className="form-check-label"
                        htmlFor="is_mandatory"
                      >
                        MK Wajib
                      </label>
                    </div>
                  </div>
                </div>
              </Col>
            </Row>
            <div className="d-flex justify-content-end mt-3 gap-2">
              <button
                className="bg-transparent text-primary rounded px-3"
                type="button"
                style={{ border: "1px solid #10487A" }}
                disabled={isSubmitting}
                onClick={() =>
                  setModalState((prev) => ({
                    ...prev,
                    open: false,
                  }))
                }
              >
                <span>Batal</span>
              </button>
              <Button
                color="primary"
                className="d-flex flex-grow-0 justify-content-center align-items-center"
                disabled={isSubmitting}
              >
                <span>
                  {modalState.state === "add" ? "Tambah" : null}
                  {modalState.state === "edit" ? "Ubah" : null}
                </span>
              </Button>
            </div>
          </form>
        </ModalBody>
      </Modal>
    </>
  );
};
