import { z } from "zod";

export const FormDocumentSchema = z.object({
  npwp: z.string().optional(),
  npwp_filepath: z
    .union([
      z.instanceof(File).optional().nullable().optional().nullable(),
      z.string().url().optional().nullable(),
    ])
    .optional()
    .refine(
      (file) => {
        if (!file || typeof file === "string") return true;
        return file.size <= 2 * 1024 * 1024;
      },
      { message: "File maksimal 2MB" }
    )
    .refine(
      (file) => {
        if (!file || typeof file === "string") return true;
        return file.type === "image/png" || file.type === "image/jpeg";
      },
      { message: "File harus dalam format PNG atau JPG" }
    ),

  // BPJS Kesehatan (Health Insurance)
  bpjs_healthcare: z.string().optional(),
  bpjs_healthcare_filepath: z
    .union([
      z.instanceof(File).optional().nullable(),
      z.string().url().optional().nullable(),
    ])
    .optional()
    .refine(
      (file) => {
        if (!file || typeof file === "string") return true;
        return file.size <= 2 * 1024 * 1024;
      },
      { message: "File maksimal 2MB" }
    )
    .refine(
      (file) => {
        if (!file || typeof file === "string") return true;
        return file.type === "image/png" || file.type === "image/jpeg";
      },
      { message: "File harus dalam format PNG atau JPG" }
    ),

  // BPJS Ketenagakerjaan (Employment Insurance)
  bpjs_employment: z.string().optional(),
  bpjs_employment_filepath: z
    .union([
      z.instanceof(File).optional().nullable(),
      z.string().url().optional().nullable(),
    ])
    .optional()
    .refine(
      (file) => {
        if (!file || typeof file === "string") return true;
        return file.size <= 2 * 1024 * 1024;
      },
      { message: "File maksimal 2MB" }
    )
    .refine(
      (file) => {
        if (!file || typeof file === "string") return true;
        return file.type === "image/png" || file.type === "image/jpeg";
      },
      { message: "File harus dalam format PNG atau JPG" }
    ),
});

export type FormDocumentSchemaType = z.infer<typeof FormDocumentSchema>;
