interface KRS {
  krs_header_id: string;
  academic_periode_id: string;
  academic_periode_name: string;
  student_name: string;
  student_nim: string;
  total_sks: number;
  student_id: string;
}

interface KrsItem {
  krs_item_id: string;
  krs_id: string;
  class_id: string;
  subject_id: string;
  subject_code: string;
  subject_name: string;
  subject_name_id: string;
  subject_name_en: string;
  schedule: string;
  lecturer_names: string;
  class_code: string;
  class_name: string;
  sks: number;
  total_sks: number;
  item_status: string;
}

interface KrsAcademicPeriod {
  id: string;
  code: string;
  fullname: string;
  shortname: string;
  is_active: boolean;
  start_date_of_college: number;
  end_date_of_college: number;
}

interface KrsAvailableClass {
  class_id: string;
  subject_id: string;
  subject_code: string;
  subject_name: string;
  schedule: string;
  class_code: string;
  class_name: string;
  sks: number;
  capacity: number;
  used_quota: number;
  quota_text: string;
  button_state: string;
}

interface KrsListData {
  academic_periods: KrsAcademicPeriod[];
  classes: KrsAvailableClass[];
}

interface KrsDetail
  extends Pick<KRS, "student_name" | "student_nim" | "academic_periode_name"> {
  student_id: string;
  study_program_name: string;
  total_sks_taken: number;
  krs_items: KrsItem[];
}

interface InfoKrs {
  student_id: string;
  academic_periode_id: string;
  academic_periode_id_before: string;
  ips_before: number;
  sks_limit_id: string;
  ips_min: number;
  ips_max: number;
  max_sks: number;
}
