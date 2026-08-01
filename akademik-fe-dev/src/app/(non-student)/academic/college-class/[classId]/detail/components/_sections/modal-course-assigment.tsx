"use client";


import { useModalContext } from "@/lib/hooks/use-modal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

import { schemeFormCourseAssigment,FormSchemeCourseAssigment } from "@/lib/validations/academic/settings/college-class/form-course-assigment";

import { useGetClassScheduleSubDetail } from "@/services/api/academic/lecturer/class-schedule/detail-class/class-schedule/use-get-class-schedule-sub-detail";
import { useGetDetailCourseAssignment } from "@/services/api/academic/lecturer/class-schedule/detail-class/course-assignment/use-get-detail-course-assigment";
import { createCourseAssignment } from "@/services/api/academic/lecturer/class-schedule/detail-class/course-assignment/create-course-assignment";
import { zodResolver } from "@hookform/resolvers/zod";

import { useParams } from "next/navigation";
import React,{useEffect} from "react";
import { Controller, SubmitHandler, useForm } from "react-hook-form";
import { useQueryClient } from "@tanstack/react-query";

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

import { CloseIcon } from "@/components/icons/close";
import { SelectComponent } from "@/components/ui/select";


import { handleInputNumberOnly } from "@/lib/utils/input-number-only";

function ModalCourseAssigment() {
    const queryClient = useQueryClient()
    const {classId} = useParams()
  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    setValue,
    reset,
    clearErrors,
  } = useForm<FormSchemeCourseAssigment>({
    resolver: zodResolver(schemeFormCourseAssigment),
    defaultValues: {
        title:'',
        description:'',
        deadline_of_assignment_submission:'',
        time_to_open:'',
        retake:''
    },
    mode: "onChange",
  });

  const { modalState, setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const { data: classSchedule, isLoading: isloadingClassSchedule } =
    useGetClassScheduleSubDetail(classId as string,modalState.open);

  const { data: detailCourseAssigment, isLoading: isloadingCourseAssigment } =
    useGetDetailCourseAssignment(classId as string,modalState.id);

  //! maping option schedule
  const mappingClassSchedule = classSchedule?.data?.data?.map((el: IClassScheduleSubDetail) => {
      return {
        label: `${el.session}`,
        value: el.id,
      };
    }) ?? [];

   
     const handleToogleModal = () => {
       setModalState((prev) => ({
         ...prev,
         id: null,
         open: false,
       }));
       reset();
       clearErrors();
     };
   
     const onSubmit : SubmitHandler<FormSchemeCourseAssigment> = async (data) => {
       try {
           const response = await createCourseAssignment(classId as string,data);
           if (response.error) {
             throw new Error(response.message);
           }
   
           setModalConfirmationState({
             open: true,
             message: "berhasil menambahkan data tugas kuliah",
             state: "success",
           });
            queryClient.invalidateQueries({
                queryKey:['get-course-assignment']
            })
            queryClient.refetchQueries({
                queryKey:['get-detail-course-assignment']
            })
           handleToogleModal();
       } catch (err: any) {
         setModalConfirmationState({
           open: true,
           message: err.message,
           state: "failed",
         });
       }
     };

  useEffect(() => {
    if (modalState.state !== "add" && detailCourseAssigment) {
      setValue("schedule_id", {
        label: `${detailCourseAssigment?.data?.session_schedule}`,
        value: detailCourseAssigment?.data?.schedule_id.toString(),
      });
      setValue(
        "title",
        detailCourseAssigment?.data?.title 
      );
      setValue(
        "description",
      detailCourseAssigment?.data?.description
      );
      setValue(
        "deadline_of_assignment_submission",
        String(detailCourseAssigment?.data?.deadline_of_assignment_submission as number)
      );
      setValue(
        "time_to_open",
        String(detailCourseAssigment?.data?.time_to_open ?? '')
      );

      setValue(
        "retake",
        String(detailCourseAssigment?.data?.retake ?? '')
      );
      setValue(
        "is_gradeable",
        detailCourseAssigment?.data?.is_gradeable as boolean 
      );
      setValue(
        "is_use_deadline",
        detailCourseAssigment?.data?.is_use_deadline as boolean
      );;
    }
  }, [modalState, detailCourseAssigment]);
   
  return (
        <Modal isOpen={modalState.open} centered>
          <section className="px-3 py-2  d-flex border-bottom justify-items-center">
            <p
              className="fs-4 fw-semibold p-0 mb-0 flex-grow-1"
              style={{ color: "#909090" }}
            >
                {
                    modalState.state === 'add' ? 'Tambah' :
                    modalState.state === 'edit' ? 'Ubah'
                    :
                    "Detail"
                } Tugas Kuliah
            </p>
            <Button className="bg-white border-0 p-0" onClick={handleToogleModal}>
              <CloseIcon color="#909090" />
            </Button>
          </section>

            <ModalBody>
                <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
                    <Row className="gap-2">
                        {/*//! schedule */}
                        <Col sm={12}>
                        <Label htmlFor="schedule_id" className="form-label mb-1 ">
                            Jadwal Perkuliahan
                        </Label>
                        <Controller
                            name="schedule_id"
                            control={control}
                            render={({ field }) => (
                            <SelectComponent
                                {...field}
                                options={mappingClassSchedule as OptionType[]}
                                isDisabled={isloadingCourseAssigment || modalState.state === "detail"}
                                isLoading={isloadingClassSchedule}
                                placeholder="Pilih Jadwal Perkuliahan"
                                isError={!!errors.schedule_id}
                                id={"schedule_id"}
                                // onChange={(value) => {
                                //   setValue("schedule_id", value.value);
                                // }}
                            />
                            )}
                        />
                        {errors?.schedule_id && (
                            <div className="text-danger fs-6">
                            {errors?.schedule_id.message}
                            </div>
                        )}
                        </Col>

                        {/*//! title  */}
                        <Col sm={12}>
                            <Label htmlFor="title" className="form-label mb-1">
                                Judul
                            </Label>
                            <Controller
                                control={control}
                                name="title"
                                render={({ field }) => (
                                <div className="">
                                    <Input
                                    {...field}
                                    className={`${
                                        errors.title && "border border-danger"
                                    }`}
                                    id="title"
                                    placeholder="Masukkan Judul"
                                    type="text"
                                    disabled={
                                    isloadingCourseAssigment || modalState.state === "detail"
                                    }
                                    />
                                </div>
                                )}
                            />
                            {errors?.title && (
                                <div className="text-danger">{errors?.title.message}</div>
                            )}
                        </Col>                          

                        {/*//! Deskripsi */}
                        <Col sm={12}>
                        <Label htmlFor="description" className="form-label mb-1">
                            Deskripsi
                        </Label>
                        <Controller
                            control={control}
                            name="description"
                            render={({ field }) => (
                            <div className="">
                                <Input
                                {...field}
                                className={`${
                                    errors.description && "border border-danger"
                                }`}
                                id="description"
                                placeholder="Masukkan deskripsi"
                                type="textarea"
                                disabled={
                                  isloadingCourseAssigment || modalState.state === "detail"
                                }
                                />
                            </div>
                            )}
                        />
                        {errors?.description && (
                            <div className="text-danger">{errors?.description.message}</div>
                        )}
                        </Col>

                        {/*//! deadline */}
                        <Col sm={12}>
                        <Label htmlFor="deadline_of_assignment_submission" className="form-label mb-1 ">
                            Batas Waktu Penyerahan 
                        </Label>
                        <Controller
                            control={control}
                            name="deadline_of_assignment_submission"
                            render={({ field }) => (
                            <div className="">
                                <Input
                                {...field}
                                className={`${
                                    errors.deadline_of_assignment_submission ? "border border-danger" : ""
                                }`}
                                id="deadline_of_assignment_submission"
                                placeholder="Batas waktu penyerahan"
                                type="text"
                                onChange={(e) => {
                                    const { numberValue } = handleInputNumberOnly(e);
                                    field.onChange(String(numberValue));
                                }}
                                disabled={
                                   isloadingCourseAssigment || modalState.state === "detail"
                                }
                                />
                            </div>
                            )}
                        />
                        {errors?.deadline_of_assignment_submission && (
                            <div className="text-danger fs-6">
                            {errors?.deadline_of_assignment_submission.message}
                            </div>
                        )}
                        </Col>

                        {/*//! time to open */}
                        <Col sm={12}>
                        <Label htmlFor="time_to_open" className="form-label mb-1 ">
                            Time To open 
                        </Label>
                        <Controller
                            control={control}
                            name="time_to_open"
                            render={({ field }) => (
                            <div className="">
                                <Input
                                {...field}
                                className={`${
                                    errors.time_to_open ? "border border-danger" : ""
                                }`}
                                id="time_to_open"
                                placeholder="Masukkan time to open"
                                type="text"
                                onChange={(e) => {
                                    const { numberValue } = handleInputNumberOnly(e);
                                    field.onChange(String(numberValue));
                                }}
                                disabled={
                                   isloadingCourseAssigment || modalState.state === "detail"
                                }
                                />
                            </div>
                            )}
                        />
                        {errors?.time_to_open && (
                            <div className="text-danger fs-6">
                            {errors?.time_to_open.message}
                            </div>
                        )}
                        </Col>

                        {/*//! retake */}
                        <Col sm={12}>
                        <Label htmlFor="retake" className="form-label mb-1 ">
                            Retake 
                        </Label>
                        <Controller
                            control={control}
                            name="retake"
                            render={({ field }) => (
                            <div className="">
                                <Input
                                {...field}
                                className={`${
                                    errors.retake ? "border border-danger" : ""
                                }`}
                                id="retake"
                                placeholder="Masukkan retake"
                                type="text"
                                onChange={(e) => {
                                    const { numberValue } = handleInputNumberOnly(e);
                                    field.onChange(String(numberValue));
                                }}
                                disabled={
                                   isloadingCourseAssigment || modalState.state === "detail"
                                }
                                />
                            </div>
                            )}
                        />
                        {errors?.retake && (
                            <div className="text-danger fs-6">
                            {errors?.retake.message}
                            </div>
                        )}
                        </Col>   

                        {/*//! is_gradeable  */}
                        <Col sm={12}>
                            <Label htmlFor="is_gradeable" className="form-label mb-1">
                                Is gradeable
                            </Label>
                            <Controller
                                control={control}
                                name="is_gradeable"
                                render={({ field }) => (
                                <div className=" form-check form-switch ms-1">
                                    <Input
                                    {...field}
                                    className={`${
                                        errors.is_gradeable && "border border-danger"
                                    }`}
                                    disabled={isloadingCourseAssigment || modalState.state === 'detail'}
                                    id="is_gradeable"
                                    role='switch'
                                    type="switch"
                                    checked={field.value ?? false}
                                    value={field.value ? "true" : "false"}
                                    onChange={e => field.onChange(e.target.checked)}
                                    style={{
                                        transform:'scale(1.5)'
                                    }}
                                    />
                                </div>
                                )}
                            />
                            {errors?.is_gradeable && (
                                <div className="text-danger">{errors?.is_gradeable.message}</div>
                            )}
                        </Col>     

                        {/*//! is use deadline  */}
                        <Col sm={12}>
                            <Label htmlFor="is_use_deadline" className="form-label mb-1">
                                Is Use Deadline
                            </Label>
                            <Controller
                                control={control}
                                name="is_use_deadline"
                                render={({ field }) => (
                                <div className=" form-check form-switch ms-1">
                                    <Input
                                    {...field}
                                    className={`${
                                        errors.is_use_deadline && "border border-danger"
                                    }`}
                                    disabled={isloadingCourseAssigment || modalState.state === 'detail'}
                                    id="is_use_deadline"
                                    role='switch'
                                    type="switch"
                                    checked={field.value ?? false}
                                    value={field.value ? "true" : "false"}
                                    onChange={e => field.onChange(e.target.checked)}
                                    style={{
                                        transform:'scale(1.5)'
                                    }}
                                    />
                                </div>
                                )}
                            />
                            {errors?.is_use_deadline && (
                                <div className="text-danger">{errors?.is_use_deadline.message}</div>
                            )}
                        </Col>                     
                    </Row>

                    <div className="d-flex justify-content-end mt-3 gap-2">
                        <Button
                        type="button"
                        className="waves-effect waves-light"
                        color="light"
                        onClick={handleToogleModal}
                        style={{ width: "80px" }}
                        >
                        Tutup
                        </Button>

                        {
                            modalState.state !== 'detail' &&   
                            <Button
                            className="btn waves-effect waves-light"
                            disabled={isSubmitting || isloadingCourseAssigment}
                            style={{ width: "80px" }}
                            color="primary"
                        >
                            {isSubmitting ? (
                            <Spinner size={"sm"} />
                            ) : modalState.state === "add" ? (
                            "Tambah"
                            ) : (
                            "Ubah"
                            )}
                        </Button>
                        }


                    </div>
                </form>
            </ModalBody>
        </Modal>
  )
}

export default ModalCourseAssigment