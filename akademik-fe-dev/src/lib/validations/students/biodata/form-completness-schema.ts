import { z } from "zod";

// Main validation schema
export const FormCompletenessSchema = z.object({
  // NIS (Nomor Induk Siswa) - Required
  nis: z.string().optional(),

  // NISN (Nomor Induk Siswa Nasional) - Required
  nisn: z.string().optional(),

  // No. Paspor - Optional
  no_passport: z.string().optional(),

  // Google Scholar - Optional
  google_scholar: z.string().optional(),

  // ID Sinta - Optional
  sinta_id: z.string().optional(),

  // ID Scopus - Optional
  scopus_id: z.string().optional(),

  // File Tanda Tangan - Optional but with validation if provided
  signature_path_file: z
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

export type FormCompletenessSchemaType = z.infer<typeof FormCompletenessSchema>;
