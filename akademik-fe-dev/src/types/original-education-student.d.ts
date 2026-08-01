interface OriginalEducationStudent {
  id: string;
  student_id: string;
  school_id?: string | null;
  educational_level_id?: string | null;
  institution_name?: string | null;
  school_major?: string | null;
  year_of_graduation?: string | null;
  nisn?: string | null;
  province_id?: string | null;
  city_id?: string | null;
  national_exam_score?: string | null;
  certificate_number?: string | null;
  certificate_filepath?: string | null;
  transcripts_filepath?: string | null;
}

type FormOriginalEducationStudent = Pick<
  OriginalEducationStudent,
  | "institution_name"
  | "school_major"
  | "nisn"
  | "national_exam_score"
  | "certificate_number"
  | "certificate_filepath"
  | "transcripts_filepath"
>;
