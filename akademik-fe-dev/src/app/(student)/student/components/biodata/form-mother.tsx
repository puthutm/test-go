"use client";

import { Button, Col, Input, Label, Row } from "reactstrap";
import { zodResolver } from "@hookform/resolvers/zod";
import { useCallback, useEffect, useState } from "react";
import { Controller, useForm } from "react-hook-form";

import { FormErrorMessage } from "@/components/ui/form-error-message";
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import { SelectComponent } from "@/components/ui/select";
import { FormDescription } from "@/components/ui/form-description";
import { DatePicker } from "@/components/ui/date-picker";
import { LIFE_STATUS_OPTIONS } from "@/lib/constants/life-status";
import { EditIcon } from "@/components/icons/edit";
import {
  FormParentSchema,
  FormParentSchemaType,
} from "@/lib/validations/students/biodata/form-parent-schema";
import { STATUS_KINSHIP_OPTIONS } from "@/lib/constants/status-kinship";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { updateParent } from "@/services/api/students/biodata/parent/update-parent";
import { CalendarTodayIcon } from "@/components/icons/calendar-today";
import { useEducationalLevels } from "@/services/api/data-referensi/educational-level/use-get-educational-level";
import { useCities } from "@/services/api/data-referensi/city/use-get-cities";
import { useJobs } from "@/services/api/data-referensi/job/use-get-jobs";

export const FormMother = ({
  mother,
}: {
  mother: ApiResponse<ParentStudent>;
}) => {
  const [isEdit, setIsEdit] = useState<boolean>(false);
  const { setModalConfirmationState } = useModalConfirmationContext();

  const { data: educationalLevels, isLoading: isLoadingEducationalLevels } =
    useEducationalLevels();

  const educationalLevelOptions = educationalLevels?.data?.map((edu) => ({
    label: edu.name,
    value: edu.id,
  })) as OptionType[];

  const { data: cities, isLoading: isLoadingCities } = useCities();

  const cityOptions = cities?.data?.map((city) => ({
    label: city.name,
    value: city.id,
  })) as OptionType[];

  const { data: jobs, isLoading: isLoadingJob } = useJobs();

  const jobOptions = jobs?.data?.map((job) => ({
    label: job.name,
    value: job.id,
  })) as OptionType[];

  const {
    formState: { errors },
    control,
    handleSubmit,
    setValue,
    clearErrors,
  } = useForm<FormParentSchemaType>({
    resolver: zodResolver(FormParentSchema),
    defaultValues: {
      address: "",
      birth_date: [],
      name: "",
      nik: "",
      phone: "",
      phone2: "",
      income: "",
    },
  });

  const handleSetFormValue = useCallback(() => {
    setValue("name", mother?.data.name);
    setValue("nik", mother?.data.nik);
    setValue("email", (mother?.data?.email as string) || undefined);
    setValue("phone", (mother?.data?.phone as string) || "");
    setValue("phone2", (mother?.data?.phone2 as string) || "");
    if (mother?.data.educational_level_id) {
      setValue("education_level_id", {
        label: mother?.data.educational_level_name as string,
        value: mother?.data?.educational_level_id,
      });
    }
    if (mother?.data.job_id) {
      setValue("job_id", {
        label: mother?.data.job_name as string,
        value: mother?.data?.job_id,
      });
    }
    if (mother?.data.birth_place_id) {
      setValue("birth_place_id", {
        label: mother?.data.birth_place_name as string,
        value: mother?.data?.birth_place_id,
      });
    }
    if (mother?.data.status_kinship) {
      const selectedStatusKinship = STATUS_KINSHIP_OPTIONS?.find(
        (kinship) => kinship.value === mother?.data.status_kinship
      );

      setValue("status_kinship", {
        label: selectedStatusKinship?.label as string,
        value: selectedStatusKinship?.value as string,
      });
    }
    if (mother?.data.life_status) {
      const selectedStatusKinship = LIFE_STATUS_OPTIONS?.find(
        (status) => status.value === mother?.data.life_status
      );

      setValue("life_status", {
        label: selectedStatusKinship?.label as string,
        value: selectedStatusKinship?.value as string,
      });
    }
    setValue("birth_date", [new Date(mother?.data?.birth_date as string)]);
    setValue("address", (mother?.data?.address as string) || "");
    setValue("income", (mother?.data?.income?.toString() as string) || "");
  }, [mother, setValue]);

  const onSubmit = async (payload: FormParentSchemaType) => {
    try {
      const response = await updateParent({ parentType: "mother", payload });

      if (response.status === 200) {
        setIsEdit(false);
        setModalConfirmationState((prev) => ({
          ...prev,
          open: true,
          message: "Data berhasil di-update",
          state: "success",
        }));
      }

      return response;
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
    if (mother) {
      handleSetFormValue();
    }
  }, [mother, handleSetFormValue]);

  if (mother?.error) {
    return <h1>{mother.message}</h1>;
  }
  return (
    <>
      <div className="d-flex border-2 border-bottom justify-content-between align-items-center">
        <h5 className="card-title py-3 mb-1" style={{ fontWeight: "500" }}>
          Ibu
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
      <form className="my-2" onSubmit={handleSubmit(onSubmit)}>
        <Row className="gap-1 gap-lg-0">
          {/* left section */}
          <Col md={12} lg={6}>
            <Row className="gap-2">
              {/* name */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2 "
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label htmlFor="name" className="form-label mb-0 fw-medium">
                      Nama
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
                          placeholder="Masukkan Nama"
                          disabled={!isEdit}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.name} />
                  </Col>
                </Row>
              </Col>
              {/* nik */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label htmlFor="nik" className="form-label mb-0 fw-medium">
                      NIK
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="nik"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.nik ? "border border-danger" : ""
                          }`}
                          id="nik"
                          placeholder="Masukkan Nama Induk Keluarga"
                          disabled={!isEdit}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.nik} />
                  </Col>
                </Row>
              </Col>
              {/* educational level */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="education_level_id"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Pendidikan
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="education_level_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={educationalLevelOptions}
                          id={"education_level_id"}
                          placeholder="Pendidikan Terakhir"
                          isDisabled={!isEdit}
                          isError={!!errors.education_level_id}
                          isLoading={isLoadingEducationalLevels}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.education_level_id} />
                  </Col>
                </Row>
              </Col>
              {/* birth place */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="birth_place_id"
                      className="form-label mb-0 fw-medium"
                    >
                      Tempat Lahir
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="birth_place_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={cityOptions}
                          isLoading={isLoadingCities}
                          isDisabled={!isEdit}
                          placeholder="Pilih Tempat Lahir"
                          isError={!!errors.birth_place_id}
                          id={"birth_place_id"}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.birth_place_id} />
                  </Col>
                </Row>
              </Col>
              {/* birth date */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="number_of_study_program"
                      className="form-label mb-0 fw-medium"
                    >
                      Tanggal Lahir
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <div className="form-icon">
                      <Controller
                        name="birth_date"
                        control={control}
                        render={({ field }) => {
                          return (
                            <DatePicker
                              onChange={(e) => field.onChange(e)}
                              value={field.value}
                              className={`p-0 ${
                                errors.birth_date ? "border border-danger" : ""
                              }`}
                              classNameFlatpickr={`form-control form-control-icon disabled-input`}
                              options={{
                                mode: "single",
                                dateFormat: "d F Y",
                              }}
                              disabled={!isEdit}
                            />
                          );
                        }}
                      />
                      <i style={{ left: "15px" }}>
                        <CalendarTodayIcon color="#878A99" />
                      </i>
                    </div>
                    <FormErrorMessage errors={errors.birth_date} />
                  </Col>
                </Row>
              </Col>
              {/* job */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="job_id"
                      className="form-label mb-0 fw-medium optional"
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
                          options={jobOptions}
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
              {/* income */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="income"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Penghasilan
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="income"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.income ? "border border-danger" : ""
                          }`}
                          id="income"
                          placeholder="Penghasilan"
                          disabled={!isEdit}
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);

                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.income} />
                  </Col>
                </Row>
              </Col>
            </Row>
          </Col>
          {/* right section */}
          <Col md={12} lg={6}>
            <Row className="gap-2">
              {/* Email */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="email"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Email
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="email"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.email ? "border border-danger" : ""
                          }`}
                          id="email"
                          placeholder="Masukan Email"
                          disabled={!isEdit}
                          {...field}
                          value={field.value ? field.value : ""}
                        />
                      )}
                    />
                    {errors.email ? (
                      <FormErrorMessage errors={errors.email} />
                    ) : (
                      <FormDescription message="Email@gmail.com" />
                    )}
                  </Col>
                </Row>
              </Col>
              {/* phone*/}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="phone"
                      className="form-label mb-0 fw-medium optional"
                    >
                      No. Hp
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
                          placeholder="Masukkan No Hp"
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
              {/* phone2 */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="phone2"
                      className="form-label mb-0 fw-medium optional"
                    >
                      No Telephone
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="phone2"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.phone2 ? "border border-danger" : ""
                          }`}
                          id="phone2"
                          placeholder="Masukkan No Telephone"
                          disabled={!isEdit}
                          {...field}
                          onChange={(e) => {
                            const { numberValue } = handleInputNumberOnly(e);

                            field.onChange(numberValue);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.phone2} />
                  </Col>
                </Row>
              </Col>
              {/* address */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="address"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Alamat
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="address"
                      control={control}
                      render={({ field }) => (
                        <textarea
                          className={`form-control form-control-icon ${
                            errors.address ? "border border-danger" : ""
                          }`}
                          id="address"
                          placeholder="Text"
                          disabled={!isEdit}
                          {...field}
                          style={{ height: "120px" }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.address} />
                  </Col>
                </Row>
              </Col>
              {/* life status */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="life_status"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Status Hidup
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="life_status"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={LIFE_STATUS_OPTIONS}
                          isDisabled={!isEdit}
                          placeholder="Pilih Golongan Darah"
                          id={"life_status"}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.life_status} />
                  </Col>
                </Row>
              </Col>
              {/* status kinship */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="status_kinship"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Status Kekerabatan
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="status_kinship"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={STATUS_KINSHIP_OPTIONS}
                          // isLoading={isLoadingBloodType}
                          isDisabled={!isEdit}
                          placeholder="Pilih Golongan Darah"
                          id={"status_kinship"}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.status_kinship} />
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
                setIsEdit(false);
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
