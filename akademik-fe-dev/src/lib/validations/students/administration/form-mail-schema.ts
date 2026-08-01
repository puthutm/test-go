import { z } from "zod";

export const FormMailSchema = z.object({
  type_mail_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Jenis surat harus diisi",
    }
  ),
  reason_mail: z
    .string({ required_error: "Keperluan surat harus diisi" })
    .min(1, { message: "Keperluan surat harus diisi" }),
  document: z
    .instanceof(File)
    .optional()
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
        return file.type === "image/png" || file.type === "image/jpeg";
      },
      { message: "File dengan format png dan jpg" }
    ),
  send_mail: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Kanal harus diisi",
    }
  ),
});

export type FormMailSchemaType = z.infer<typeof FormMailSchema>;
