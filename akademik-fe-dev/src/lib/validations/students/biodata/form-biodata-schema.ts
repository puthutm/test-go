import { z } from "zod";

import { NIK_REGEX } from "@/lib/constants/nik-regex";

export const FormBiodataSchema = z.object({
  nim: z.string().optional(),
  name: z
    .string({ required_error: "Nama lengkap harus diisi" })
    .min(1, { message: "Nama lengkap harus diisi" }),
  back_degree: z.string().optional(),
  birth_place_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Tempat lahir harus diisi",
    }
  ),
  birth_date: z
    .array(z.date(), {
      message: "Tanggal lahir harus diisi",
    })
    .min(1, {
      message: "Tanggal lahir harus diisi",
    }),
  gender: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Jenis kelamin harus diisi",
    }
  ),
  religion_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Agama harus diisi",
    }
  ),
  ethnic_id: z
    .object({
      label: z.string(),
      value: z.string(),
    })
    .optional(),
  status_id: z
    .object({
      label: z.string(),
      value: z.string(),
    })
    .optional(),
  nik: z
    .string({ required_error: "NIK harus diisi" })
    .min(16, { message: "NIK harus 16 angka" })
    .regex(NIK_REGEX, { message: "NIK tidak valid" }),
  no_kk: z
    .string({
      required_error: "Nomor KK harus diisi",
    })
    .min(16, {
      message: "Nomor KK harus 16 angka",
    })
    .regex(NIK_REGEX, { message: "Nomor KK tidak valid" }),
  height: z.string().optional(),
  weight: z.string().optional(),
  blood_type_id: z
    .object({
      label: z.string(),
      value: z.string(),
    })
    .optional(),
});

export type FormBiodataSchemaType = z.infer<typeof FormBiodataSchema>;
