import { z } from "zod";

// {
//     "study_program_id": "123e4567-e89b-12d3-a456-426614174000",  // UUID for the study program (required)
//     "grade_id": "123e4567-e89b-12d3-a456-426614174001",           // UUID for the grade (required)
//     "weight_value": 85.5,                                           // Weight value for the grade (required, between 0 and 100)
//     "lower_value": 60.2,                                            // Lower boundary of the scale (required)
//     "upper_value": 90.2,                                            // Upper boundary of the scale (required)
//     "description": "This scale represents a range of grade values." // Description for the scale (required, max length 1000 characters)
// }

export const GradeScaleSchema = z.object({
  study_program_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "program studi harus dipilih",
    }
  ),
  grade_id: z.object(
    {
      label: z.string(),
      value: z.string(),
    },
    {
      message: "nilai harus dipilih",
    }
  ),
  weight_value: z
    .string()
    .min(1, "bobot nilai harus diisi")
    .refine(
      (val) => !isNaN(Number(val)) && Number(val) >= 1 && Number(val) <= 100,
      "bobot nilai harus berupa angka antara 1 sampai 100"
    ),
  lower_value: z
    .string()
    .min(1, "nilai bawah harus diisi")
    .refine(
      (val) => !isNaN(Number(val)) && Number(val) >= 1 && Number(val) <= 100,
      "nilai Bawah harus berupa angka antara 1 sampai 100"
    ),
  upper_value: z
    .string()
    .min(1, "nilai Atas harus diisi")
    .refine(
      (val) => !isNaN(Number(val)) && Number(val) >= 1 && Number(val) <= 100,
      "nilai Bawah harus berupa angka antara 1 sampai 100"
    ),
  description: z
    .string()
    .min(1, "deskripsi harus di isi")
    .max(1000, "deskripsi tidak boleh lebih dari 1000 character"),
});

export type GradeScaleFormType = z.infer<typeof GradeScaleSchema>;
