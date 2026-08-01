"use client";

import { Button, Col, Label, Modal, ModalBody, Row, Spinner } from "reactstrap";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useParams } from "next/navigation";
import { useEffect } from "react";

import { useModalContext } from "@/lib/hooks/use-modal";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { SelectComponent } from "@/components/ui/select";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import {
  classScheduleTemplateFormSchema,
  ClassScheduleFormTemplateSchemaType,
} from "@/lib/validations/settings/academic-period/form-class-schedule";
import { DatePicker } from "@/components/ui/date-picker";
import { useWeekDaysOptions } from "@/lib/hooks/use-days";
import { classMerge } from "@/lib/utils/class-merge";
import { ScheduleIcon } from "@/components/icons/schedule";
import { createClassScheduleTemplate } from "@/services/api/settings/academic-period/class-schedule/create-class-schedulte-template";
import { useGetClassScheduleTemplateById } from "@/services/api/settings/academic-period/class-schedule/use-get-detail-class-schedule-template-by-id";
import { updateClassScheduleTemplate } from "@/services/api/settings/academic-period/class-schedule/update-class-schedule-template";

export const ModalClassScheduleTemplate = () => {
  const { modalState, setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const { classId } = useParams();

  const {
    data: classSchedule,
    isLoading: isLoadingClassSchedule,
    refetch: refetchClassSchedule,
  } = useGetClassScheduleTemplateById({
    classId: classId as string,
    classScheduleTemplateId: modalState.id as string,
  });

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    reset,
    watch,
    setValue,
  } = useForm({
    resolver: zodResolver(classScheduleTemplateFormSchema),
    defaultValues: {
      start_time: [],
      end_time: [],
    },
  });

  const startTimeForm = watch("start_time");

  const dayOptions = useWeekDaysOptions();

  const typeOfMeetingOptions = [
    {
      label: "Offline",
      value: "offline",
    },
    {
      label: "Online",
      value: "online",
    },
  ];

  const findDay = dayOptions.find(
    (data) => data.value === classSchedule?.data?.day_name
  );

  const findTypeOfMeeting = typeOfMeetingOptions.find(
    (data) => data.value === classSchedule?.data?.type_of_meeting
  );

  const toggleModal = () => {
    setModalState((prev) => ({
      ...prev,
      open: !modalState.open,
    }));
    reset();
  };

  const onSubmit = async (payload: ClassScheduleFormTemplateSchemaType) => {
    try {
      const response =
        modalState.state === "add"
          ? await createClassScheduleTemplate({
              classId: classId as string,
              payload,
            })
          : await updateClassScheduleTemplate({
              classId: classId as string,
              scheduleTemplateId: modalState.id as string,
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
        message: `${
          modalState.state === "add" ? "Tambah" : "Update"
        } data berhasil`,
        state: "success",
      }));

      return toggleModal();
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
    if (modalState.open && modalState.state !== "add") {
      refetchClassSchedule();
    }
  }, [modalState.open, modalState.state]);

  useEffect(() => {
    if (classSchedule && modalState.open && modalState.state !== "add") {
      setValue("day_name", findDay as OptionType);
      setValue("type_of_meeting", findTypeOfMeeting as OptionType);
    } else {
      reset();
    }
  }, [modalState.open, classSchedule, modalState.state]);

  return (
    <Modal isOpen={modalState.open} centered>
      <div className="d-flex justify-content-between align-items-center pt-3 px-3 border-bottom pb-2">
        <p className="fs-4 fw-semibold mb-0 text-black">
          {modalState.state === "add"
            ? "Tambah Jadwal Mingguan"
            : modalState.state === "edit"
            ? "Ubah Jadwal Mingguan"
            : "Detail Jadwal Mingguan"}
        </p>
        <Button className="bg-white border-0 p-0" onClick={toggleModal}>
          <i className=" ri-close-fill text-black fs-4"></i>
        </Button>
      </div>
      <ModalBody>
        <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
          <Row className="gap-3">
            {/* day_name */}
            <Col sm={12}>
              <Label htmlFor="day_name" className="form-label mb-1">
                Hari
              </Label>
              <Controller
                name="day_name"
                control={control}
                render={({ field }) => (
                  <SelectComponent
                    options={dayOptions as OptionType[]}
                    placeholder="Pilih Hari"
                    isError={!!errors.day_name}
                    id={"day_name"}
                    {...field}
                    isDisabled={isLoadingClassSchedule}
                    isLoading={isLoadingClassSchedule}
                  />
                )}
              />
              <FormErrorMessage errors={errors.day_name} />
            </Col>
            {/* start time */}
            <Col sm={12}>
              <Label htmlFor="start_time" className="form-label mb-1">
                Jam Mulai
              </Label>
              <div className="form-icon">
                <Controller
                  name="start_time"
                  control={control}
                  render={({ field }) => {
                    return (
                      <DatePicker
                        onChange={(e) => field.onChange(e)}
                        value={field.value}
                        className={`p-0 ${
                          errors.start_time ? "border border-danger" : ""
                        }`}
                        classNameFlatpickr={`form-control form-control-icon disabled-input`}
                        options={{
                          dateFormat: "H:i",
                          enableTime: true,
                          noCalendar: true,
                          time_24hr: true,
                        }}
                        placeholder="Masukkan Jam Mulai"
                      />
                    );
                  }}
                />
                <i style={{ left: "15px" }}>
                  <ScheduleIcon color="#878A99" />
                </i>
              </div>
              <FormErrorMessage errors={errors.start_time} />
            </Col>
            {/* end time */}
            <Col sm={12}>
              <Label htmlFor="end_time" className="form-label mb-1">
                Jam Selesai
              </Label>
              <div className="form-icon">
                <Controller
                  name="end_time"
                  control={control}
                  render={({ field }) => {
                    return (
                      <DatePicker
                        onChange={(e) => field.onChange(e)}
                        value={field.value}
                        className={`p-0 ${
                          errors.end_time ? "border border-danger" : ""
                        }`}
                        classNameFlatpickr={`form-control form-control-icon disabled-input`}
                        options={{
                          dateFormat: "H:i",
                          enableTime: true,
                          noCalendar: true,
                          time_24hr: true,
                          minTime: startTimeForm[0],
                        }}
                        placeholder="Masukkan Jam Selesai"
                      />
                    );
                  }}
                />
                <i style={{ left: "15px" }}>
                  <ScheduleIcon color="#878A99" />
                </i>
              </div>
              <FormErrorMessage errors={errors.end_time} />
            </Col>
            {/* type_of_meeting */}
            <Col sm={12}>
              <Label htmlFor="type_of_meeting" className="form-label mb-1">
                Jenis Pertemuan
              </Label>
              <Controller
                name="type_of_meeting"
                control={control}
                render={({ field }) => (
                  <SelectComponent
                    options={typeOfMeetingOptions as OptionType[]}
                    placeholder="Pilih Jenis Pertemuan"
                    isError={!!errors.type_of_meeting}
                    id={"type_of_meeting"}
                    {...field}
                    isLoading={isLoadingClassSchedule}
                    isDisabled={isLoadingClassSchedule}
                  />
                )}
              />
              <FormErrorMessage errors={errors.type_of_meeting} />
            </Col>
          </Row>
          <div className="d-flex justify-content-end mt-3">
            <Button
              type="button"
              className={classMerge(
                modalState.state === "detail" ? "btn-success" : "btn-light",
                " waves-effect waves-light me-2"
              )}
              onClick={toggleModal}
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
