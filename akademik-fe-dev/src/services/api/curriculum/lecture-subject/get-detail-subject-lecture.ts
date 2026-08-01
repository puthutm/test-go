'use server'

import { fetchApi } from "@/lib/utils/fetch-server"


export const getSubjectDetailLecture = async (idSubject:string) : Promise<ApiResponse<ILectureSubjects>> =>{
    try{
        const xhr = await fetchApi(`/lecturer/curriculum/subjects/${idSubject}`,{
            method:'GET'
        })
        return xhr
    }
    catch(err:any){
        throw new Error(err.message || 'gagal get data mata kuliah dosen')
    }
}