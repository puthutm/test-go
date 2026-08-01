"use client";

import { useModalContext } from "@/lib/hooks/use-modal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";

import {
  CreditLimitFormSchema,
  CreditLimitFormType,
} from "@/lib/validations/settings/credit-limit";

import { createSksLimits } from "@/services/api/settings/sks-limit/create-sks-limits";
import { editSksLimit } from "@/services/api/settings/sks-limit/update-sks-limits";
import { useGetDetailSksLimit } from "@/services/api/settings/sks-limit/use-get-detail-sks-limits";
import { zodResolver } from "@hookform/resolvers/zod";
import React, { useEffect } from "react";
import { Controller, useForm } from "react-hook-form";


import {
  Button,
  Col,
  Input,
  Label,
  Modal,
  ModalBody,
  ModalHeader,
  Row,
  Spinner,
} from "reactstrap";

const ModalCreditLimit = () => {
  const { modalState, setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();


  const {
    data: detailSksLimitData,
    isLoading: isLoadingSksLimit,
  } = useGetDetailSksLimit(modalState?.id as string | null);

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    setValue,
    reset,
    clearErrors
  } = useForm<CreditLimitFormType>({
    resolver: zodResolver(CreditLimitFormSchema),
    mode: "onChange",
    defaultValues: {
      ips_min: "",
      ips_max: "",
      sks_limit: "",
    }
  });

  const handleToggleModal = () => {
    setModalState((prev) => ({
      ...prev,
      id:null,
      open: !prev.open,
    }));
    reset();
    clearErrors()
  };



  const onSubmit = async (data: CreditLimitFormType) => {
    try {
      if (modalState.state === "add") {
        const response = await createSksLimits(data);
        if (response.error) {
          throw new Error(response.message);
        }

        setModalConfirmationState({
          open: true,
          message: "berhasil menambahkan data batas sks",
          state: "success",
        });
        handleToggleModal();
      } else {
        const response = await editSksLimit(modalState.id as string, data);
        if (response.error) {
          throw new Error(response.message);
        }

        setModalConfirmationState({
          open: true,
          message: "berhasil mengubah data batas sks",
          state: "success",
        });
        handleToggleModal();
      }
    } catch (err: any) {
      setModalConfirmationState({
        open: true,
        message: err.message,
        state: "failed",
      });
    }
  };


    useEffect(() => {
      if (modalState.state === "edit" && detailSksLimitData) {
          setValue("ips_min", String(detailSksLimitData?.data?.ips_min));
          setValue("ips_max", String(detailSksLimitData?.data?.ips_max));
          setValue("sks_limit", String(detailSksLimitData?.data?.sks_limit));
      }
    }, [modalState, detailSksLimitData]);



  return (
    <Modal isOpen={modalState.open} centered>
      <div className="d-flex justify-content-end px-2 pt-1">
        <Button className="bg-white border-0 p-0" onClick={handleToggleModal}>
          <i className=" ri-close-fill text-black fs-4"></i>
        </Button>
      </div>

      <ModalHeader className="ps-3 pt-0">
        <p className="fs-4 fw-semibold mb-0 text-black">
          {modalState.state === "add"
            ? "Tambah"
            : modalState.state === "edit"
            ? "Ubah"
            : "Detail"}{" "}
          Batas SKS
        </p>
      </ModalHeader>

      <ModalBody>
        <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
          <Row className="gap-2">
            <Col sm={12}>
              <Label htmlFor="ips_min" className="form-label  mb-1">
                IPS Min
              </Label>
              <Controller
                control={control}
                name="ips_min"
                render={({ field }) => (
                  <Input
                    {...field}
                    className={`form-control ${errors.ips_min && 'border border-danger'}`}
                    id="ips_min"
                    placeholder="Masukkan Nilai Minimal"
                    type="text"
                    disabled={
                      isLoadingSksLimit || modalState.state === "detail"
                    }
                    onChange={(e) => {
                     const {fractionalValue} = handleInputNumberOnly(e);
                     field.onChange(fractionalValue);
                    }}
                  />
                )}
              />
              {errors.ips_min && (
                <p className="text-danger">{errors.ips_min.message}</p>
              )}
            </Col>

            <Col sm={12}>
              <Label htmlFor="ips_max" className="form-label  mb-1">
                IPS Max
              </Label>
              <Controller
                control={control}
                name="ips_max"
                render={({ field }) => (
                  <Input
                    {...field}
                    className={`form-control ${errors.ips_max && 'border border-danger'}`}
                    id="ips_max"
                    placeholder="Masukkan Nilai Maksimal"
                    type="text"
                    disabled={
                      isLoadingSksLimit || modalState.state === "detail"
                    }
                    onChange={(e) => {
                     const {fractionalValue} = handleInputNumberOnly(e);
                     field.onChange(fractionalValue);
                    }}
                  />
                )}
              />
              {errors.ips_max && (
                <p className="text-danger">{errors.ips_max.message}</p>
              )}
            </Col>

            <Col sm={12}>
              <Label htmlFor="sks_limit" className="form-label  mb-1">
                Batas SKS
              </Label>
              <Controller
                control={control}
                name="sks_limit"
                render={({ field }) => (
                  <Input
                    {...field}
                    className={`form-control ${errors.sks_limit && 'border border-danger'}`}
                    id="sks_limit"
                    placeholder="Masukkan Batas SKS"
                    type="text"
                    disabled={
                      isLoadingSksLimit || modalState.state === "detail"
                    }
                     onChange={(e) => {
                     const {stringValue} = handleInputNumberOnly(e);
                     field.onChange(stringValue);
                    }}
                  />
                )}
              />
              {errors.sks_limit && (
                <p className="text-danger">{errors.sks_limit.message}</p>
              )}
            </Col>
          </Row>

          <div className="d-flex justify-content-end mt-3 gap-2">
            <Button
              type="button"
              className={'waves-effect waves-light'}
              color="light"
              onClick={handleToggleModal}
              style={{ width: "80px" }}
            >
              Tutup
            </Button>
            {modalState.state !== "detail" && (
              <Button
                className="btn  waves-effect waves-light"
                disabled={isSubmitting}
                color="primary"
                style={{ width: "80px" }}
              >
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

export default ModalCreditLimit;
