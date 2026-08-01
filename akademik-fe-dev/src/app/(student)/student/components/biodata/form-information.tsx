"use client";

import { Button, Col, Input, Label, Row } from "reactstrap";
import { useCallback, useEffect, useState } from "react";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { SelectComponent } from "@/components/ui/select";

import {
  FormInformationSchema,
  FormInformationSchemaType,
} from "@/lib/validations/students/biodata/form-information-schema";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import { EditIcon } from "@/components/icons/edit";
import { FormDescription } from "@/components/ui/form-description";
import { useTransportations } from "@/services/api/data-referensi/transportation/use-get-transportation";
import { useAlmamaterSizes } from "@/services/api/data-referensi/almamater-size/use-get-religions";
import { useJobs } from "@/services/api/data-referensi/job/use-get-jobs";
import { useCountries } from "@/services/api/data-referensi/country/use-get-countries";
import { updateInformationStudent } from "@/services/api/students/biodata/information/update-information";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

export const FormInformation = ({
  information,
}: {
  information: ApiResponse<InformationStudent>;
}) => {
  const [isEdit, setIsEdit] = useState<boolean>(false);

  const { setModalConfirmationState } = useModalConfirmationContext();

  const {
    formState: { errors },
    handleSubmit,
    control,
    clearErrors,
    setValue,
  } = useForm<FormInformationSchemaType>({
    resolver: zodResolver(FormInformationSchema),
    defaultValues: {
      college_email: "",
      private_email: "",
      phone: "",
    },
  });

  const { data: dataTransportation, isLoading: isLoadingTransportation } =
    useTransportations();

  const { data: dataAlmamaterSize, isLoading: isLoadingAlmamaterSize } =
    useAlmamaterSizes();

  const { data: dataJob, isLoading: isLoadingJob } = useJobs();

  const { data: dataCountry, isLoading: isLoadingCountry } = useCountries();

  const transportationOptions = dataTransportation?.data?.map((val) => ({
    label: val.name,
    value: val.id,
  }));

  const almamaterOptions = dataAlmamaterSize?.data?.map((val) => ({
    label: `${val.code} (${val.size})`,
    value: val.id,
  }));

  const jobOptions = dataJob?.data?.map((val) => ({
    label: val.name,
    value: val.id,
  }));

  const countryOptions = dataCountry?.data?.map((val) => ({
    label: val.name,
    value: val.id,
  }));

  const onSubmit = async (payload: FormInformationSchemaType) => {
    try {
      const response = await updateInformationStudent(payload);

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

  const handleSetFormValue = useCallback(() => {
    setValue("private_email", information?.data?.private_email);
    setValue("phone", information?.data?.phone);
    setValue("college_email", information?.data?.college_email as string);
    setValue("almamater_size_id", {
      label: information?.data?.almamater_size_name as string,
      value: information?.data?.almamater_size_id,
    });
    setValue("citizenship_id", {
      label: information?.data?.citizenship_name as string,
      value: information?.data?.citizenship_id,
    });
    setValue("job_id", {
      label: information?.data?.job_name as string,
      value: information?.data?.job_id,
    });
    setValue("transportation_id", {
      label: information?.data?.transportation_name as string,
      value: information?.data?.transportation_id,
    });
  }, [information, setValue]);

  useEffect(() => {
    if (information) handleSetFormValue();
  }, [information, handleSetFormValue]);

  if (information.error) {
    return <h1>{information.message}</h1>;
  }
  return (
    <>
      <div className="d-flex border-2 border-bottom justify-content-between align-items-center">
        <h5 className="card-title py-3 mb-1" style={{ fontWeight: "500" }}>
          Informasi
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
      <form onSubmit={handleSubmit((data) => onSubmit(data))} className="my-2">
        <Row className="gap-1 gap-lg-0">
          {/* left section */}
          <Col md={12} lg={6}>
            <Row className="gap-2">
              {/* study_program_id*/}
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
                      Program Studi
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="study_program_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={[]}
                          // isLoading={isLoadingCity}
                          isDisabled
                          placeholder="Pilih Program Studi"
                          isError={!!errors.study_program_id}
                          id={"study_program_id"}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.study_program_id} />
                  </Col>
                </Row>
              </Col>

              {/* college_email*/}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="college_email"
                      className="form-label mb-0 fw-medium"
                    >
                      Email Kampus
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="college_email"
                      control={control}
                      rules={{
                        pattern: {
                          value:
                            /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/,
                          message: "Format email tidak valid",
                        },
                      }}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.college_email ? "border border-danger" : ""
                          }`}
                          id="college_email"
                          placeholder="Masukkan Email Kampus"
                          disabled
                          type="email"
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.college_email} />
                    <FormDescription message="Emailkampus@unsia.ac.id" />
                  </Col>
                </Row>
              </Col>

              {/* private_email*/}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="private_email"
                      className="form-label mb-0 fw-medium"
                    >
                      Email Pribadi
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="private_email"
                      control={control}
                      rules={{
                        pattern: {
                          value:
                            /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$/,
                          message: "Format email tidak valid",
                        },
                      }}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.private_email ? "border border-danger" : ""
                          }`}
                          id="private_email"
                          placeholder="Masukkan Email pribadi"
                          disabled={!isEdit}
                          {...field}
                        />
                      )}
                    />
                    {errors.private_email ? (
                      <FormErrorMessage errors={errors.private_email} />
                    ) : (
                      <FormDescription message="Emailpribadi@gmail.com" />
                    )}
                  </Col>
                </Row>
              </Col>

              {/* phone */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="phone"
                      className="form-label mb-0 fw-medium"
                    >
                      No. HP
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="phone"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.phone ? "border border-danger" : ""
                          }`}
                          id="phone"
                          placeholder="Masukkan No. HP"
                          disabled={!isEdit}
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);

                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.phone} />
                  </Col>
                </Row>
              </Col>
            </Row>
          </Col>
          {/* right section */}
          <Col md={12} lg={6}>
            <Row className="gap-2">
              {/* transportation_id */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="transportation_id"
                      className="form-label mb-0 fw-medium"
                    >
                      Transportasi
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="transportation_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={transportationOptions as OptionType[]}
                          isLoading={isLoadingTransportation}
                          isDisabled={!isEdit}
                          placeholder="Pilih Jenis Transportasi"
                          isError={!!errors.transportation_id}
                          id={"transportation_id"}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.transportation_id} />
                  </Col>
                </Row>
              </Col>

              {/* citizenship_id*/}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="citizenship_id"
                      className="form-label mb-0 fw-medium"
                    >
                      Kewarganegaraan
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="citizenship_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={countryOptions as OptionType[]}
                          isLoading={isLoadingCountry}
                          isDisabled={!isEdit}
                          placeholder="Pilih Kewarganegaraan"
                          isError={!!errors.citizenship_id}
                          id={"citizenship_id"}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.citizenship_id} />
                  </Col>
                </Row>
              </Col>

              {/* almamater_size_id */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="almamater_size_id"
                      className="form-label mb-0 fw-medium"
                    >
                      Ukuran Jas Almamater
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="almamater_size_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={almamaterOptions as OptionType[]}
                          isLoading={isLoadingAlmamaterSize}
                          isDisabled={!isEdit}
                          placeholder="Pilih Ukuran Jas Almamater"
                          isError={!!errors.almamater_size_id}
                          id={"almamater_size_id"}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.almamater_size_id} />
                  </Col>
                </Row>
              </Col>

              {/* job_id */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="job_id"
                      className="form-label mb-0 fw-medium"
                    >
                      Pekerjaan
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="job_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={jobOptions as OptionType[]}
                          isLoading={isLoadingJob}
                          isDisabled={!isEdit}
                          placeholder="Pilih Pekerjaan"
                          isError={!!errors.job_id}
                          id={"job_id"}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.job_id} />
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
            >
              <span>Batal</span>
            </button>
            <Button
              color="primary"
              className="d-flex flex-grow-0 justify-content-center align-items-center"
            >
              <span>Update</span>
            </Button>
          </div>
        )}
      </form>
    </>
  );
};
