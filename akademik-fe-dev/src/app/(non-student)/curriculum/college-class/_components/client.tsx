'use client'
import React from "react";

// import components
import { Card, CardBody, CardHeader, Col, Row,Button,Input } from "reactstrap";
import DataTables from "@/components/ui/datatable";
import { SearchIcon } from "@/components/icons/search";
import { FilterListIcon } from "@/components/icons/filter-list";
import { RefreshIcon } from "@/components/icons/refresh";
import { FileDownloadIcon } from "@/components/icons/file-download";
import { AddIcon } from "@/components/icons/add";
import TabsSectionLecture from "./tabs";
import { SelectComponent } from "@/components/ui/select";
import useColumnDefcollegeClass from "./column-def-college-class";

export interface IDummyValueTable {
  id:string,
  periode:string,
  kurikulum:string,
  kode:string,
  nama_kelas:string,
  nama_matkul:string,
  kapasitas:string,
  peserta:string,
  dosen_pengajar:string,
}


const dummyValueTable : PaginationData<IDummyValueTable> = {
  metadata:{
    page: 1,
    size: 10,
    total_data: 10,
    total_page: 1,
  },
  data:[
    {
      id:'1',
      periode:'Akademik Genap 2024/2025',
      kurikulum:'2021',
      kode:'200001102',
      nama_kelas:'IT303',
      nama_matkul:'Sistem Operasi',
      kapasitas:'60',
      peserta:'55',
      dosen_pengajar:'Novi Dian Natashia, S.Kom., MMSI ',
    },
    {
      id:'2',
      periode:'Akademik Genap 2024/2025',
      kurikulum:'2021',
      kode:'200001102',
      nama_kelas:'IT303',
      nama_matkul:'Interaksi Manusia dan Komputer (IMK) ',
      kapasitas:'60',
      peserta:'55',
      dosen_pengajar:'Novi Dian Natashia, S.Kom., MMSI ',
    },
]
}


function ClientComponentCollegeClass() {
    // dumy data
  const dummyData = [{
    value:'All',
    label:'Semua'
  }]

  const {columns} = useColumnDefcollegeClass()

  return (
    <Row>
        <Col>
        
        <Card className='p-3 rounded-3 bg-white'>
          {/*//! card header */}
          <CardHeader className='p-0 m-0'>
            {/*//! TITLE */}
            <section className="d-flex align-items-center">
              {/* //! text */}
              <h2 className="m-0 p-0 fs-4 fw-medium flex-grow-1">
                Kelas Kuliah
              </h2>
            </section>

            {/* //! tabs */}
            <TabsSectionLecture/>

          </CardHeader>

          {/* //! card body */}
          <CardBody className="p-0 m-0">
            {/*//! action */}
            <section className='position-relative gap-2 d-flex align-items-center flex-column flex-md-row w-100 my-3'>
              {/*//! left */}
              <div className="position-relative flex-grow-1 d-flex gap-2 flex-wrap">
                {/*//! button filter */}
                <Button
                disabled
                className='border d-flex align-items-center gap-2'
                color='transparent'
                style={{color:'#909090'}}
                // onClick={()=>{
                //   setShowModalFilter(()=>({
                //     status:true,
                //     title:'Filter'
                //   }))
                // }}
                >
                  <FilterListIcon className='fs-4' color='#909090' />
                  Filter
                </Button>

                {/*//! select component */}
                <SelectComponent
                          options={dummyData}
                          placeholder="Semua"
                          isDisabled
                          // isDisabled={!isEdit || isLoadingRegistrantBiodataSystem}
                          // isError={!!errors.gender}
                          id={'selectFilter'}
                          isClearable
                          onChange={(value)=>{
                            console.log(value)
                          }}
                />

                {/*//! input search */}
                <div className='form-icon'>

                  <Input
                    autoComplete='off'
                    type="text"
                    disabled
                    name="inputSearchSubject"
                    className={`form-control form-control-icon py-2 `}
                    id="inputNamaLengkap"
                    placeholder="Cari Mata Kuliah"
                  />

                  <i className="">
                    <SearchIcon/>
                  </i>
                </div>

                {/*//! button search */}
                <Button className='btn  d-flex align-items-center gap-2'
                color='transparent'
                disabled
                style={{color:'#10487A',border:'1px solid #10487A'}}
                >
                  Cari
                </Button>

                {/*//! button clear */}
                <Button className='btn  d-flex align-items-center gap-2'
                  color='transparent'
                  disabled
                  style={{color:'#10487A',border:'1px solid #10487A'}}
                >
                  <RefreshIcon color='#10487A'/>
                </Button>
              </div>

              {/*//! right */}
              <div className='position-relative flex-shrink-0 gap-2 d-flex'>

                {/*//! button add */}
                <Button className='btn  d-flex align-items-center gap-2 '
                  color='transparent'
                  disabled
                  style={{color:'#10487A',border:'1px solid #10487A'}}
                  >
                    <AddIcon color='#10487A'/>
                    Tambah
                </Button>
              
                {/*//! button delete */}
                {/* <Button className='btn  d-flex align-items-center gap-2 border border-danger text-danger'
                    color='transparent'
                  >
                    <DeleteIcon color='#DC4C64'/>
                    Hapus
                </Button> */}

                {/*//! button downlod */}
                <Button className='btn btn-primary d-flex align-items-center gap-2'
                  color='#10487A'
                  disabled
                  >
                    <FileDownloadIcon color='#fff'/>
                    Download
                </Button>
              </div>
            </section>


            {/* //! tables */}
            <section className="table-responsive">
              <DataTables
                    columns={columns}
                    data={dummyValueTable}
                    pageCount={1}
                    pagination={dummyValueTable.metadata.page}
                    setPagination={() => {}}
                    isLoading={false}
                    total={1}
                  />
            </section>


          </CardBody>
        </Card>

        </Col>
    </Row>
  )
}

export default ClientComponentCollegeClass