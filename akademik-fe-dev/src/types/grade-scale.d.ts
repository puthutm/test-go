// {
//   "id": "2817631a-9aea-4924-b339-0f3d98eac3c4",
//   "study_program_id": "123e4567-e89b-12d3-a456-426614174000",
//   "grade_id": "123e4567-e89b-12d3-a456-426614174001",
//   "weight_value": 85.5,
//   "lower_value": 60,
//   "upper_value": 90,
//   "description": "This scale represents a range of grade values.",
//   "created_at": 1742787661254,
//   "created_by": "591047e8-f32e-48fc-857d-19a7ec0b5361",
//   "updated_at": 1742787716196,
//   "updated_by": "591047e8-f32e-48fc-857d-19a7ec0b5361",
//   "deleted_at": null,
//   "deleted_by": null,
//   "study_program_name": "",
//   "grade_name": ""
// }

interface IGradeScale {
  id: string;
  study_program_id: string;
  study_program_name: string;
  grade_id: string;
  grade_name: string;
  weight_value: number;
  lower_value: number;
  upper_value: number;
  description: string;
  created_at: number | null;
  created_by: string;
  updated_at: number | null;
  updated_by: string;
  deleted_at: number | null;
  deleted_by: string | null;
}

interface IQueryParamsGradeScale extends QueryParam {
  sort?:string,
  study_program_id?:string,
  grade_id?:string
}

type IFormGradeScale = Pick<IGradeScale,'study_program_id'|'grade_id'|'weight_value' | 'lower_value' | 'upper_value' | 'description'>