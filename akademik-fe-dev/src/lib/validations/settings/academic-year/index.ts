import { z } from "zod";

export const formAcademicYearSchema = z.object({
  name: z.string({ required_error: "Nama harus diisi" }).min(1, {
    message: "Nama harus diisi",
  }),
  years: z
    .string({ required_error: "Tahun harus diisi" })
    .min(4, {
      message: "Tahun tidak valid",
    })
    .max(4, {
      message: "Tahun tidak valid",
    }),
});

export type FormAcademicYearSchemaType = z.infer<typeof formAcademicYearSchema>;

export const formAcademicYearSchemaDefaultValues = {
  name: "",
  years: "",
};
