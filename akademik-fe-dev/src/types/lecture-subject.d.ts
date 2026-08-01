type ILectureSubjects  = Subject

type ILectureSubjectsCordinator = Subject

interface IQueryParamsLectureSubject {
    sort?:string | null,
    sort_by?:string | null,
    page:number,
    limit?:number,
    search?:string | null,
    academic_periode_id?:string | null,
    study_program_id?:string|null,
    subject_group_id?:string|null,
    curriculum_year_id?:string|null
}

interface IQueryParamsLectureSubjectCordinator {
    sort?:string | null,
    sort_by?:string | null,
    page:number,
    limit?:number,
    search?:string | null,
    subject_type_id?:string|null,
    subject_group_id?:string|null,
    curriculum_year_id?:string|null
}