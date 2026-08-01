'use server'

import { fetchApi } from "@/lib/utils/fetch-server"


export const getSubjectLecture = async (queryParam:IQueryParamsLectureSubject) : Promise<ApiResponse<PaginationData<ILectureSubjects>>> =>{
  const params = new URLSearchParams();
  if (queryParam.search) {
    params.append("search", queryParam.search);
  }

  if (queryParam.page !==  undefined) {
    params.append("page", String(queryParam.page));
  }

  if (queryParam.limit !== undefined) {
    params.append("limit", String(queryParam.limit));
  }
    if (queryParam.curriculum_year_id !== undefined) {
    params.append("curriculum_year_id", String(queryParam.curriculum_year_id));
  }

  if (queryParam.sort_by) {
    params.append("sort_by", queryParam.sort_by);
  }
    try{
        const xhr = await fetchApi(`/lecturer/curriculum/subjects?${params.toString()}`,{
            method:'GET'
        })
        return xhr
    }
    catch(err:any){
        throw new Error(err.message || 'gagal get data mata kuliah dosen')
    }
}