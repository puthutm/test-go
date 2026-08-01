"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { Controller, useForm } from "react-hook-form";
import { Button, Col, Label, Modal, ModalBody, Row, Spinner } from "reactstrap";

import { FormErrorMessage } from "@/components/ui/form-error-message";
import { useModalContext } from "@/lib/hooks/use-modal";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import {
  formKrsStatusSchema,
  FormKrsStatusSchemaType,
} from "@/lib/validations/lecturer/update-status-reject-approve-krs";
import { useUpdateStatusKrs } from "@/services/api/lectures/krs/use-update-status-krs";

export const ModalUpdateStatusKrs = () => {
  const { modalState, setModalState } = useModalContext();
  const { setModalConfirmationState } = useModalConfirmationContext();

  const { mutateAsync: updateKrsStatus, isPending } = useUpdateStatusKrs();

  const {
    control,
    handleSubmit,
    formState: { isSubmitting },
  } = useForm<FormKrsStatusSchemaType>({
    resolver: zodResolver(formKrsStatusSchema),
    defaultValues: {
      item_status: "rejected",
      reject_reason: "",
    },
  });

  const handleToggleModal = () => {
    setModalState((prev) => ({
      ...prev,
      open: !prev.open,
      id: undefined,
    }));
  };

  const onSubmit = async (data: FormKrsStatusSchemaType) => {
    try {
      const response = await updateKrsStatus({
        krsItemId: modalState?.id as string,
        payload: data,
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
      <div className="d-flex justify-content-between align-items-center pt-3 px-3">
        <p className="fs-4 fw-semibold mb-0 text-black">KRS Status</p>
        <Button className="bg-white border-0 p-0" onClick={handleToggleModal}>
          <i className=" ri-close-fill text-black fs-4"></i>
        </Button>
      </div>
      <ModalBody>
        <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
          <Row>
            <Controller
              name="reject_reason"
              control={control}
              render={({ field, fieldState }) => (
                <Col sm={12}>
                  <Label htmlFor="reject_reason" className="form-label mb-1">
                    Alasan
                  </Label>
                  <textarea
                    className={`form-control ${
                      fieldState.error?.message ? "border border-danger" : ""
                    }`}
                    id="reject_reason"
                    placeholder="Alasan ditolak"
                    disabled={isPending}
                    {...field}
                  />
                  <FormErrorMessage errors={fieldState?.error} />
                </Col>
              )}
            />
          </Row>
          <div className="d-flex justify-content-end mt-3">
            <Button
              type="button"
              className={"btn-light waves-effect waves-light me-2"}
              onClick={handleToggleModal}
              disabled={isPending}
            >
              Tutup
            </Button>
            {modalState.state !== "detail" && (
              <Button disabled={isSubmitting} color="primary">
                {isSubmitting ? <Spinner size={"sm"} /> : "Simpan"}
              </Button>
            )}
          </div>
        </form>
      </ModalBody>
    </Modal>
  );
};
