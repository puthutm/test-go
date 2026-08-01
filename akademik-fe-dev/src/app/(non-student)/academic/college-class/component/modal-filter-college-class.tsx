
'use client'
import React,{Dispatch,SetStateAction, useEffect,useState} from 'react'
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

// import use form hook
import { useForm,Controller,SubmitHandler } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import schemeFormFilterCollegeClass,{IFormSchemeFilterCollegeClass} from '@/lib/validations/academic/settings/college-class/form-filter-college-class'

import { useGetUnsiaStudyProgram } from '@/services/api/data-referensi/study-program/use-get-unsia-study-program'
import { useGetOptionCurriculumYear } from '@/services/api/data-referensi/curriculum-year/use-get-option-curriculum-year'
import { useSearchParams } from 'next/navigation'
import { useRouter } from 'next/navigation'
import { usePathname } from 'next/navigation'


import { IModalManipulationFilterCollegeClass } from './client'

function ModalFilterCollegeClass({
    showModal,
    setShowModal
}:{
    showModal:IModalManipulationFilterCollegeClass,
    setShowModal:Dispatch<SetStateAction<IModalManipulationFilterCollegeClass>>,
}) {
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
    } = useForm<IFormSchemeFilterCollegeClass>({
        resolver: zodResolver(schemeFormFilterCollegeClass)
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


    //! get option study program
    const {
        data:dataOptionStudyProgram,
        isFetching:isLoadingOptionStudyProgram
    } = useGetUnsiaStudyProgram(showModal.status)


    useEffect(() => {
        if(searchParams.get('curriculum_year_id') && dataOptionCurriculumYear){
            const filterOption = dataOptionCurriculumYear?.data?.find((el:CurriculumYear)=>{
                return searchParams.get('curriculum_year_id') === el.id
            })

            const mapingOption = {
                value:filterOption?.id as string,
                label:filterOption?.years as string
            }
            setValue('kurikulum',[mapingOption])
        }
    }, [showModal,dataOptionCurriculumYear])

  const curriculumYearOptions = dataOptionCurriculumYear?.data?.map((val: CurriculumYear) => ({
    label: val.years,
    value: val.id,
  })) || [];
  const studyProgramOptions = dataOptionStudyProgram?.data?.map((val: UnsiaStudyProgram) => ({
    label: val.name,
    value: val.id,
  })) || [];

    // event handle close
    const handleCloseModal = ()=>{
            setShowModal(()=>({
            status:false,
            title:'Filter',
        }))
        reset()
        clearErrors()
        setValue('program_studi',[])
        setValue('sistem_kuliah',[])
        setValue('jenis_status',[])
        setValue('prodi_pengampu',[])
        setValue('kurikulum',[])
        setValue('kelas',[])
    }

    // handle reset
    const handleReset = ()=>{
        setValue('program_studi',[])
        setValue('sistem_kuliah',[])
        setValue('jenis_status',[])
        setValue('prodi_pengampu',[])
        setValue('kurikulum',[])
        setValue('kelas',[])
        router.push(`${pathname}`)
    }

    // handle submit
    const handleSubmitFilter:SubmitHandler<IFormSchemeFilterCollegeClass> = async (dataInput,event) => {
        event?.preventDefault()
        if(dataInput.kurikulum.length !== 0){
            params.set("curriculum_year_id", dataInput.kurikulum[0].value);
        }
        if(dataInput.program_studi.length !== 0){
            params.set("study_program_id", dataInput.program_studi[0].value);
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
                        {/*//! col program studi */}
                        <Row
                            className="align-items-center gap-2"
                            >
                            {/* <Col sm={12}>
                                <Label htmlFor="selectProgramStudi" className='form-label  m-0 d-flex align-items-center gap-1'> 
                                 Tahun Kurikulum
                                    <span className="m-0 p-0 fst-italic fw-semibold" style={{fontSize:"10px",color:"#3A3A3A"}}>
                                    - Optional
                                    </span> 
                                </Label>
                            </Col> */}
                            <Col sm={12}>
                                <Controller
                                    name="program_studi"
                                    control={control}
                                    render={({ field }) => {
                                        return (
                                            <SelectComponent
                                            {...field}
                                            options={studyProgramOptions}
                                            isLoading={isLoadingOptionStudyProgram}
                                            // isDisabled={!isEdit || isLoadingRegistrantAddress}
                                            placeholder="Program Studi"
                                            hasIcon={false}
                                            id={'selectProgramStudi'}
                                            isClearable
                                            isError={errors.program_studi ? true : false}
                                            onChange={(value)=>{
                                                field.onChange(value !== null ? [value] : [])
                                            }}
                                            />
                                    );
                                    }}
                                />
                                <FormErrorMessage errors={errors.program_studi} />
                            </Col>
                        </Row>

                        {/*//! col sistem kuliah */}
                        <Row
                            className="align-items-center gap-2"
                            >
                            {/* <Col sm={12}>
                                <Label htmlFor="selectSistemKuliah" className='form-label  m-0 d-flex align-items-center gap-1'> 
                                 Tahun Kurikulum
                                    <span className="m-0 p-0 fst-italic fw-semibold" style={{fontSize:"10px",color:"#3A3A3A"}}>
                                    - Optional
                                    </span> 
                                </Label>
                            </Col> */}
                            <Col sm={12}>
                                <Controller
                                    name="sistem_kuliah"
                                    control={control}
                                    render={({ field }) => {
                                        return (
                                            <SelectComponent
                                            {...field}
                                            isDisabled
                                            options={dummyValue!}
                                            // isLoading={isLoadingOptionCountries || isLoadingRegistrantAddress}
                                            // isDisabled={!isEdit || isLoadingRegistrantAddress}
                                            placeholder="Sistem Kuliah"
                                            hasIcon={false}
                                            id={'selectSistemKuliah'}
                                            isClearable
                                            isError={errors.sistem_kuliah ? true : false}
                                            onChange={(value)=>{
                                                field.onChange(value !== null ? [value] : [])
                                            }}
                                            />
                                    );
                                    }}
                                />
                                <FormErrorMessage errors={errors.sistem_kuliah} />
                            </Col>
                        </Row>

                        {/*//! col jenis status */}
                        <Row
                            className="align-items-center gap-2"
                            >
                            {/* <Col sm={12}>
                                <Label htmlFor="selectJenisStatus" className='form-label  m-0 d-flex align-items-center gap-1'> 
                                 Tahun Kurikulum
                                    <span className="m-0 p-0 fst-italic fw-semibold" style={{fontSize:"10px",color:"#3A3A3A"}}>
                                    - Optional
                                    </span> 
                                </Label>
                            </Col> */}
                            <Col sm={12}>
                                <Controller
                                    name="jenis_status"
                                    control={control}
                                    render={({ field }) => {
                                        return (
                                            <SelectComponent
                                            {...field}
                                            options={dummyValue}
                                            isDisabled
                                            // isLoading={isLoadingOptionStudyProgram}
                                            // isDisabled={!isEdit || isLoadingRegistrantAddress}
                                            placeholder="Jenis status"
                                            hasIcon={false}
                                            id={'selectJenisStatus'}
                                            isClearable
                                            isError={errors.jenis_status ? true : false}
                                            onChange={(value)=>{
                                                field.onChange(value !== null ? [value] : [])
                                            }}
                                            />
                                    );
                                    }}
                                />
                                <FormErrorMessage errors={errors.jenis_status} />
                            </Col>
                        </Row>

                        {/*//! col prodi pengampu */}
                        <Row
                            className="align-items-center gap-2"
                            >
                            {/* <Col sm={12}>
                                <Label htmlFor="selectProdiPengampu" className='form-label  m-0 d-flex align-items-center gap-1'> 
                                 Tahun Kurikulum
                                    <span className="m-0 p-0 fst-italic fw-semibold" style={{fontSize:"10px",color:"#3A3A3A"}}>
                                    - Optional
                                    </span> 
                                </Label>
                            </Col> */}
                            <Col sm={12}>
                                <Controller
                                    name="prodi_pengampu"
                                    control={control}
                                    render={({ field }) => {
                                        return (
                                            <SelectComponent
                                            {...field}
                                            isDisabled
                                            options={dummyValue!}
                                            // isLoading={isLoadingOptionCountries || isLoadingRegistrantAddress}
                                            // isDisabled={!isEdit || isLoadingRegistrantAddress}
                                            placeholder="Jenis status"
                                            hasIcon={false}
                                            id={'selectProdiPengampu'}
                                            isClearable
                                            isError={errors.prodi_pengampu ? true : false}
                                            onChange={(value)=>{
                                                field.onChange(value !== null ? [value] : [])
                                            }}
                                            />
                                    );
                                    }}
                                />
                                <FormErrorMessage errors={errors.prodi_pengampu} />
                            </Col>
                        </Row>


                        {/*//! col kurikulum */}
                        <Row
                            className="align-items-center gap-2"
                            >
                            {/* <Col sm={12}>
                                <Label htmlFor="selectKurikulum" className='form-label  m-0 d-flex align-items-center gap-1'> 
                                 Tahun Kurikulum
                                    <span className="m-0 p-0 fst-italic fw-semibold" style={{fontSize:"10px",color:"#3A3A3A"}}>
                                    - Optional
                                    </span> 
                                </Label>
                            </Col> */}
                            <Col sm={12}>
                                <Controller
                                    name="kurikulum"
                                    control={control}
                                    render={({ field }) => {
                                        return (
                                            <SelectComponent
                                            {...field}
                                            options={curriculumYearOptions}
                                            isLoading={isLoadingCurriculumYear}
                                            // isDisabled={!isEdit || isLoadingRegistrantAddress}
                                            placeholder="Kurikulum"
                                            hasIcon={false}
                                            id={'selectKurikulum'}
                                            isError={errors.kurikulum ? true : false}
                                            onChange={(value)=>{
                                                field.onChange(value !== null ? [value] : [])
                                            }}
                                            />
                                    );
                                    }}
                                />
                                <FormErrorMessage errors={errors.kurikulum} />
                            </Col>
                        </Row>

                        {/*//! col kelas / kelompok */}
                        <Row
                            className="align-items-center gap-2"
                            >
                            {/* <Col sm={12}>
                                <Label htmlFor="selectKelas" className='form-label  m-0 d-flex align-items-center gap-1'> 
                                 Tahun Kurikulum
                                    <span className="m-0 p-0 fst-italic fw-semibold" style={{fontSize:"10px",color:"#3A3A3A"}}>
                                    - Optional
                                    </span> 
                                </Label>
                            </Col> */}
                            <Col sm={12}>
                                <Controller
                                    name="kelas"
                                    control={control}
                                    render={({ field }) => {
                                        return (
                                            <SelectComponent
                                            {...field}
                                            isDisabled
                                            options={dummyValue!}
                                            // isLoading={isLoadingOptionCountries || isLoadingRegistrantAddress}
                                            // isDisabled={!isEdit || isLoadingRegistrantAddress}
                                            placeholder="Kelas/Kelompok"
                                            hasIcon={false}
                                            id={'selectKelas'}
                                            isClearable
                                            isError={errors.kelas ? true : false}
                                            onChange={(value)=>{
                                                field.onChange(value !== null ? [value] : [])
                                            }}
                                            />
                                    );
                                    }}
                                />
                                <FormErrorMessage errors={errors.kelas} />
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

                        <Button className='px-3 py-2 btn-primary flex-grow-1 gap-2'
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

export default ModalFilterCollegeClass