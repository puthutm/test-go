interface ClassScheduleTemplate {
  id: string;
  class_id: string;
  day_name: string;
  start_time: string;
  end_time: string;
  type_of_meeting: string;
}

interface ClassScheduleSession {
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
  material_plan: string | null;
  material_realization: string | null;
  is_uts: boolean;
  is_uas: boolean;
}
