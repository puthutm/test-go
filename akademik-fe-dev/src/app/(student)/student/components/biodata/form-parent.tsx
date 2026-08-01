import { FormFather } from "./form-father";
import { FormMother } from "./form-mother";
import { getFather } from "@/services/api/students/biodata/parent/get-father";
import { getMother } from "@/services/api/students/biodata/parent/get-mother";

export async function FormParent() {
  const [father, mother] = await Promise.all([getFather(), getMother()]);

  return (
    <>
      <FormFather father={father as ApiResponse<ParentStudent>} />
      <hr className="mt-3 mb-3" />
      <FormMother mother={mother as ApiResponse<ParentStudent>} />
    </>
  );
}
