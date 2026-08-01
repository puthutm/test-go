import { z } from "zod";

import { NIK_REGEX } from "@/lib/constants/nik-regex";

export const FormParentSchema = z.object({
  name: z.string().min(1, { message: "Nama orang tua harus diisi" }),
  email: z.string().email("Email tidak valid").optional(),
  nik: z
    .string()
    .min(16, "NIK harus 16 digit")
    .max(16, "NIK harus 16 digit")
    .regex(NIK_REGEX, { message: "NIK tidak valid" })
    .optional(),
  phone: z.string().optional(),
  phone2: z.string().optional(),
  education_level_id: z
    .object({
      label: z.string(),
      value: z.string(),
    })
    .optional(),
  job_id: z
    .object({
      label: z.string(),
      value: z.string(),
    })
    .optional(),
  birth_place_id: z
    .object({
      label: z.string(),
      value: z.string(),
    })
    .optional(),
  status_kinship: z
    .object({
      label: z.string(),
      value: z.string(),
    })
    .optional(),
  life_status: z
    .object({
      label: z.string(),
      value: z.string(),
    })
    .optional(),
  birth_date: z.array(z.date()).optional(),
  address: z.string().optional(),
  income: z.string().optional(),
  kinship: z.string().optional(),
});

export type FormParentSchemaType = z.infer<typeof FormParentSchema>;
