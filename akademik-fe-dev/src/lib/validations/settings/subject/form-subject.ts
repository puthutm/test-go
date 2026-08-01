import { z } from "zod";

export const FormSubjectSchema = z.object({
  curriculum_year_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Tahun kurikulum harus dipilih",
    }
  ),
  study_program_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Unit pengampu harus dipilih",
    }
  ),
  course_type_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Jenis mata kuliah harus dipilih",
    }
  ),
  course_group_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Kelompok mata kuliah harus dipilih",
    }
  ),
  code: z
    .string({ message: "Kode mata kuliah harus diisi" })
    .min(1, { message: "Kode mata kuliah harus diisi" }),
  name_id: z
    .string({ message: "Nama mata kuliah harus diisi" })
    .min(1, { message: "Nama mata kuliah harus diisi" }),
  name_en: z
    .string({ message: "Nama mata kuliah dalam bahasa inggris harus diisi" })
    .min(1, { message: "Nama mata kuliah dalam bahasa inggris harus diisi" }),
  face_to_face_sks: z
    .number({ message: "SKS tatap muka harus diisi" })
    .min(1, { message: "SKS tatap muka harus diisi" }),
  practicum_sks: z.number().optional(),
  field_practice_sks: z.number().optional(),
  simulation_sks: z.number().optional(),
  field_of_studies_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "Rumpun mata kuliah harus dipilih",
    }
  ),
  supporting_lecturer_id: z.array(
    z.object(
      {
        label: z.string(),
        value: z.string(),
      },
      {
        message: "Dosen pengampu harus dipilih",
      }
    )
  ),
  developer_rps_lecturer_id: z.array(
    z.object(
      {
        label: z.string(),
        value: z.string(),
      },
      {
        message: "Dosen pembuat RPS harus dipilih",
      }
    )
  ),
  subject_coordinator_lecturer_id: z.array(
    z.object(
      {
        label: z.string(),
        value: z.string(),
      },
      {
        message: "Koordinator pengampu mata kuliah harus dipilih",
      }
    )
  ),
  is_mku: z.boolean().default(false),
  is_sap: z.boolean().default(false),
  is_silabus: z.boolean().default(false),
  is_teaching_material: z.boolean().default(false),
  is_diktat: z.boolean().default(false),
});

export type FormSubjectSchemaType = z.infer<typeof FormSubjectSchema>;
