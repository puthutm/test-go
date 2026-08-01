import { z } from "zod";

// {
//     "curriculum_year_id": "a8f8a44f-d54f-42a3-b590-df3812b330f9",
//     "study_program_id": "aab7fa45-d5fc-453b-8009-1a12c2e8bcba", // program study
//     "course_type_id": "def3d490-f09e-42e5-a7b4-e72793b08361",
//     "course_group_id": "4e6cb67d-f60b-48c9-a018-b9d9a923b68f",
//     "code": "CS102",
//     "name_id": "Computer Science 101",
//     "name_en": "Computer Science 101",
//     "face_to_face_sks": 3,
//     "practicum_sks": 2,
//     "field_practice_sks": 1,
//     "simulation_sks": 1, // optional
//     "total_sks": 7,
//     "field_of_studies_id": "77e5c42c-93f2-42ec-86cf-f508f38a2672",
//     "supporting_lecturer_id": "d4f7d8a4-5e2f-4a3b-9f61-cff410a57f95",
//     "developer_rps_lecturer_id": "15f3b09f-84f3-4b93-88db-2649442d2432",
//     "subject_coordinator_lecturer_id": "5178b19c-df99-47d2-ae35-015d4bffb66c",
//     "is_mku": true,
//     "is_sap": false,
//     "is_silabus": true,
//     "is_teaching_material": true,
//     "is_diktat": false
//   }

export const SubjectFormSchema = z.object({
  curriculum_year_id: z.string().nonempty(),
  study_program_id: z.string().nonempty(),
  course_type_id: z.string().nonempty(),
  course_group_id: z.string().nonempty(),
  code: z.string().nonempty(),
  name_id: z.string().nonempty(),
  name_en: z.string().nonempty(),
  face_to_face_sks: z.number().min(0),
  practicum_sks: z.number().min(0),
  field_practice_sks: z.number().min(0),
  simulation_sks: z.number().min(0).optional(),
  total_sks: z.number().min(0),
  field_of_studies_id: z.string().nonempty(),
  supporting_lecturer_id: z.string().nonempty(),
  developer_rps_lecturer_id: z.string().nonempty(),
  subject_coordinator_lecturer_id: z.string().nonempty(),
  is_mku: z.boolean(),
  is_sap: z.boolean(),
  is_silabus: z.boolean(),
  is_teaching_material: z.boolean(),
  is_diktat: z.boolean(),
});

export type SubjectFormType = z.infer<typeof SubjectFormSchema>;

export const SubjectInitialValue: SubjectFormType = {
  curriculum_year_id: "",
  study_program_id: "",
  course_type_id: "",
  course_group_id: "",
  code: "",
  name_id: "",
  name_en: "",
  face_to_face_sks: 0,
  practicum_sks: 0,
  field_practice_sks: 0,
  simulation_sks: 0,
  total_sks: 0,
  field_of_studies_id: "",
  supporting_lecturer_id: "",
  developer_rps_lecturer_id: "",
  subject_coordinator_lecturer_id: "",
  is_mku: false,
  is_sap: false,
  is_silabus: false,
  is_teaching_material: false,
  is_diktat: false,
};
