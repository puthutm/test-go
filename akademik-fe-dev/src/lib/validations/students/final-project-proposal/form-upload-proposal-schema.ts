import { z } from "zod";

export const FormUploadFinalProjectProposalSchema = z.object({
  title_id: z
    .string({ required_error: "Judul proposal harus diisi" })
    .min(1, { message: "Judul proposal harus diisi" }),
  title_en: z
    .string({ required_error: "Judul proposal harus diisi" })
    .min(1, { message: "Judul proposal harus diisi" }),
  topic: z
    .string({ required_error: "Topik harus diisi" })
    .min(1, { message: "Topik harus diisi" }),
  abstract: z
    .string({ required_error: "Abstrak harus diisi" })
    .min(1, { message: "Abstrak harus diisi" }),
  file: z
    .union([
      z.instanceof(File).optional().nullable(),
      z.string().url().optional().nullable(),
    ])
    .optional()
    .refine(
      (file) => {
        if (!file || typeof file === "string") return true;
        return file.size <= 10 * 1024 * 1024;
      },
      { message: "File maksimal 10MB" }
    )
    .refine(
      (file) => {
        if (!file || typeof file === "string") return true;
        return file.type === "application/pdf";
      },
      { message: "File harus dalam format PDF" }
    ),
});

export type FormUploadPropFinalProjectosalSchemaType = z.infer<
  typeof FormUploadFinalProjectProposalSchema
>;
