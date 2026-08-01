interface DocumentStudent {
  id?: string;
  npwp?: string | null;
  npwp_filepath?: string | null;
  bpjs_healthcare?: string | null;
  bpjs_healthcare_filepath?: string | null;
  bpjs_employment?: string | null;
  bpjs_employment_filepath?: string | null;
}

type FormDocumentStudent = Omit<DocumentStudent, "id">;
