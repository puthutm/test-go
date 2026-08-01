interface FinalProjectProposal {
  id: string;
  student_id?: string;
  student_name?: string;
  title_id: string;
  title_en: string;
  topic: string;
  academic_period_id: string;
  study_program_id: string;
  study_program_name?: string;
  lecturer_name?: string[];
  abstract: string;
  file_path: string;
  status: number;
  date: number;
  confirmation_status_date: Date;
  confirmation_by: string;
  feedback: string;
}

type FinalProjectProposalStudent = Pick<
  FinalProjectProposal,
  "id" | "title_id" | "title_en" | "status" | "date"
>;
