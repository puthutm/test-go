"use client";

import { Button, Col, Input, Label, Row } from "reactstrap";
import { useCallback, useEffect, useState } from "react";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { debounce } from "lodash";

import { SelectComponent } from "@/components/ui/select";
import { CalendarTodayIcon } from "@/components/icons/calendar-today";
import { DatePicker } from "@/components/ui/date-picker";
import {
  FormBiodataSchema,
  FormBiodataSchemaType,
} from "@/lib/validations/students/biodata/form-biodata-schema";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import { EditIcon } from "@/components/icons/edit";
import { FormDescription } from "@/components/ui/form-description";
import { useEthnics } from "@/services/api/data-referensi/ethnic/use-get-ethnics";
import { useReligions } from "@/services/api/data-referensi/religion/use-get-religions";
import { useStudentStatuses } from "@/services/api/data-referensi/student-status/use-get-student-statuses";
import { useBloodTypes } from "@/services/api/data-referensi/blood-type/use-get-blood-types";
import { useCities } from "@/services/api/data-referensi/city/use-get-cities";
import { GENDER_OPTIONS } from "@/lib/constants/gender";
import { updateBiodataStudent } from "@/services/api/students/biodata/biodata/update-biodata";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";

export const FormBiodata = ({
  biodata,
}: {
  biodata: ApiResponse<BiodataStudent>;
}) => {
  const [isEdit, setIsEdit] = useState<boolean>(false);
  const [search, setSearch] = useState("");
  const { setModalConfirmationState } = useModalConfirmationContext();

  const { data: dataEthnic, isLoading: isLoadingEthnic } = useEthnics();
  const { data: dataReligion, isLoading: isLoadingReligion } = useReligions();
  const { data: dataStudentStatus, isLoading: isLoadingStudentStatus } =
    useStudentStatuses();
  const { data: dataBloodTypes, isLoading: isLoadingBloodTypes } =
    useBloodTypes();
  const { data: dataCities, isLoading: isLoadingCities } = useCities(search);

  const ethnicOptions = dataEthnic?.data?.map((val: Ethnic) => ({
    label: val.name,
    value: val.id,
  }));

  const religionOptions = dataReligion?.data?.map((val: Religion) => ({
    label: val.name,
    value: val.id,
  }));

  const cityOptions = dataCities?.data?.map((val: City) => ({
    label: val.name,
    value: val.id,
  }));

  const studentStatusOptions = dataStudentStatus?.data?.map(
    (val: Omit<StudentStatus, "is_default">) => ({
      label: val.name,
      value: val.id,
    })
  );

  const bloodTypeOptions = dataBloodTypes?.data?.map(
    (val: Omit<BloodType, "code">) => ({
      label: val.name,
      value: val.id,
    })
  );

  const handleSearch = debounce((e) => {
    setSearch(e);
  }, 1000);

  const {
    formState: { errors, isSubmitting },
    handleSubmit,
    control,
    clearErrors,
    setValue,
  } = useForm<FormBiodataSchemaType>({
    resolver: zodResolver(FormBiodataSchema),
    defaultValues: {
      nim: "",
      name: "",
      back_degree: "",
      height: "",
      weight: "",
      no_kk: "",
      nik: "",
      birth_date: [],
    },
  });

  const handleSetFormValue = useCallback(() => {
    const selectedGender = GENDER_OPTIONS.find(
      (val) => val.value === biodata?.data?.gender
    );

    setValue("name", (biodata?.data?.name as string) || "");
    setValue("back_degree", (biodata?.data?.back_degree as string) || "");
    setValue("gender", selectedGender as OptionType);
    setValue("nik", (biodata?.data?.nik as string) || "");
    setValue("no_kk", (biodata?.data?.no_kk as string) || "");
    if (biodata?.data?.birth_place_id) {
      setValue("birth_place_id", {
        label: biodata?.data?.birth_place_name as string,
        value: biodata?.data?.birth_place_id as string,
      });
    }
    if (biodata?.data?.ethnic_id) {
      setValue("ethnic_id", {
        label: biodata?.data?.ethnic_name as string,
        value: biodata?.data?.ethnic_id as string,
      });
    }
    if (biodata?.data?.religion_id) {
      setValue("religion_id", {
        label: biodata?.data?.religion_name as string,
        value: biodata?.data?.religion_id as string,
      });
    }

    if (biodata?.data?.status_id) {
      setValue("status_id", {
        label: biodata?.data?.status_name as string,
        value: biodata?.data?.status_id as string,
      });
    }

    if (biodata?.data?.blood_type_id) {
      setValue("blood_type_id", {
        label: biodata?.data?.blood_type_name as string,
        value: biodata?.data?.blood_type_id as string,
      });
    }

    setValue("height", biodata?.data?.height?.toString());
    setValue("weight", biodata?.data?.weight?.toString());
    setValue("birth_date", [new Date(biodata?.data?.birth_date as string)]);
  }, [biodata, setValue]);

  const onSubmit = async (payload: FormBiodataSchemaType) => {
    try {
      const response = await updateBiodataStudent(payload);

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
    if (biodata) handleSetFormValue();
  }, [biodata, handleSetFormValue]);

  if (biodata.error) {
    return <h1>{biodata.message}</h1>;
  }

  return (
    <>
      <div className="d-flex border-2 border-bottom justify-content-between align-items-center">
        <h5 className="card-title py-3 mb-1" style={{ fontWeight: "500" }}>
          Biodata
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
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="my-2"
        autoComplete="off"
      >
        <Row className="gap-1 gap-lg-0">
          {/* left section */}
          <Col md={12} lg={6}>
            <Row className="gap-2">
              {/* nim */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label htmlFor="nim" className="form-label mb-0 fw-medium">
                      NIM
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="nim"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.nim ? "border border-danger" : ""
                          }`}
                          id="nim"
                          placeholder="Masukkan NIM"
                          disabled
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);

                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.nim} />
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
                      Nama Lengkap
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
                          placeholder="Masukkan Nama Lengkap"
                          disabled={!isEdit}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.name} />
                  </Col>
                </Row>
              </Col>
              {/* back title */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="back_degree"
                      className="form-label mb-0 fw-medium"
                    >
                      Gelar Belakang
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="back_degree"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.name ? "border border-danger" : ""
                          }`}
                          id="back_degree"
                          placeholder="Masukkan Gelar Belakang"
                          disabled={!isEdit}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.back_degree} />
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
                      No. NIK
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
                          placeholder="Masukkan NIK"
                          disabled={!isEdit}
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);

                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.nik} />
                  </Col>
                </Row>
              </Col>
              {/* kk */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label htmlFor="kk" className="form-label mb-0 fw-medium">
                      No. KK
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="no_kk"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.no_kk ? "border border-danger" : ""
                          }`}
                          id="no_kk"
                          placeholder="Masukkan No. KK"
                          disabled={!isEdit}
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);

                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.no_kk} />
                  </Col>
                </Row>
              </Col>
              {/* Birthday place */}
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
                          options={cityOptions as OptionType[]}
                          isLoading={isLoadingCities}
                          isDisabled={!isEdit}
                          placeholder="Pilih Tempat Lahir"
                          isError={!!errors.birth_place_id}
                          id={"birth_place_id"}
                          onInputChange={(e) => handleSearch(e)}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.birth_place_id} />
                  </Col>
                </Row>
              </Col>
              {/* Birthday date */}
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
            </Row>
          </Col>
          {/* right section */}
          <Col md={12} lg={6}>
            <Row className="gap-2">
              {/* Gender */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="gender"
                      className="form-label mb-0 fw-medium"
                    >
                      Jenis Kelamin
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="gender"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={GENDER_OPTIONS}
                          // isLoading={isLoadingCity}
                          isDisabled={!isEdit}
                          placeholder="Pilih Jenis Kelamin"
                          isError={!!errors.gender}
                          id={"gender"}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.gender} />
                  </Col>
                </Row>
              </Col>
              {/* Status */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="status_id"
                      className="form-label mb-0 fw-medium"
                    >
                      Status
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="status_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={studentStatusOptions as OptionType[]}
                          isLoading={isLoadingStudentStatus}
                          isDisabled={!isEdit}
                          placeholder="Pilih Status"
                          isError={!!errors.status_id}
                          id={"status_id"}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.status_id} />
                  </Col>
                </Row>
              </Col>
              {/* Religion */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="religion_id"
                      className="form-label mb-0 fw-medium"
                    >
                      Agama
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="religion_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={religionOptions as OptionType[]}
                          isLoading={isLoadingReligion}
                          isDisabled={!isEdit}
                          placeholder="Pilih Agama"
                          isError={!!errors.religion_id}
                          id={"religion_id"}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.religion_id} />
                  </Col>
                </Row>
              </Col>
              {/* Ethnic */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="ethnic_id"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Suku
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="ethnic_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={ethnicOptions as OptionType[]}
                          isLoading={isLoadingEthnic}
                          isDisabled={!isEdit}
                          placeholder="Pilih Suku"
                          id={"ethnic_id"}
                          {...field}
                        />
                      )}
                    />
                  </Col>
                </Row>
              </Col>
              {/* height */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="height"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Tinggi Badan
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="height"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon`}
                          id="height"
                          placeholder="Masukkan Tinggi Badan"
                          disabled={!isEdit}
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);

                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    <FormDescription message="Dalam cm" />
                  </Col>
                </Row>
              </Col>
              {/* weight */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="weight"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Berat Badan
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="weight"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon`}
                          id="weight"
                          placeholder="Masukkan Berat Badan"
                          disabled={!isEdit}
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);

                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    <FormDescription message="Dalam kg" />
                  </Col>
                </Row>
              </Col>
              {/* Blood Type */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="blood_type_id"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Golongan Darah
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="blood_type_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={bloodTypeOptions as OptionType[]}
                          isLoading={isLoadingBloodTypes}
                          isDisabled={!isEdit}
                          placeholder="Pilih Golongan Darah"
                          id={"blood_type_id"}
                          {...field}
                        />
                      )}
                    />
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
