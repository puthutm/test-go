interface Presences {
  id: string;
  academic_periode_id: string;
  academic_periode_name: string;
  study_program_id: string;
  study_program_name: string;
}

interface SubjectsPresence {
  subject_id: string;
  subject_name_id: string;
  subject_name_en: string;
  academic_periode_id: string;
  academic_periode_name: string;
  study_program_id: string;
  lecturer_id: string;
  class_count: string;
}

interface ClassPresence {
  class_id: string;
  class_name: string;
  class_code: string;
  study_program_id: string;
}

interface PresenceComponent {
  use_open_session: boolean;
  open_session_percentage: number;

  use_document_material: boolean;
  document_material_percentage: number;

  use_quiz: boolean;
  quiz_percentage: number;

  use_task: boolean;
  task_percentage: number;

  use_video: boolean;
  video_percentage: number;

  use_uts: boolean;
  uts_percentage: number;

  use_uas: boolean;
  uas_percentage: number;

  use_comment: boolean;
  comment_percentage: number;
}

interface ClassPresenceSession {
  session_id: string;
  session: number;
  session_date: string;
  presence_percentage: number;
}

interface StudentPresenceSession {
  student_id: string;
  student_name: string;
  student_nim: string;
  open_session_percentage: number;
  document_material_percentage: number;
  quiz_percentage: number;
  task_percentage: number;
  video_percentage: number;
  uts_percentage: number;
  uas_percentage: number;
  comment_percentage: number;
  total_percentage: number;
  presence_flag: number;
}

type StudentPresenceComponent =
  | "open_session"
  | "document_material"
  | "quiz"
  | "task"
  | "video"
  | "uts"
  | "uas"
  | "comment";
