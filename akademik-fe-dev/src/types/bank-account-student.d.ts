interface BankAccountStudent {
  id?: string;
  bank_id: string | null;
  bank_name?: string | null;
  account_number?: string | null;
  account_name?: string | null;
  account_filepath?: string | null;
  bank_name?: string | null;
}

type FormBankAccountStudent = Omit<BankAccountStudent, "id">;
