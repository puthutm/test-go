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

import { FormErrorMessage } from "@/components/ui/form-error-message";
import { SelectComponent } from "@/components/ui/select";
import { useModalContext } from "@/lib/hooks/use-modal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import {
  StudentAccountSchema,
  StudentAccountSchemaType,
} from "@/lib/validations/academic/portal/student-account-schema";
import { createStudent } from "@/services/api/portal/academic/create-student";
import { useGetBatches } from "@/services/api/portal/academic/use-get-batches";

export const ModalStudentAccount = () => {
  const { modalState, setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const { data: batches, isLoading: isLoadingBatches } = useGetBatches();

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    reset,
  } = useForm<StudentAccountSchemaType>({
    resolver: zodResolver(StudentAccountSchema),
    defaultValues: {
      batch_detail_id: "",
      nik: "",
      name: "",
      email: "",
      password: "",
      phone: "",
    },
  });

  const handleToggleModal = () => {
    setModalState((prev) => ({
      ...prev,
      open: !prev.open,
      id: undefined,
    }));
  };

  const onSubmit = async (payload: StudentAccountSchemaType) => {
    try {
      const response = await createStudent(payload);

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
        message: "Tambah data mahasiswa berhasil",
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

  useEffect(() => {
    if (!modalState.open) {
      reset();
    }
  }, [modalState.open, reset]);

  const batchOptions =
    batches?.data?.map((batch) => ({
      label: batch.batch_name,
      value: batch.batch_detail_id,
    })) || [];

  return (
    <Modal isOpen={modalState.open} centered>
      <div className="d-flex justify-content-between align-items-center pt-3 px-3">
        <p className="fs-4 fw-semibold mb-0 text-black">
          {modalState.state === "add"
            ? "Tambah Akun Mahasiswa"
            : modalState.state === "edit"
            ? "Ubah Akun Mahasiswa"
            : "Detail Akun Mahasiswa"}
        </p>
        <Button className="bg-white border-0 p-0" onClick={handleToggleModal}>
          <i className="ri-close-fill text-black fs-4"></i>
        </Button>
      </div>
      <ModalBody>
        <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
          <Row className="gy-3">
            <Col sm={12}>
              <Label htmlFor="batch_detail_id" className="form-label mb-1">
                Batch <span className="text-danger">*</span>
              </Label>
              <Controller
                name="batch_detail_id"
                control={control}
                render={({ field: { onChange, value, name } }) => (
                  <SelectComponent
                    id={name}
                    options={batchOptions}
                    value={batchOptions.find((opt) => opt.value === value)}
                    onChange={(val: any) => onChange(val?.value)}
                    placeholder="Pilih Batch"
                    isLoading={isLoadingBatches}
                  />
                )}
              />
              <FormErrorMessage errors={errors.batch_detail_id} />
            </Col>

            <Col sm={12}>
              <Label htmlFor="nik" className="form-label mb-1">
                NIK <span className="text-danger">*</span>
              </Label>
              <Controller
                name="nik"
                control={control}
                render={({ field: { onChange, ...field } }) => (
                  <Input
                    {...field}
                    id="nik"
                    placeholder="Masukkan NIK"
                    className={errors.nik ? "border-danger" : ""}
                    onChange={(e) => {
                      const { stringValue } = handleInputNumberOnly(e);
                      onChange(stringValue);
                    }}
                  />
                )}
              />
              <FormErrorMessage errors={errors.nik} />
            </Col>

            <Col sm={12}>
              <Label htmlFor="name" className="form-label mb-1">
                Nama Lengkap <span className="text-danger">*</span>
              </Label>
              <Controller
                name="name"
                control={control}
                render={({ field }) => (
                  <Input
                    {...field}
                    id="name"
                    placeholder="Masukkan Nama Lengkap"
                    className={errors.name ? "border-danger" : ""}
                  />
                )}
              />
              <FormErrorMessage errors={errors.name} />
            </Col>

            <Col sm={12}>
              <Label htmlFor="email" className="form-label mb-1">
                Email <span className="text-danger">*</span>
              </Label>
              <Controller
                name="email"
                control={control}
                render={({ field }) => (
                  <Input
                    {...field}
                    type="email"
                    id="email"
                    placeholder="Masukkan Email"
                    className={errors.email ? "border-danger" : ""}
                  />
                )}
              />
              <FormErrorMessage errors={errors.email} />
            </Col>

            <Col sm={12}>
              <Label htmlFor="phone" className="form-label mb-1">
                Nomor Telepon <span className="text-danger">*</span>
              </Label>
              <Controller
                name="phone"
                control={control}
                render={({ field: { onChange, ...field } }) => (
                  <Input
                    {...field}
                    id="phone"
                    placeholder="Masukkan Nomor Telepon"
                    className={errors.phone ? "border-danger" : ""}
                    onChange={(e) => {
                      const { stringValue } = handleInputNumberOnly(e);
                      onChange(stringValue);
                    }}
                  />
                )}
              />
              <FormErrorMessage errors={errors.phone} />
            </Col>

            <Col sm={12}>
              <Label htmlFor="password" className="form-label mb-1">
                Password
              </Label>
              <Controller
                name="password"
                control={control}
                render={({ field }) => (
                  <Input
                    {...field}
                    type="password"
                    id="password"
                    placeholder="Default: 12345678"
                    className={errors.password ? "border-danger" : ""}
                  />
                )}
              />
              <FormErrorMessage errors={errors.password} />
              <small className="text-muted">
                Kosongkan jika ingin menggunakan password default.
              </small>
            </Col>
          </Row>

          <div className="d-flex justify-content-end mt-4">
            <Button
              type="button"
              className="btn-light waves-effect waves-light me-2"
              onClick={handleToggleModal}
            >
              Batal
            </Button>
            <Button disabled={isSubmitting} color="primary">
              {isSubmitting ? <Spinner size="sm" /> : "Simpan"}
            </Button>
          </div>
        </form>
      </ModalBody>
    </Modal>
  );
};
