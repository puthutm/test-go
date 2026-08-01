// {
//     "id": "81afa892-7a24-453b-b3d4-028b1c23fe43",
//     "value_element_id": "123e4567-e89b-12d3-a456-426614174000",
//     "percentage": 15.5,
//     "is_passing_requirement": true,
//     "created_at": 1742788201121,
//     "created_by": "591047e8-f32e-48fc-857d-19a7ec0b5361",
//     "updated_at": 1742789028359,
//     "updated_by": "591047e8-f32e-48fc-857d-19a7ec0b5361",
//     "deleted_at": null,
//     "deleted_by": "591047e8-f32e-48fc-857d-19a7ec0b5361",
//     "value_element_name": ""
// }

interface IGradeComposition {
  id: string;
  academic_periode_id: string;
  academic_periode_name: string;
  value_element_id: string;
  percentage: number;
  is_passing_requirement: boolean;
  created_at: number;
  created_by: string;
  updated_at: number;
  updated_by: string;
  deleted_at: number | null;
  deleted_by: string | null;
  value_element_name: string;
}

interface IQueryParamsGradeComposition extends QueryParam {
  sort?: string;
  value_element_id?: string;
}

type iGradeCompositionForm = Pick<
  IGradeComposition,
  "value_element_id" | "percentage" | "is_passing_requirement"
>;
