import { z } from "zod";

export const FormOriginalEducationSchema = z.object({
  institution_name: z
    .string({
      required_error: "Asal sekolah atau perguruan tinggi harus diisi",
    })
    .min(1, { message: "Asal sekolah atau perguruan tinggi harus diisi" }),
  school_major: z
    .string({ required_error: "Jurusuan atau program studi harus diisi" })
    .min(1, { message: "Jurusuan atau program studi harus diisi" }),
  nisn: z
    .string({ required_error: "NIM atau NISN harus diisi" })
    .min(1, { message: "NIM atau NISN harus diisi" }),
  national_exam_score: z
    .string({ required_error: "IPK atau nilai rata-rata harus diisi" })
    .min(1, { message: "IPK atau nilai rata-rata harus diisi" }),
  certificate_number: z
    .string({ required_error: "Nomor ijazah harus diisi" })
    .min(1, { message: "Nomor ijazah harus diisi" }),
  certificate_filepath: z
    .instanceof(File, { message: "File ijazah harus diunggah" })
    .refine(
      (file) => {
        if (!file) return true;
        return file.size <= 2 * 1024 * 1024; // 2mb max
      },
      { message: "File maksimal 2mb" }
    )
    .refine(
      (file) => {
        if (!file) return true;
        return file.type === "application/pdf";
      },
      { message: "File dengan format pdf" }
    ),
  transcripts_filepath: z
    .instanceof(File, { message: "File transkrip nilai harus diunggah" })
    .refine(
      (file) => {
        if (!file) return true;
        return file.size <= 2 * 1024 * 1024; // 2mb max
      },
      { message: "File maksimal 2mb" }
    )
    .refine(
      (file) => {
        if (!file) return true;
        return file.type === "application/pdf";
      },
      { message: "File dengan format pdf" }
    ),
});

export type FormOriginalEducationSchemaType = z.infer<
  typeof FormOriginalEducationSchema
>;
