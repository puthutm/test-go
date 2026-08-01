import { z } from "zod";

export const FormAddressSchema = z
  .object({
    country_id: z
      .object(
        {
          label: z.string().nullable(),
          value: z.string().nullable(),
        },
        {
          message: "Negara harus diisi",
        }
      )
      .nullish(),
    province_id: z
      .object(
        {
          label: z.string().nullable(),
          value: z.string().nullable(),
        },
        {
          message: "Provinsi harus diisi",
        }
      )
      .nullish(),
    city_id: z
      .object(
        {
          label: z.string().nullable(),
          value: z.string().nullable(),
        },
        {
          message: "Kabupaten/Kota harus diisi",
        }
      )
      .nullish(),
    district_id: z
      .object(
        {
          label: z.string().nullable(),
          value: z.string().nullable(),
        },
        {
          message: "Kecamatan harus diisi",
        }
      )
      .nullish(),
    village_id: z
      .object(
        {
          label: z.string().nullable(),
          value: z.string().nullable(),
        },
        {
          message: "Kelurahan/Desa harus diisi",
        }
      )
      .nullish(),
    rt: z.string().min(1, { message: "RT harus diisi" }),
    rw: z.string().min(1, { message: "RW harus diisi" }),
    address: z.string().min(1, "Alamat harus diisi"),
    postal_code: z
      .string()
      .min(1, { message: "Kode pos harus diisi" })
      .max(5, { message: "Kode pos maksimal 5 angka" }),
    distance: z.string().optional(),
  })
  .refine(
    (data) => !data.country_id || (data.province_id && data.province_id.value),
    {
      message: "Provinsi harus diisi",
      path: ["province_id"],
    }
  )
  .refine((data) => !data.province_id || (data.city_id && data.city_id.value), {
    message: "Kabupaten/Kota harus diisi",
    path: ["city_id"],
  })
  .refine(
    (data) => !data.city_id || (data.district_id && data.district_id.value),
    {
      message: "Kecamatan harus diisi",
      path: ["district_id"],
    }
  )
  .refine(
    (data) => !data.district_id || (data.village_id && data.village_id.value),
    {
      message: "Kelurahan/Desa harus diisi",
      path: ["village_id"],
    }
  );

export type FormAddressSchemaType = z.infer<typeof FormAddressSchema>;
