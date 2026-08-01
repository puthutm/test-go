interface Subject {
  id: string;
  curriculum_year_id: string;
  study_program_id: string;
  code: string;
  name_id: string;
  name_en: string;
  course_type_id: string;
  course_group_id: string;
  face_to_face_sks: number;
  practicum_sks: number;
  field_practice_sks: number;
  simulation_sks?: number | null;
  total_sks: number;
  field_of_studies_id: string;
  is_mku: boolean;
  is_sap: boolean;
  is_silabus: boolean;
  is_teaching_material: boolean;
  is_diktat: boolean;
  created_at: number;
  updated_at: number;
  deleted_at: number | null;
  study_program_name: string;
  curriculum_year_name: string;
  course_type_name: string;
  course_group_name: string;
  field_study_name: string;
  supporting_lecturers: {
    id: string;
    lecturer_name: string;
    lecturer_front_title: string;
    lecturer_back_title: string;
  }[];
  developer_rps_lecturers: {
    id: string;
    lecturer_name: string;
    lecturer_front_title: string;
    lecturer_back_title: string;
  }[];
  subject_coordinator_lecturers: {
    id: string;
    lecturer_name: string;
    lecturer_front_title: string;
    lecturer_back_title: string;
  }[];
}
