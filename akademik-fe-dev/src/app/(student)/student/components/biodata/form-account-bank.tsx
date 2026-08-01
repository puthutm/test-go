"use client";

import { Button, Col, Input, Label, Row } from "reactstrap";
import { useCallback, useEffect, useState } from "react";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { SelectComponent } from "@/components/ui/select";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { EditIcon } from "@/components/icons/edit";
import { FormDescription } from "@/components/ui/form-description";
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import {
  FormBankAccountSchema,
  FormBankAccountSchemaType,
} from "@/lib/validations/students/biodata/form-bank-account-schema";
import { useBanks } from "@/services/api/data-referensi/bank/use-banks";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { useUpdateBankAccountStudent } from "@/services/api/students/biodata/bank-account/use-update-bank-account";
import { splitFileNameUploaded } from "@/lib/utils/split-filename-upload";
import { useBankAccountStudent } from "@/services/api/students/biodata/bank-account/use-get-bank-account";
import { useGetFileStorage } from "@/services/api/file-storage/use-get-file-storage";

export const FormBankAccount = () => {
  const [isEdit, setIsEdit] = useState<boolean>(false);
  const { setModalConfirmationState } = useModalConfirmationContext();

  const { data: dataBank, isLoading: isLoadingBank } = useBanks();
  const { data: dataBankAccount, isLoading: isLoadingBankAccount } =
    useBankAccountStudent();
  const { mutateAsync: updateBankAccountStudent } =
    useUpdateBankAccountStudent();
  const { getFileStorage, loading: loadingFile } = useGetFileStorage();

  const bankOptions = dataBank?.data?.map((bank) => ({
    label: bank.name,
    value: bank.id,
  }));

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    clearErrors,
    setValue,
  } = useForm<FormBankAccountSchemaType>({
    resolver: zodResolver(FormBankAccountSchema),
    defaultValues: {
      account_name: "",
      account_number: "",
      account_file_path: undefined,
    },
  });

  const handleSetFormValue = useCallback(() => {
    if (dataBankAccount?.data.bank_id) {
      setValue("bank_id", {
        label: dataBankAccount?.data.bank_name || null,
        value: dataBankAccount?.data.bank_id || null,
      });
    }
    setValue(
      "account_name",
      (dataBankAccount?.data.account_name as string) || ""
    );
    setValue(
      "account_number",
      (dataBankAccount?.data.account_number as string) || ""
    );
  }, [dataBankAccount?.data, setValue]);

  const onSubmit = async (payload: FormBankAccountSchemaType) => {
    try {
      const response = await updateBankAccountStudent(payload);

      if (response.error) {
        return setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          state: "failed",
          message: response.message,
        }));
      }

      setIsEdit(false);
      return setModalConfirmationState((prev) => ({
        ...prev,
        open: true,
        message: "Data berhasil di-update",
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
    if (dataBankAccount?.data) handleSetFormValue();
  }, [dataBankAccount?.data, handleSetFormValue]);

  if (dataBankAccount?.error) {
    return <h1>{dataBankAccount.message}</h1>;
  }

  return (
    <>
      <div className="d-flex border-2 border-bottom justify-content-between align-items-center">
        <h5 className="card-title py-3 mb-1" style={{ fontWeight: "500" }}>
          Rekening
        </h5>
        {!isEdit ? (
          <button
            className="bg-transparent rounded px-3 d-flex gap-1 align-items-center justify-content-center text-primary"
            style={{ border: "1px solid #10487A", paddingBlock: "8px" }}
            onClick={() => setIsEdit(true)}
          >
            <EditIcon />
            <span>Edit</span>
          </button>
        ) : null}
      </div>
      <form onSubmit={handleSubmit(onSubmit)} className="my-2">
        <Row className="gap-1 gap-lg-0">
          {/* left section */}
          <Col md={12} lg={6}>
            <Row className="gap-2">
              {/* Bank Name */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="bank_id"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Nama Bank
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="bank_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={bankOptions as OptionType[]}
                          isDisabled={!isEdit || isLoadingBankAccount}
                          isLoading={isLoadingBank}
                          placeholder="Pilih Bank"
                          isError={!!errors.bank_id}
                          id={"bank_id"}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.bank_id} />
                  </Col>
                </Row>
              </Col>
              {/* Account Number */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="account_number"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Nomor Rekening
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="account_number"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.account_number ? "border border-danger" : ""
                          }`}
                          id="account_number"
                          placeholder="Masukkan Nomor Rekening"
                          disabled={!isEdit || isLoadingBankAccount}
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);
                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.account_number} />
                  </Col>
                </Row>
              </Col>
            </Row>
          </Col>
          {/* right section */}
          <Col md={12} lg={6}>
            <Row className="gap-2">
              {/* Account Holder */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="account_name"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Nama Pemilik Rekening
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="account_name"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.account_name ? "border border-danger" : ""
                          }`}
                          id="account_name"
                          placeholder="Masukkan Nama Pemilik Rekening"
                          disabled={!isEdit || isLoadingBankAccount}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.account_name} />
                  </Col>
                </Row>
              </Col>
              {/* File Upload */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="account_file_path"
                      className="form-label mb-0 fw-medium optional"
                    >
                      File Rekening
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <div className="d-flex gap-2">
                      <Controller
                        name="account_file_path"
                        control={control}
                        render={({ field: { onChange, value, ...field } }) => (
                          <div className="position-relative w-100">
                            <Input
                              type="text"
                              className={`form-control form-control-icon ${
                                errors.account_file_path
                                  ? "border border-danger"
                                  : ""
                              }`}
                              value={
                                value instanceof File
                                  ? value.name
                                  : splitFileNameUploaded(
                                      dataBankAccount?.data
                                        ?.account_filepath as string
                                    ) || ""
                              }
                              disabled={!isEdit}
                              readOnly
                              placeholder="Pilih file"
                              {...field}
                            />
                            {isEdit && (
                              <Input
                                type="file"
                                id="account_file_path"
                                accept=".png,.jpg,.jpeg"
                                onChange={(e) => {
                                  const file = e.target.files?.[0];
                                  onChange(file);
                                }}
                                disabled={isLoadingBankAccount}
                                hidden
                              />
                            )}
                          </div>
                        )}
                      />
                      {isEdit && (
                        <label
                          htmlFor="account_file_path"
                          className="btn d-flex align-items-center btn-light mb-0"
                          style={{ whiteSpace: "nowrap" }}
                        >
                          Upload File
                        </label>
                      )}
                      {!isEdit && (
                        <button
                          className="btn d-flex align-items-center btn-light"
                          style={{ whiteSpace: "nowrap" }}
                          type="button"
                          onClick={
                            dataBankAccount?.data?.account_filepath
                              ? async () =>
                                  await getFileStorage(
                                    dataBankAccount?.data
                                      ?.account_filepath as string
                                  )
                              : () => null
                          }
                          disabled={loadingFile}
                        >
                          Lihat File
                        </button>
                      )}
                    </div>
                    {errors.account_file_path ? (
                      <FormErrorMessage errors={errors.account_file_path} />
                    ) : (
                      <FormDescription message="File dengan format .png dan .jpg maksimal 2mb" />
                    )}
                  </Col>
                </Row>
              </Col>
            </Row>
          </Col>
        </Row>
        {isEdit && (
          <div className="d-flex justify-content-between mt-3 gap-3">
            <button
              onClick={() => {
                setIsEdit(!isEdit);
                clearErrors();
              }}
              className="bg-transparent text-primary rounded px-3"
              type="button"
              style={{ border: "1px solid #10487A" }}
              disabled={isSubmitting}
            >
              <span>Batal</span>
            </button>
            <Button
              color="primary"
              className="d-flex flex-grow-0 justify-content-center align-items-center"
              type="submit"
              disabled={isSubmitting}
            >
              <span>Update</span>
            </Button>
          </div>
        )}
      </form>
    </>
  );
};
