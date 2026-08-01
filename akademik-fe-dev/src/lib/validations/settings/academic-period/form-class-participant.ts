import { z } from "zod";

export const FormClassParticipantSchema = z.object({
  student_id: z
    .string({ message: "Mahasiswa harus diisi" })
    .min(1, { message: "Mahasiswa harus diisi" }),
});

export type FormClassParticipantSchemaType = z.infer<
  typeof FormClassParticipantSchema
>;
