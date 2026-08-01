import { z } from "zod";

export const FormInformationSchema = z.object({
  //   program studi
  study_program_id: z
    .object({
      label: z.string(),
      value: z.string(),
    })
    .optional(),

  //   email kampus
  college_email: z.string().optional(),

  //   email pribadi
  private_email: z
    .string()
    .min(1, "Email pribadi harus diisi")
    .email({ message: "Email tidak valid" }),

  //   no hp
  phone: z
    .string({
      required_error: "No.HP harus diisi", // Pesan kesalahan jika tidak diisi
    })
    .min(11, { message: "No.HP minimil 11 angka" })
    .max(13, "No.HP tidak boleh lebih dari 13 angka"),

  //   transportasi
  transportation_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Jenis transportasi harus diisi",
    }
  ),

  //   kewarganegaraan
  citizenship_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Kewarganegaraan harus diisi",
    }
  ),

  //   pekerjaan
  job_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Pekerjaan harus diisi",
    }
  ),

  //   ukuran jas
  almamater_size_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Ukuran jas almamater harus diisi",
    }
  ),
});

export type FormInformationSchemaType = z.infer<typeof FormInformationSchema>;
