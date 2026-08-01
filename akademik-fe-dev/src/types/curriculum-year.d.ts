interface CurriculumYear {
  id: string;
  years: string;
  starts: string;
  start_date: string;
  end_date: string;
  description?: string | null;
  academic_periode_name?: string;
}

type CurriculumYearForm = Omit<CurriculumYear, "id">;
