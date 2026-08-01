interface PortalStudent {
  student_id: string;
  student_name: string;
  student_nim: string;
  study_program_name: string;
  student_status: string;
}

interface PortalStudentAcademic {
  student_id: string;
  student_study_program_id: string;
  nim: string;
  student_name: string;
  student_status: "active" | "inactive" | string;

  study_program_id: string;
  study_program_code: string;
  study_program_name: string;

  lecturer_pa_biodata_id: string | null;
  lecturer_pa_user_id: string | null;
  lecturer_pa_name: string;

  entry_academic_period_id: string;
  entry_academic_period_name: string;

  current_academic_period_id: string;
  current_academic_period_name: string;
  current_academic_period_shortname: string;

  entry_period_rank: number;
  current_period_rank: number;
  current_semester: number;
  semester_label: string;

  total_sks_taken: number;
  ipk: number;

  created_at: number; // timestamp (ms)
}

interface Batch {
  batch_detail_id: string;
  batch_name: string;
}
