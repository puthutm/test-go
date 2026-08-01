interface CurriculumStudyProgram {
  data: {
    id: string;
    subject_code: string;
    curriculum_year_id: string;
    subject_id: string;
    semester_number_id: string;
    limit_grade_id: string;
    precondition: boolean;
    study_program_id: string;
    is_mandatory: boolean;
    subject_name_id: string;
    subject_name_en: string;
    subject_total_sks: number;
    grade_name: string;
    grade_code: string;
    field_study_concentration_id: string;
    field_study_concentration_name: string;
    field_study_concentration_code: string;
    subject_prerequisites: SubjectPrerequisites[];
  }[];
  total: {
    sks: number;
    mandatory: number;
    no_mandatory: number;
  };
}

interface CurriculumStudyProgramTable {
  id: string;
  subject_code: string;
  curriculum_year_id: string;
  subject_id: string;
  semester_number_id: string;
  limit_grade_id: string;
  precondition: boolean;
  study_program_id: string;
  is_mandatory: boolean;
  subject_name_id: string;
  subject_name_en: string;
  subject_total_sks: number;
  grade_name: string;
  grade_code: string;
  field_study_concentration_id: string;
  field_study_concentration_name: string;
  field_study_concentration_code: string;
  subject_prerequisites: SubjectPrerequisites[];
}
