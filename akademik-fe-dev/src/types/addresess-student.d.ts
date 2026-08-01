interface AddressStudent {
  id?: string;
  student_id?: string | null;
  country_id?: string | null;
  country_name?: string | null;
  province_id?: string | null;
  province_name?: string | null;
  city_id?: string | null;
  city_name?: string | null;
  district_id?: string | null;
  district_name?: string | null;
  village_id?: string | null;
  village_name?: string | null;
  rt?: string | null;
  rw?: string | null;
  address?: string | null;
  postal_code?: string | null;
  distance?: string | null;
}

type FormAddressStudent = Omit<
  AddressStudent,
  | "id"
  | "student_id"
  | "country_name"
  | "province_name"
  | "city_name"
  | "district_name"
  | "village_name"
>;
