"use client";

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
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { useModalContext } from "@/lib/hooks/use-modal";
import {
  FormClassAcademicPeriodDetailSchema,
  FormClassAcademicPeriodDetailType,
} from "@/lib/validations/settings/academic-period/form-class";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import { SelectComponent } from "@/components/ui/select";
import { classMerge } from "@/lib/utils/class-merge";
import { useSearchCurriculumYear } from "@/services/api/data-referensi/curriculum-year/use-get-search-curriculum-year";
import { useGetUnsiaStudyProgram } from "@/services/api/data-referensi/study-program/use-get-unsia-study-program";
import { useGetSearchSubject } from "@/services/api/settings/subject/use-get-search-subject";
import { createClass } from "@/services/api/settings/academic-period/class/create-class";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

export const ModalAddClass = ({
  academicPeriodId,
}: {
  academicPeriodId: string;
}) => {
  const { modalState, setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    reset,
    watch,
    setValue,
  } = useForm({
    resolver: zodResolver(FormClassAcademicPeriodDetailSchema),
    defaultValues: {
      code: "",
      name: "",
      capacity: 0,
      number_of_meeting: 0,
      academic_periode_id: {
        label: academicPeriodId,
        value: academicPeriodId,
      },
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

  const handleToggleModal = () => {
    setModalState((prev) => ({
      ...prev,
      open: !modalState.open,
    }));
    reset();
  };

  const onSubmit = async (data: FormClassAcademicPeriodDetailType) => {
    try {
      const response = await createClass(data);

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
        message: error.toString(),
      }));
    }
  };

  return (
    <Modal isOpen={modalState.open} centered>
      <div className="d-flex justify-content-between align-items-center pt-3 px-3 border-bottom pb-2">
        <p className="fs-4 fw-semibold mb-0 text-black">
          {modalState.state === "add"
            ? "Tambah Kelas"
            : modalState.state === "edit"
            ? "Ubah Kelas"
            : "Detail Kelas"}
        </p>
        <Button className="bg-white border-0 p-0" onClick={handleToggleModal}>
          <i className=" ri-close-fill text-black fs-4"></i>
        </Button>
      </div>
      <ModalBody>
        <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
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
                        isDisabled={!curriculumYearId || !studyProgramId}
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.subject_id} />
                </Col>
              </Row>
            </Col>
            {/* code */}
            <Col sm={12}>
              <Row
                className="align-items-center gap-2"
                style={{ paddingBottom: "10px" }}
              >
                <Col sm={12}>
                  <Label htmlFor="code" className="form-label mb-0 fw-medium">
                    Kode Kelas
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
                        placeholder="Masukkan Kode Kelas"
                        {...field}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.code} />
                </Col>
              </Row>
            </Col>
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
                        placeholder="Masukkan Jumlah Pertemuan"
                        {...field}
                        onChange={(e) => {
                          const { numberValue } = handleInputNumberOnly(e);

                          field.onChange(numberValue);
                        }}
                      />
                    )}
                  />
                  <FormErrorMessage errors={errors.number_of_meeting} />
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
