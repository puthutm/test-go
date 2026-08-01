interface SubjectPrerequisites {
  id: string;
  subject_id: string;
  subject_code: string;
  subject_name_id: string;
  subject_name_en: string;
  semester_number_id: string;
  semester_number: string;
}

interface SubjectPrerequisitesParams {
  curriculumYearId: string;
  studyProgramId?: string;
  semesterNumberId: string;
}
