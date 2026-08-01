interface Class {
  id: string;
  name: string;
  code: string;
  academic_periode_id: string;
  academic_periode_name: string;
  subject_id: string;
  study_program_id: string;
  curriculum_year_id: string;
  curriculum_year_name: string;
  capacity: number;
  number_of_meeting: number;
  contract_description: string;
  contract_file_path: string;
  total_participant: number;
  subject_name_id: string;
  subject_name_en: string;
  lecturer_name: string;
  study_program_name: string;
  day_name: string;
  start_time: string;
  end_time: string;
  subject_total_sks: number;
}

interface ClassSchedule {
  id: string;
  code: string;
  name: string | null;
  capacity: number;
  total_participant: number;
  subject_name_id: string;
  subject_name_en: string;
  lecturer_name: string;
  study_program_name: string;
  curriculum_year_name: string;
  day_name: string;
  start_time: string;
  end_time: string;
}

interface ClassScheduleDetail {
  id: string;
  code: string;
  name: string | null;
  academic_periode_id: string;
  subject_id: string;
  study_program_id: string;
  curriculum_year_id: string;
  capacity: number;
  number_of_meeting: number;
  contract_description: string | null;
  contract_file_path: string | null;
  created_at: number;
  created_by: string;
  updated_at: number;
  updated_by: string;
  deleted_at: string | null;
  deleted_by: string | null;
  total_participant: number;
  subject_name_id: string;
  subject_name_en: string;
  lecturer_name: string;
  study_program_name: string;
  lecturer_system: string;
  subject_total_sks: number;
  curriculum_year_name: string;
  start_date_of_college: string;
  end_date_of_college: string;
  academic_periode_fullname: string;
}

interface IQueryParamsClassSchedule extends QueryParam {
  study_program_id?: string | null;
  curriculum_year_id?: string | null;
  class_id?: string | null;
}

interface WeeklySchedule {
  day_name: string;
  date: string;
  start_time: string;
  end_time: string;
  type_of_meeting: string;
  material_attachment_file_path: string | null;
  attendance_document_file_path: string | null;
  journal_document_file_path: string | null;
  material_plan: string | null;
  material_realization: string | null;
}

type DistributionOfStudyProgram = Pick<
  IClassScheduleDetail,
  "id" | "code" | "name"
>;

interface AcademicSystemDistribution {
  type_of_meeting: string;
}

type StudentClassDistribution = Pick<
  IClassScheduleDetail,
  "id" | "code" | "name"
>;

interface IClassParticipant {
  id: string;
  class_id: string;
  student_id: string;
  student_nim: string;
  student_name: string;
  study_program_name: string;
  year_of_entry: string;
}

interface IClassAttendance {
  id: string;
  class_id: string;
  session: number;
  day_name: string;
  date: string;
  start_time: string;
  end_time: string;
  type_of_meeting: string;
  material_attachment_file_path: string | null;
  attendance_document_file_path: string | null;
  journal_document_file_path: string | null;
  material_plan: string;
  material_realization: string;
}

interface IClassScheduleSubDetail {
  id: string;
  class_id: string;
  session: number;
  day_name: string;
  date: string;
  start_time: string;
  end_time: string;
  type_of_meeting: string;
  material_attachment_file_path: string | null;
  attendance_document_file_path: string | null;
  journal_document_file_path: string | null;
  material_plan: string;
  material_realization: string;
}

type ICourseContract = Pick<
  ClassScheduleDetail,
  | "id"
  | "code"
  | "name"
  | "academic_periode_id"
  | "subject_id"
  | "study_program_id"
  | "curriculum_year_id"
  | "capacity"
  | "number_of_meeting"
  | "contract_description"
  | "contract_file_path"
  | "created_at"
  | "created_by"
  | "updated_at"
  | "updated_by"
  | "deleted_at"
  | "deleted_by"
  | "total_participant"
  | "subject_name_id"
  | "subject_name_en"
  | "lecturer_name"
  | "subject_total_sks"
  | "curriculum_year_name"
  | "start_date_of_college"
  | "end_date_of_college"
  | "academic_periode_fullname"
  | "lecturer_system"
  | "study_program_name"
>;

interface ICourseAssignment {
  id: string;
  title: string;
  description: string;
  is_gradeable: boolean | string;
  is_sharing: boolean | string;
  is_use_deadline: boolean | string;
  retake: number | string;
  schedule_id: string;
  session_schedule: number;
  sharing_date: string | null;
  deadline_of_assignment_submission: number | string;
  time_to_open: number | string;
  total_collect: number;
  views: number;
  updated_at: number | null;
  created_at: number;
  deleted_at: number | null;
}

type IFormCourseAssignment = Pick<
  ICourseAssignment,
  | "schedule_id"
  | "title"
  | "description"
  | "is_gradeable"
  | "is_use_deadline"
  | "deadline_of_assignment_submission"
  | "time_to_open"
  | "retake"
>;

type ClassForm = Pick<
  Class,
  | "code"
  | "name"
  | "academic_periode_id"
  | "subject_id"
  | "study_program_id"
  | "curriculum_year_id"
  | "capacity"
>;

interface ClassScore {
  student_id: string;
  nim: string;
  student_name: string;

  presence_score: number;
  task_score: number;
  uts_score: number;
  uas_score: number;
  final_score: number;

  quality_value: number;

  grade_id: string;
  grade_code: string;
  grade_name: string;
  grade_description: string;

  is_passed: boolean;
  pass_note: string;

  limit_grade_id: string;
  limit_grade_code: string;
  limit_grade_lower_limit: number;
}
