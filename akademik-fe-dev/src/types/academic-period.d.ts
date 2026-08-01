interface AcademicPeriod {
  id: string;
  code: string;
  fullname: string;
  shortname: string;
  academic_year_id: string;
  academic_year: string;
  semester_id: string;
  semester: string;
  start_date_of_college: string;
  end_date_of_college: string;
  start_date_of_uts: string;
  end_date_of_uts: string;
  start_date_of_uas: string;
  end_date_of_uas: string;
  number_of_lecture_meeting: string;
  is_active: boolean;
  created_at: number;
  updated_at: number;
}

type AcademicPeriodForm = Pick<
  AcademicPeriod,
  | "code"
  | "fullname"
  | "academic_year_id"
  | "semester_id"
  | "shortname"
  | "start_date_of_college"
  | "end_date_of_college"
  | "start_date_of_uas"
  | "end_date_of_uas"
  | "start_date_of_uts"
  | "end_date_of_uts"
  | "number_of_lecture_meeting"
  | "is_active"
>;
