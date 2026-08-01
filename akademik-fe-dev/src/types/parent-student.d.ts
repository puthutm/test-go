interface ParentStudent {
  id?: string;
  student_id?: string;
  name: string;
  nik: string;
  educational_level_id?: string;
  educational_level_name?: string;
  type?: string;
  phone?: string | null;
  phone2?: string | null;
  email?: string | null;
  kinship?: string | null;
  status_kinship?: string | null;
  life_status?: string | null;
  address?: string | null;
  birth_place_id?: string | null;
  birth_place_name?: string | null;
  birth_date?: string | null;
  job_id: string;
  job_name: string;
  income?: string | null;
}

type FormParentStudent = Omit<
  ParentStudent,
  "id" | "student_id",
  "educational_level_name"
>;
