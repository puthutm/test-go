
'use client'
import React,{Dispatch,SetStateAction, useState,useEffect} from 'react'
// import third pary component
import {
  Modal,
  ModalBody,
  Row,
  Button,
  Col,
} from 'reactstrap'

// import component
import { CloseIcon } from '@/components/icons/close'
import { SelectComponent } from '@/components/ui/select'
import { FormErrorMessage } from '@/components/ui/form-error-message'

import { useGetOptionCurriculumYear } from '@/services/api/data-referensi/curriculum-year/use-get-option-curriculum-year'

// import use form hook
import { useForm,Controller,SubmitHandler } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import schemeFormFilterSubject,{IFormSchemeFilterSubject} from '@/lib/validations/subjects/form-validation-filter-subject'
import { useSearchParams } from 'next/navigation'
import { useRouter } from 'next/navigation'
import { usePathname } from 'next/navigation'

import { IModalManipulationFilterSubjects } from './page-subject-lecturer';

function ModalFilterSubjects({
    showModal,
    setShowModal
}:{
    showModal:IModalManipulationFilterSubjects,
    setShowModal:Dispatch<SetStateAction<IModalManipulationFilterSubjects>>,
}
) {
  const router = useRouter();
  const pathname = usePathname()
  const searchParams = useSearchParams();
  const params = new URLSearchParams(searchParams);
    const dummyValue = [{
        value:'2025',
        label:'2025',
    }]

    // reack hook form
    const {
        control,
        handleSubmit,
        setValue,
        formState: { errors },
        reset,
        clearErrors
    } = useForm<IFormSchemeFilterSubject>({
        resolver: zodResolver(schemeFormFilterSubject)
    });

    const [queryParamsCurriculumYear,] = useState<QueryParamDataRefensi>({
        page:1,
        page_size:8,
        filter:''
    })

    //! get option curriculum year
    const {
        data:dataOptionCurriculumYear,
        isFetching:isLoadingCurriculumYear
    } = useGetOptionCurriculumYear(queryParamsCurriculumYear,showModal.status)

    useEffect(() => {
        if(searchParams.get('curriculum_year_id') && dataOptionCurriculumYear){
            const filterOption = dataOptionCurriculumYear?.data?.find((el:CurriculumYear)=>{
                return searchParams.get('curriculum_year_id') === el.id
            })

            const mapingOption = {
                value:filterOption?.id as string,
                label:filterOption?.years as string
            }
            setValue('year_curriculum',[mapingOption])
        }
    }, [showModal,dataOptionCurriculumYear])
    


  const curriculumYearOptions = dataOptionCurriculumYear?.data?.map((val: CurriculumYear) => ({
    label: val.years,
    value: val.id,
  }))||[];
    // event handle close
    const handleCloseModal = ()=>{
            setShowModal(()=>({
            status:false,
            title:'Filter',
        }))
        reset()
        clearErrors()
        setValue('year_curriculum',[])
        setValue('subject_type',[])
        setValue('subject_group',[])
    }

    // handle reset
    const handleReset = ()=>{
        setValue('year_curriculum',[])
        setValue('subject_type',[])
        setValue('subject_group',[])
        router.push(`${pathname}`)
    }
    // handle submit
    const handleSubmitFilter:SubmitHandler<IFormSchemeFilterSubject> =(dataInput,event) => {
        event?.preventDefault()
        if(dataInput.year_curriculum.length !== 0){
            params.set("curriculum_year_id", dataInput.year_curriculum[0].value);
        }
        router.push(`?${params.toString()}`)
    }
  

  return (
    <Modal
            isOpen={showModal.status}
            centered
            size="md"
            className="position-relative p-0"
            style={{ border: "0" }}
            >

            {/*//! modal header */}
            <section  className="px-4 ">
                <section className="position-relative py-3 d-flex align-items-center justify-content-end gap-2 border-bottom border-3">
                    <h2
                            style={{ fontSize: 20,color:'#3A3A3A'}}
                            className="m-0 p-0 fw-semibold flex-grow-1 w-100"
                    >
                            {showModal.title}
                    </h2>
                    
                    <Button
                    className='p-0'
                    color={'transparent'}
                    onClick={handleCloseModal}
                    >
                        <CloseIcon  width='25' hanging={'25'}/>
                    </Button>
                </section>
            </section>

            {/*//! modal body */}
            <ModalBody className="p-4 ">
                {/* form */}
                <form action=""
                 autoComplete="off"
                 onSubmit={handleSubmit(handleSubmitFilter)}
                 >
                    {/*//! input */}
                    <section className="gap-4 d-flex flex-column">
                        {/*//! col year curriculum */}
                        <Row
                            className="align-items-center gap-2"
                            >
                            {/* <Col sm={12}>
                                <Label htmlFor="selectYearCurriculum" className='form-label  m-0 d-flex align-items-center gap-1'> 
                                 Tahun Kurikulum
                                    <span className="m-0 p-0 fst-italic fw-semibold" style={{fontSize:"10px",color:"#3A3A3A"}}>
                                    - Optional
                                    </span> 
                                </Label>
                            </Col> */}
                            <Col sm={12}>
                                <Controller
                                    name="year_curriculum"
                                    control={control}
                                    render={({ field }) => {
                                        return (
                                            <SelectComponent
                                            {...field}
                                            options={curriculumYearOptions}
                                            isLoading={isLoadingCurriculumYear}
                                            // isDisabled={!isEdit || isLoadingRegistrantAddress}
                                            placeholder="Tahun Kurikulum"
                                            hasIcon={false}
                                            id={'selectYearCurriculum'}
                                            isError={errors.year_curriculum ? true : false}
                                            onChange={(value)=>{
                                                field.onChange(value !== null ? [value] : [])
                                            }}
                                            />
                                    );
                                    }}
                                />
                                <FormErrorMessage errors={errors.year_curriculum} />
                            </Col>
                        </Row>

                        {/*//! col subject_type */}
                        <Row
                            className="align-items-center gap-2"
                            >
                            {/* <Col sm={12}>
                                <Label htmlFor="selectsubjectType" className='form-label  m-0 d-flex align-items-center gap-1'> 
                                 Jenis Mata Kuliah
                                    <span className="m-0 p-0 fst-italic fw-semibold" style={{fontSize:"10px",color:"#3A3A3A"}}>
                                    - Optional
                                    </span> 
                                </Label>
                            </Col> */}
                            <Col sm={12}>
                                <Controller
                                    name="subject_type"
                                    control={control}
                                    render={({ field }) => {
                                        return (
                                            <SelectComponent
                                            {...field}
                                            options={dummyValue!}
                                            // isLoading={isLoadingOptionCountries || isLoadingRegistrantAddress}
                                            // isDisabled={!isEdit || isLoadingRegistrantAddress}
                                            placeholder="Jenis Mata Kuliah"
                                            hasIcon={false}
                                            id={'selectsubjectType'}
                                            isError={errors.subject_type ? true : false}
                                            onChange={(value)=>{
                                                field.onChange(value !== null ? [value] : [])
                                            }}
                                            />
                                    );
                                    }}
                                />
                                <FormErrorMessage errors={errors.subject_type} />
                            </Col>
                        </Row>          

                        {/*//! col subject_Group */}
                        <Row
                            className="align-items-center gap-2"
                            >
                            {/* <Col sm={12}>
                                <Label htmlFor="selectsubjectGroup" className='form-label  m-0 d-flex align-items-center gap-1'> 
                                 Jenis Mata Kuliah
                                    <span className="m-0 p-0 fst-italic fw-semibold" style={{fontSize:"10px",color:"#3A3A3A"}}>
                                    - Optional
                                    </span> 
                                </Label>
                            </Col> */}
                            <Col sm={12}>
                                <Controller
                                    name="subject_group"
                                    control={control}
                                    render={({ field }) => {
                                        return (
                                            <SelectComponent
                                            {...field}
                                            options={dummyValue!}
                                            // isLoading={isLoadingOptionCountries || isLoadingRegistrantAddress}
                                            // isDisabled={!isEdit || isLoadingRegistrantAddress}
                                            placeholder="Kelompok Mata Kuliah"
                                            hasIcon={false}
                                            id={'selectsubjectGroup'}
                                            
                                            isError={errors.subject_group ? true : false}
                                            onChange={(value)=>{
                                                field.onChange(value !== null ? [value] : [])
                                            }}
                                            />
                                    );
                                    }}
                                />
                                <FormErrorMessage errors={errors.subject_group} />
                            </Col>
                        </Row>              
                    </section>

                    {/*//! button action */}
                    <section className="d-flex gap-2 mt-4 ">
                        <Button 
                        className='px-3 py-2 flex-grow-1 border border-2'
                        color='#F3F6F9'
                        type='button'
                        onClick={handleReset}
                        >
                            Reset
                        </Button>

                        <Button type='submit' className='px-3 py-2 btn-primary flex-grow-1 gap-2'
                        color='#10487A'
                        >
                            Terapkan
                        </Button>
                    </section>
                </form>    
            </ModalBody>
    </Modal>
  )
}

export default ModalFilterSubjects