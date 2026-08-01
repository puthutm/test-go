interface BiodataStudent {
  id?: string;
  nik: string;
  no_kk: string;
  name: string;
  birth_date: string;
  back_degree?: string | null;
  birth_place_id: string;
  birth_place_name: string;
  blood_type_id?: string | null;
  blood_type_name: string;
  ethnic_id: string;
  ethnic_name: string;
  gender: string;
  religion_id: string;
  religion_name: string;
  height?: number | null;
  weight?: number | null;
  status_id: string;
  status_name: string;
}

type FormBiodataStudent = Omit<
  BiodataStudent,
  | "id"
  | "status_name"
  | "birth_place_name"
  | "blood_type_name"
  | "ethnic_name"
  | "religion_name"
>;
