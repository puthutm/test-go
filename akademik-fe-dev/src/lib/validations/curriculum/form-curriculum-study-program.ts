import { z } from "zod";

export const curriculumStudyProgramForProgramHeadSchema = z.object({
  curriculum_year_id: z
    .string({ required_error: "Tahun kurikulum harus diisi" })
    .min(1, { message: "Tahun kurikulum harus diisi" }),
  subject_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Mata kuliah harus dipilih",
    }
  ),
  semester_number_id: z
    .string({ message: "Semester harus diisi" })
    .min(1, { message: "Semester harus diisi" }),
  limit_grade_id: z
    .string({ message: "Nilai harus diisi" })
    .min(1, { message: "Nilai harus diisi" }),
  is_mandatory: z.boolean().default(false),
  field_study_concentration_id: z
    .object({
      label: z.string(),
      value: z.string(),
    })
    .optional()
    .nullable(),
  subject_prerequisites: z
    .array(
      z.object({
        label: z.string(),
        value: z.string(),
      })
    )
    .optional()
    .nullable(),
});

export const curriculumStudyProgramForAcademicSchema = z.object({
  curriculum_year_id: z
    .string({ required_error: "Tahun kurikulum harus diisi" })
    .min(1, { message: "Tahun kurikulum harus diisi" }),
  study_program_id: z.string().optional(),
  subject_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Mata kuliah harus dipilih",
    }
  ),
  semester_number_id: z
    .string({ message: "Semester harus diisi" })
    .min(1, { message: "Semester harus diisi" }),
  limit_grade_id: z
    .string({ message: "Nilai harus diisi" })
    .min(1, { message: "Nilai harus diisi" }),
  is_mandatory: z.boolean().default(false),
  field_study_concentration_id: z
    .object({
      label: z.string(),
      value: z.string(),
    })
    .optional()
    .nullable(),
  subject_prerequisites: z
    .array(
      z.object({
        label: z.string(),
        value: z.string(),
      })
    )
    .optional()
    .nullable(),
});

export type CurriculumStudyProgramForProgramHeadSchemaType = z.infer<
  typeof curriculumStudyProgramForProgramHeadSchema
>;

export type CurriculumStudyProgramForAcademicSchemaType = z.infer<
  typeof curriculumStudyProgramForAcademicSchema
>;
