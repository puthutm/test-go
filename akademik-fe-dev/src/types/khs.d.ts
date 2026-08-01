interface KhsSemester {
  academic_periode_id: string;
  academic_periode_name: string;
  total_weight: number;
  total_sks: number;
  ips: number;
  subjects: KhsSubject[];
}

interface KhsSubject {
  subject_code: string;
  subject_name: string;
  total_sks: number;
  final_score: number;
  grade_code: string;
  quality_value: number;
  weight: number;
  is_passed: boolean;
}

type NotPassedItem = Pick<
  KhsSubject,
  "subject_code" | "subject_name" | "final_score" | "grade_code"
> &
  Pick<KhsSemester, "academic_periode_id" | "academic_periode_name">;

interface Khs {
  semesters: KhsSemester[];
  not_passed: NotPassedItem[];
}
