import { z } from "zod";

export const FormSeminarSchema = z.object({


  // Lembar Pengesahan
  lembar_pengesahan: z
    .instanceof(File)
    .refine((file) => file.type === "application/pdf", {
      message: "File harus berformat .pdf",
    })
    .refine((file) => file.size <= 10 * 1024 * 1024, {
      message: "Ukuran file maksimal 10MB",
    })
    .optional(),

  // Kartu Konsultasi
  kartu_konsultasi: z
    .instanceof(File)
    .refine((file) => file.type === "application/pdf", {
      message: "File harus berformat .pdf",
    })
    .refine((file) => file.size <= 10 * 1024 * 1024, {
      message: "Ukuran file maksimal 10MB",
    })
    .optional(),


  // Lembar Penilaian
  lembar_penilaian: z
    .instanceof(File)
    .refine((file) => file.type === "application/pdf", {
      message: "File harus berformat .pdf",
    })
    .refine((file) => file.size <= 10 * 1024 * 1024, {
      message: "Ukuran file maksimal 10MB",
    })
    .optional(),




});

export type FormSeminarSchemaType = z.infer<typeof FormSeminarSchema>;
