
import { fetchApi } from "@/lib/utils/fetch-server";
import { FormSchemeCourseAssigment } from "@/lib/validations/academic/settings/college-class/form-course-assigment";

export const createCourseAssignment = async (idClass:string,dataForm:FormSchemeCourseAssigment) : Promise<ApiResponse<ICourseAssignment | undefined>> =>{
    const payload :IFormCourseAssignment = {
        schedule_id:dataForm.schedule_id.value,
        title:dataForm.title,
        description:dataForm.description,
        deadline_of_assignment_submission:String(dataForm.deadline_of_assignment_submission),
        time_to_open:String(dataForm.time_to_open),
        retake:String(dataForm.retake),
        is_gradeable:String(dataForm.is_gradeable),
        is_use_deadline:String(dataForm.is_use_deadline)
    } 

    try{
        const response = await fetchApi(`/lecturer/academic/class-schedules/${idClass}/course-assisments`,{
            method:"POST",
            body:JSON.stringify(payload),
            headers: {
                "Content-Type": "application/json",
            },
        })

        return response
    }
    catch (error: any) {
    throw new Error(error);
  }
}