"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect } from "react";
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

import { CalendarTodayIcon } from "@/components/icons/calendar-today";
import { DatePicker } from "@/components/ui/date-picker";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { SelectComponent } from "@/components/ui/select";
import { useModalContext } from "@/lib/hooks/use-modal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { classMerge } from "@/lib/utils/class-merge";
import {
  AcademicPeriodFormType,
  AcademicPeriodInitValues,
  academicPeriodSchema,
} from "@/lib/validations/academic/settings/academic-period";
import { createAcademicPeriod } from "@/services/api/data-referensi/academic-period/create-academic-period";
import { updateAcademicPeriod } from "@/services/api/data-referensi/academic-period/update-academic-period";
import { useGetAcademicPeriodById } from "@/services/api/data-referensi/academic-period/use-get-academic-period-by-id";
import { useAcademicYears } from "@/services/api/data-referensi/academic-year/use-get-academic-year";
import { useSemesters } from "@/services/api/data-referensi/semester/use-get-academic-year";

const ModalAcademicPeriod = () => {
  const { modalState, setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    reset,
    watch,
    setValue,
  } = useForm<AcademicPeriodFormType>({
    resolver: zodResolver(
      modalState.state === "add"
        ? academicPeriodSchema({ isEdit: false })
        : academicPeriodSchema({ isEdit: true })
    ),
    defaultValues: AcademicPeriodInitValues,
  });

  const academicYear = watch("academic_year_id");
  const semester = watch("semester_id");
  const startCollege = watch("start_date_of_college");
  const startUas = watch("start_date_of_uas");
  const startUts = watch("start_date_of_uts");

  const { data: academicYears, isLoading: isLoadingAcademicYears } =
    useAcademicYears();

  const { data: academicPeriod, isLoading: isLoadingAcademicPeriod } =
    useGetAcademicPeriodById(modalState?.id as string);

  const academicYearOptions = academicYears?.data?.map((val: AcademicYear) => ({
    label: val.name,
    value: val.id,
  }));

  const { data: semesters, isLoading: isLoadingSemesters } = useSemesters();

  const semesterOptions = semesters?.data?.map((val: Semester) => ({
    label: val.name,
    value: val.id,
  }));

  const handleToggleModal = () => {
    setModalState((prev) => ({
      ...prev,
      open: !prev.open,
      id: undefined,
    }));
  };

  const onSubmit = async (data: AcademicPeriodFormType) => {
    try {
      const response =
        modalState.state === "add"
          ? await createAcademicPeriod(data)
          : await updateAcademicPeriod(modalState.id as string, data);

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

  const handleSetFormValue = () => {
    if (academicPeriod) {
      setValue("fullname", academicPeriod?.data?.fullname);
      setValue("code", academicPeriod?.data?.code);
      setValue("shortname", academicPeriod?.data?.shortname);
      setValue("is_active", academicPeriod?.data?.is_active);
      setValue("start_date_of_college", [
        new Date(academicPeriod?.data?.start_date_of_college),
      ]);
      setValue("end_date_of_college", [
        new Date(academicPeriod?.data?.end_date_of_college),
      ]);
      setValue("start_date_of_uas", [
        new Date(academicPeriod?.data?.start_date_of_uas),
      ]);
      setValue("end_date_of_uas", [
        new Date(academicPeriod?.data?.end_date_of_uas),
      ]);
      setValue("start_date_of_uts", [
        new Date(academicPeriod?.data?.start_date_of_uts),
      ]);
      setValue("end_date_of_uts", [
        new Date(academicPeriod?.data?.end_date_of_uts),
      ]);
      setValue("academic_year_id", {
        label: academicPeriod?.data?.academic_year,
        value: academicPeriod?.data?.academic_year_id,
      });
      setValue("semester_id", {
        label: academicPeriod?.data?.semester,
        value: academicPeriod?.data?.semester_id,
      });
      setValue("number_of_lecture_meeting", {
        label: academicPeriod?.data?.number_of_lecture_meeting,
        value: academicPeriod?.data?.number_of_lecture_meeting,
      });
      setValue("is_active", academicPeriod?.data?.is_active);
    }
  };

  useEffect(() => {
    if (academicYear && semester) {
      const findAcademicYear = academicYears?.data?.find(
        (data) => data.id === academicYear.value
      );

      const convertSemesterToNumber = (semester: string) => {
        switch (semester.toLocaleLowerCase()) {
          case "ganjil":
            return 1;
          case "genap":
            return 2;
          case "antara":
            return 3;
          default:
            return 0;
        }
      };

      setValue(
        "code",
        `${findAcademicYear?.years}${convertSemesterToNumber(semester.label)} `
      );
      setValue("fullname", `${semester.label} ${academicYear.label}`);
      setValue("shortname", `${semester.label} ${academicYear.label}`);
    }
  }, [academicYear, semester, setValue]);

  useEffect(() => {
    if (modalState.state === "edit" && modalState.open) {
      handleSetFormValue();
    } else {
      reset();
    }
  }, [modalState.state, modalState.open, academicPeriod]);

  return (
    <Modal isOpen={modalState.open} centered size="lg">
      <div className="d-flex justify-content-between align-items-center pt-3 px-3">
        <p className="fs-4 fw-semibold mb-0 text-black">
          {modalState.state === "add"
            ? "Tambah Periode Akademik"
            : modalState.state === "edit"
            ? "Ubah Periode Akademik"
            : "Detail Periode Akademik"}
        </p>
        <Button className="bg-white border-0 p-0" onClick={handleToggleModal}>
          <i className=" ri-close-fill text-black fs-4"></i>
        </Button>
      </div>

      <ModalBody>
        <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
          <Row>
            {/* left section */}
            <Col className="d-flex flex-column gap-4">
              {/* code and is active */}
              <Col sm={12}>
                <Label htmlFor="period_code" className="form-label mb-1">
                  Kode Periode
                </Label>
                <Controller
                  name="code"
                  control={control}
                  render={({ field }) => (
                    <Input
                      className="form-control form-control-icon"
                      id="code"
                      placeholder="Kode Periode"
                      disabled
                      {...field}
                    />
                  )}
                />
              </Col>
              {/* academic year id */}
              <Col sm={12}>
                <Label htmlFor="academic_year_id" className="form-label mb-1">
                  Tahun Ajaran
                </Label>
                <Controller
                  name="academic_year_id"
                  control={control}
                  render={({ field }) => (
                    <SelectComponent
                      options={academicYearOptions as OptionType[]}
                      isLoading={isLoadingAcademicYears}
                      placeholder="Pilih Tahun Ajaran"
                      isError={!!errors.academic_year_id}
                      id={"birth_place_id"}
                      isDisabled={
                        modalState.state === "detail" ||
                        isLoadingAcademicYears ||
                        isLoadingAcademicPeriod
                      }
                      {...field}
                    />
                  )}
                />
                <FormErrorMessage errors={errors.academic_year_id} />
              </Col>
              {/* semester_id */}
              <Col sm={12}>
                <Label htmlFor="semester_id" className="form-label mb-1">
                  Semester
                </Label>

                <Controller
                  name="semester_id"
                  control={control}
                  render={({ field }) => (
                    <SelectComponent
                      options={semesterOptions as OptionType[]}
                      isLoading={isLoadingSemesters}
                      isDisabled={isLoadingSemesters || isLoadingAcademicPeriod}
                      placeholder="Pilih Semester"
                      isError={!!errors.semester_id}
                      id={"birth_place_id"}
                      {...field}
                    />
                  )}
                />
                <FormErrorMessage errors={errors.semester_id} />
              </Col>
              {/* fullname period */}
              <Col sm={12}>
                <Label htmlFor="fullname" className="form-label mb-1">
                  Nama Periode
                </Label>

                <Controller
                  name="fullname"
                  control={control}
                  render={({ field }) => (
                    <Input
                      className={`form-control form-control-icon ${
                        errors.fullname ? "border border-danger" : ""
                      }`}
                      id="fullname"
                      placeholder="Nama Periode"
                      disabled={isLoadingAcademicPeriod}
                      {...field}
                    />
                  )}
                />
                <FormErrorMessage errors={errors.fullname} />
              </Col>
              {/* shortname period */}
              <Col sm={12}>
                <Label htmlFor="short_period_name" className="form-label mb-1">
                  Nama Singkat Periode
                </Label>
                <Controller
                  name="shortname"
                  control={control}
                  render={({ field }) => (
                    <Input
                      className={`form-control form-control-icon ${
                        errors.shortname ? "border border-danger" : ""
                      }`}
                      id="shortname"
                      placeholder="Nama Singkat Periode"
                      disabled={isLoadingAcademicPeriod}
                      {...field}
                    />
                  )}
                />
              </Col>
              {/* number of lecture meeting */}
              <Col sm={12}>
                <Label
                  htmlFor="number_of_lecture_meeting"
                  className="form-label mb-1"
                >
                  Jumlah Pertemuan
                </Label>

                <Controller
                  name="number_of_lecture_meeting"
                  control={control}
                  render={({ field }) => (
                    <SelectComponent
                      options={Array.from({ length: 20 }, (_, i) => i + 1).map(
                        (item) => ({
                          value: item.toString(),
                          label: item.toString(),
                        })
                      )}
                      isDisabled={isLoadingAcademicPeriod}
                      placeholder="Pilih Jumlah Pertemuan"
                      isError={!!errors.number_of_lecture_meeting}
                      id={"birth_place_id"}
                      {...field}
                    />
                  )}
                />
                <FormErrorMessage errors={errors.number_of_lecture_meeting} />
              </Col>
            </Col>
            {/* right section */}
            <Col className="d-flex flex-column gap-4">
              {/* start date of college */}
              <Col sm={12}>
                <Label
                  htmlFor="start_date_of_college"
                  className="form-label mb-1"
                >
                  Tanggal Mulai Kuliah
                </Label>
                <div className="form-icon">
                  <Controller
                    name="start_date_of_college"
                    control={control}
                    render={({ field }) => {
                      return (
                        <DatePicker
                          onChange={(e) => field.onChange(e)}
                          value={field.value}
                          className={`p-0 ${
                            errors.start_date_of_college
                              ? "border border-danger"
                              : ""
                          }`}
                          classNameFlatpickr={`form-control form-control-icon disabled-input`}
                          options={{
                            mode: "single",
                            dateFormat: "d F Y",
                          }}
                          disabled={isLoadingAcademicPeriod}
                        />
                      );
                    }}
                  />
                  <i style={{ left: "15px" }}>
                    <CalendarTodayIcon color="#878A99" />
                  </i>
                </div>
                <FormErrorMessage errors={errors.start_date_of_college} />
              </Col>
              {/* end date of college */}
              <Col sm={12}>
                <Label
                  htmlFor="end_date_of_college"
                  className="form-label mb-1"
                >
                  Tanggal Akhir Kuliah
                </Label>
                <div className="form-icon">
                  <Controller
                    name="end_date_of_college"
                    control={control}
                    render={({ field }) => {
                      return (
                        <DatePicker
                          onChange={(e) => field.onChange(e)}
                          value={field.value}
                          className={`p-0 ${
                            errors.end_date_of_college
                              ? "border border-danger"
                              : ""
                          }`}
                          classNameFlatpickr={`form-control form-control-icon disabled-input`}
                          options={{
                            mode: "single",
                            dateFormat: "d F Y",
                            minDate: startCollege[0],
                          }}
                          disabled={isLoadingAcademicPeriod}
                        />
                      );
                    }}
                  />
                  <i style={{ left: "15px" }}>
                    <CalendarTodayIcon color="#878A99" />
                  </i>
                </div>
                <FormErrorMessage errors={errors.end_date_of_college} />
              </Col>
              {/* start date of uts */}
              <Col sm={12}>
                <Label htmlFor="start_date_of_uts" className="form-label mb-1">
                  Tanggal Mulai UTS
                </Label>
                <div className="form-icon">
                  <Controller
                    name="start_date_of_uts"
                    control={control}
                    render={({ field }) => {
                      return (
                        <DatePicker
                          onChange={(e) => field.onChange(e)}
                          value={field.value as Date[]}
                          className={`p-0 ${
                            errors.start_date_of_uts
                              ? "border border-danger"
                              : ""
                          }`}
                          classNameFlatpickr={`form-control form-control-icon disabled-input`}
                          options={{
                            mode: "single",
                            dateFormat: "d F Y",
                          }}
                          disabled={
                            isLoadingAcademicPeriod ||
                            modalState.state === "add"
                          }
                        />
                      );
                    }}
                  />
                  <i style={{ left: "15px" }}>
                    <CalendarTodayIcon color="#878A99" />
                  </i>
                </div>
                <FormErrorMessage errors={errors.start_date_of_uts} />
              </Col>
              {/* end date of uts */}
              <Col sm={12}>
                <Label htmlFor="end_date_of_uts" className="form-label mb-1">
                  Tanggal Selesai UTS
                </Label>
                <div className="form-icon">
                  <Controller
                    name="end_date_of_uts"
                    control={control}
                    render={({ field }) => {
                      return (
                        <DatePicker
                          onChange={(e) => field.onChange(e)}
                          value={field.value as Date[]}
                          className={`p-0 ${
                            errors.end_date_of_uts ? "border border-danger" : ""
                          }`}
                          classNameFlatpickr={`form-control form-control-icon disabled-input`}
                          options={{
                            mode: "single",
                            dateFormat: "d F Y",
                            minDate: startUts?.[0],
                          }}
                          disabled={
                            isLoadingAcademicPeriod ||
                            modalState.state === "add"
                          }
                        />
                      );
                    }}
                  />
                  <i style={{ left: "15px" }}>
                    <CalendarTodayIcon color="#878A99" />
                  </i>
                </div>
                <FormErrorMessage errors={errors.end_date_of_uts} />
              </Col>
              {/* start date of uas */}
              <Col sm={12}>
                <Label htmlFor="start_date_of_uas" className="form-label mb-1">
                  Tanggal Mulai UAS
                </Label>
                <div className="form-icon">
                  <Controller
                    name="start_date_of_uas"
                    control={control}
                    render={({ field }) => {
                      return (
                        <DatePicker
                          onChange={(e) => field.onChange(e)}
                          value={field.value as Date[]}
                          className={`p-0 ${
                            errors.start_date_of_uas
                              ? "border border-danger"
                              : ""
                          }`}
                          classNameFlatpickr={`form-control form-control-icon disabled-input`}
                          options={{
                            mode: "single",
                            dateFormat: "d F Y",
                          }}
                          disabled={
                            isLoadingAcademicPeriod ||
                            modalState.state === "add"
                          }
                        />
                      );
                    }}
                  />
                  <i style={{ left: "15px" }}>
                    <CalendarTodayIcon color="#878A99" />
                  </i>
                </div>
                <FormErrorMessage errors={errors.start_date_of_uas} />
              </Col>
              {/* end date of uas */}
              <Col sm={12}>
                <Label htmlFor="end_date_of_uas" className="form-label mb-1">
                  Tanggal Selesai UAS
                </Label>
                <div className="form-icon">
                  <Controller
                    name="end_date_of_uas"
                    control={control}
                    render={({ field }) => {
                      return (
                        <DatePicker
                          onChange={(e) => field.onChange(e)}
                          value={field.value as Date[]}
                          className={`p-0 ${
                            errors.end_date_of_uas ? "border border-danger" : ""
                          }`}
                          classNameFlatpickr={`form-control form-control-icon disabled-input`}
                          options={{
                            mode: "single",
                            dateFormat: "d F Y",
                            minDate: startUas?.[0],
                          }}
                          disabled={
                            isLoadingAcademicPeriod ||
                            modalState.state === "add"
                          }
                        />
                      );
                    }}
                  />
                  <i style={{ left: "15px" }}>
                    <CalendarTodayIcon color="#878A99" />
                  </i>
                </div>
                <FormErrorMessage errors={errors.end_date_of_uas} />
              </Col>
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

export default ModalAcademicPeriod;
