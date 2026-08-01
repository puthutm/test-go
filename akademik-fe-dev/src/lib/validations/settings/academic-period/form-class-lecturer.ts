import { z } from "zod";

export const FormClassLecturerSchema = z.object({
  lecturer_id: z
    .string({ message: "Dosen pengajara harus diisi" })
    .min(1, { message: "Dosen pengajara harus diisi" }),
  subtitute_lecturer_id: z.string().optional().nullable(),
});

export type FormClassLecturerSchemaType = z.infer<
  typeof FormClassLecturerSchema
>;
