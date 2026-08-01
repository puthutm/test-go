interface InformationStudent {
  id?: string;
  college_email?: string | null;
  private_email: string;
  phone: string;
  transportation_id: string;
  transportation_name?: string;
  citizenship_id: string;
  citizenship_name?: string;
  almamater_size_id: string;
  almamater_size_name?: string;
  job_id: string;
  job_name?: string;
  study_program_id?: string;
}

type FormInformationStudent = Omit<
  InformationStudent,
  | "id"
  | "college_email"
  | "transportation_name"
  | "citizenship_name"
  | "almamater_size_name"
  | "job_name"
>;
