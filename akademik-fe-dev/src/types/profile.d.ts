interface Profile {
  nik: string;
  name: string;
  username: string;
  email: string;
  phone: string;
  email_private: string;
  avatar?: string;
  avatar_google?: string | null;
  created_at?: number;
  updated_at?: number;
  roles: Role[];
}
