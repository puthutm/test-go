"use client";

import { Button, Col, Input, Label, Row } from "reactstrap";
import { useCallback, useEffect, useState } from "react";
import { useForm, Controller } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import { SelectComponent } from "@/components/ui/select";

import {
  FormAddressSchema,
  FormAddressSchemaType,
} from "@/lib/validations/students/biodata/form-address-schema";
import { FormErrorMessage } from "@/components/ui/form-error-message";
import { handleInputNumberOnly } from "@/lib/utils/input-number-only";
import { EditIcon } from "@/components/icons/edit";
import { FormDescription } from "@/components/ui/form-description";
import { useCountries } from "@/services/api/data-referensi/country/use-get-countries";
import { useProvincesByCountryId } from "@/services/api/data-referensi/province/use-get-province-by-country-id";
import { useCitiesByProvinceId } from "@/services/api/data-referensi/city/use-get-cities-by-province-id";
import { useDistrictsByCityId } from "@/services/api/data-referensi/district/use-get-district-by-city-id";
import { useVillagesByDistrictId } from "@/services/api/data-referensi/village/use-get-village-by-district-id";
import { useModalConfirmationContext } from "@/lib/hooks/use-modal-confirmation";
import { updateAddressStudent } from "@/services/api/students/biodata/address/update-address";

export const FormAddress = ({
  address,
}: {
  address: ApiResponse<AddressStudent>;
}) => {
  const [isEdit, setIsEdit] = useState<boolean>(false);
  const { setModalConfirmationState } = useModalConfirmationContext();

  const {
    formState: { errors },
    handleSubmit,
    control,
    reset,
    watch,
    setValue,
  } = useForm<FormAddressSchemaType>({
    resolver: zodResolver(FormAddressSchema),
    defaultValues: {
      address: "",
      rt: "",
      rw: "",
      postal_code: "",
      distance: "",
    },
  });

  const { data: dataCountries, isLoading: isLoadingCountries } = useCountries();

  const countryOptions = dataCountries?.data?.map((country) => ({
    label: country.name,
    value: country.id,
  }));

  const selectedCountry = watch("country_id")?.value;

  const { data: dataProvinces, isLoading: isLoadingProvinces } =
    useProvincesByCountryId(selectedCountry as string);

  const provinceOptions = dataProvinces?.data?.map((province) => ({
    label: province.name,
    value: province.id,
  }));

  const selectedProvince = watch("province_id")?.value;

  const { data: dataCities, isLoading: isLoadingCities } =
    useCitiesByProvinceId(selectedProvince as string);

  const cityOptions = dataCities?.data?.map((city) => ({
    label: city.name,
    value: city.id,
  }));

  const selectedCity = watch("city_id")?.value;

  const { data: dataDistricts, isLoading: isLoadingDistricts } =
    useDistrictsByCityId(selectedCity as string);

  const districtOptions = dataDistricts?.data?.map((district) => ({
    label: district.name,
    value: district.id,
  }));

  const selectedDistrict = watch("district_id")?.value;

  const { data: dataVillages, isLoading: isLoadingVillages } =
    useVillagesByDistrictId(selectedDistrict as string);

  const villageOptions = dataVillages?.data?.map((village) => ({
    label: village.name,
    value: village.id,
  }));

  const onSubmit = async (payload: FormAddressSchemaType) => {
    try {
      const response = await updateAddressStudent(payload);

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
    setValue("address", address?.data?.address || "");
    setValue("distance", address?.data?.distance || "");
    setValue("rt", address?.data?.rt || "");
    setValue("rw", address?.data?.rw || "");
    setValue("postal_code", address?.data?.postal_code || "");
    if (address?.data?.country_id) {
      setValue("country_id", {
        label: address?.data?.country_name as string,
        value: address?.data?.country_id as string,
      });
    } else {
      setValue("country_id", null);
    }
    if (address?.data?.province_id) {
      setValue("province_id", {
        label: address?.data?.province_name as string,
        value: address?.data?.province_id as string,
      });
    } else {
      setValue("province_id", null);
    }
    if (address?.data?.city_id) {
      setValue("city_id", {
        label: address?.data?.city_name as string,
        value: address?.data?.city_id as string,
      });
    } else {
      setValue("city_id", null);
    }
    if (address?.data?.district_id) {
      setValue("district_id", {
        label: address?.data?.district_name as string,
        value: address?.data?.district_id as string,
      });
    } else {
      setValue("district_id", null);
    }
    if (address?.data?.village_id) {
      setValue("village_id", {
        label: address?.data?.village_name as string,
        value: address?.data?.village_id as string,
      });
    } else {
      setValue("village_id", null);
    }
  }, [address, setValue]);

  useEffect(() => {
    if (address) handleSetFormValue();
  }, [address, handleSetFormValue]);

  if (address.error) {
    return <h1>{address.message}</h1>;
  }

  return (
    <>
      <div className="d-flex border-2 border-bottom justify-content-between align-items-center">
        <h5 className="card-title py-3 mb-1" style={{ fontWeight: "500" }}>
          Alamat
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
              {/* country */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="country_id"
                      className="form-label mb-0 fw-medium"
                    >
                      Negara
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="country_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={countryOptions as OptionType[]}
                          isLoading={isLoadingCountries}
                          isDisabled={!isEdit}
                          placeholder="Pilih Negara"
                          isError={!!errors.country_id}
                          id={"country_id"}
                          {...field}
                          onChange={(e) => {
                            field.onChange(e);
                            setValue("province_id", null);
                            setValue("city_id", null);
                            setValue("district_id", null);
                            setValue("village_id", null);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.country_id} />
                  </Col>
                </Row>
              </Col>
              {/* province */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="province_id"
                      className="form-label mb-0 fw-medium"
                    >
                      Provinsi
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="province_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={provinceOptions as OptionType[]}
                          isLoading={isLoadingProvinces}
                          isDisabled={!isEdit || !selectedCountry}
                          placeholder="Pilih Provinsi"
                          isError={!!errors.province_id}
                          id={"province_id"}
                          {...field}
                          onChange={(e) => {
                            field.onChange(e);
                            setValue("city_id", null);
                            setValue("district_id", null);
                            setValue("village_id", null);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.province_id} />
                  </Col>
                </Row>
              </Col>
              {/* City*/}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="city_id"
                      className="form-label mb-0 fw-medium"
                    >
                      Kabupaten / Kota
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="city_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={cityOptions as OptionType[]}
                          isLoading={isLoadingCities}
                          isDisabled={!isEdit || !selectedProvince}
                          placeholder="Pilih Kabupaten / Kota"
                          isError={!!errors.city_id}
                          id={"city_id"}
                          {...field}
                          onChange={(e) => {
                            field.onChange(e);
                            setValue("district_id", null);
                            setValue("village_id", null);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.city_id} />
                  </Col>
                </Row>
              </Col>
              {/* District*/}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="district_id"
                      className="form-label mb-0 fw-medium"
                    >
                      Kecamatan
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="district_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={districtOptions as OptionType[]}
                          isLoading={isLoadingDistricts}
                          isDisabled={!isEdit || !selectedCity}
                          placeholder="Pilih Kecamatan"
                          isError={!!errors.district_id}
                          id={"district_id"}
                          {...field}
                          onChange={(e) => {
                            field.onChange(e);
                            setValue("village_id", null);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.district_id} />
                  </Col>
                </Row>
              </Col>
              {/* Village */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="village_id"
                      className="form-label mb-0 fw-medium"
                    >
                      Kelurahan / Desa
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="village_id"
                      control={control}
                      render={({ field }) => (
                        <SelectComponent
                          options={villageOptions as OptionType[]}
                          isLoading={isLoadingVillages}
                          isDisabled={!isEdit || !selectedDistrict}
                          placeholder="Pilih Kelurahan / Desa"
                          isError={!!errors.village_id}
                          id={"village_id"}
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.village_id} />
                  </Col>
                </Row>
              </Col>
              {/* rt*/}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label htmlFor="rt" className="form-label mb-0 fw-medium">
                      RT
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="rt"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.rt ? "border border-danger" : ""
                          }`}
                          id="rt"
                          placeholder="Masukkan RT"
                          disabled={!isEdit}
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);

                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    {errors.rt ? (
                      <FormErrorMessage errors={errors.rt} />
                    ) : (
                      <FormDescription message="Contoh : 010. Jika tidak memiliki RT mohon isi 00" />
                    )}
                  </Col>
                </Row>
              </Col>
              {/* rw */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label htmlFor="rw" className="form-label mb-0 fw-medium">
                      RW
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="rw"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.rw ? "border border-danger" : ""
                          }`}
                          id="rw"
                          placeholder="Masukkan RW"
                          disabled={!isEdit}
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);

                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    {errors.rw ? (
                      <FormErrorMessage errors={errors.rw} />
                    ) : (
                      <FormDescription message="Contoh : 010. Jika tidak memiliki RW mohon isi 00" />
                    )}
                  </Col>
                </Row>
              </Col>
            </Row>
          </Col>
          {/* right section */}
          <Col md={12} lg={6}>
            <Row className="gap-2">
              {/* alamat */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="address"
                      className="form-label mb-0 fw-medium"
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
                          placeholder="Masukkan Alamat"
                          disabled={!isEdit}
                          rows={5}
                          style={{ resize: "none" }} // Matikan resize
                          {...field}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.address} />
                  </Col>
                </Row>
              </Col>

              {/* kode_pos */}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="postal_code"
                      className="form-label mb-0 fw-medium"
                    >
                      Kode Pos
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="postal_code"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.postal_code ? "border border-danger" : ""
                          }`}
                          id="postal_code"
                          placeholder="Masukkan Kode Pos"
                          disabled={!isEdit}
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);

                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.postal_code} />
                  </Col>
                </Row>
              </Col>

              {/* jarak_rumah*/}
              <Col sm={12}>
                <Row
                  className="align-items-center gap-2"
                  style={{ paddingBottom: "10px" }}
                >
                  <Col sm={12}>
                    <Label
                      htmlFor="distance"
                      className="form-label mb-0 fw-medium optional"
                    >
                      Jarak Rumah
                    </Label>
                  </Col>
                  <Col sm={12}>
                    <Controller
                      name="distance"
                      control={control}
                      render={({ field }) => (
                        <Input
                          className={`form-control form-control-icon ${
                            errors.distance ? "border border-danger" : ""
                          }`}
                          id="distance"
                          placeholder="Masukkan Jarak Rumah"
                          disabled={!isEdit}
                          {...field}
                          onChange={(e) => {
                            const { stringValue } = handleInputNumberOnly(e);

                            field.onChange(stringValue);
                          }}
                        />
                      )}
                    />
                    <FormErrorMessage errors={errors.distance} />
                    <FormDescription message="dalam km" />
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
                reset();
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
