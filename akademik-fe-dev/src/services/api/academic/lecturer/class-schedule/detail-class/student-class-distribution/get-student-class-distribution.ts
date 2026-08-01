'use server'

import { fetchApi } from "@/lib/utils/fetch-server"


export const getStudentClassDistribution = async (idClass:string) : Promise<ApiResponse<StudentClassDistribution[]>> =>{
    try{
        const xhr = await fetchApi(`/lecturer/academic/class-schedules/${idClass}/student-class-distributions`,{
            method:'GET'
        })
        return xhr
    }
    catch(err:any){
        throw new Error(err.message || 'gagal get data kelas mahasiswa')
    }
}