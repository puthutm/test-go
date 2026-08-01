interface CompletenessStudent {
  id?: string;
  no_passport: string;
  google_scholar?: string;
  sinta_id?: string;
  scopus_id?: string;
  signature_path_file?: string;
}

type FormCompletenessStudent = Omit<CompletenessStudent, "id">;
