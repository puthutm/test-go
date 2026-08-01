'use client'
import React from 'react'
import { Button,Input } from 'reactstrap'
import { SearchIcon } from '@/components/icons/search'
import { FilterListIcon } from '@/components/icons/filter-list'
import { AddIcon } from '@/components/icons/add'
import DataTables from '@/components/ui/datatable'

import useColumnDefClassParticipants from './column-def-class-participants'

export interface IDummyValueTable {
  id:string,
  nim:string,
  nama:string
  presensi:string,
  hadir:string,
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
      nim:'1234567890',
      nama:'Muhammad Farhan',
      presensi:'16/16',
      hadir:'100%',
    },
    {
      id:'2',
      nim:'2347923847',
      nama:'Dipa Nusantara Aidit',
      presensi:'16/16',
      hadir:'100%',
    },
        {
      id:'3',
      nim:'8797437832',
      nama:'Kartosuwiryo',
      presensi:'16/16',
      hadir:'100%',
    },
        {
      id:'4',
      nim:'2347893242',
      nama:'Untung Syamsuri',
      presensi:'16/16',
      hadir:'100%',
    },
]
}

function SectionClassParticipants() {

    const { columns } = useColumnDefClassParticipants()

  return (
    <section className="position-relative p-3 bg-white rounded-3">
        {/*//! */}
        <h2 className="m-0 p-0 fs-5 fw-semibold">
            Peserta Kelas 
        </h2>

        {/*//! section data */}
        <section className="position-relative mt-2 ">
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
                {/*//! input search */}
                <div className='form-icon'>
                  <Input
                    autoComplete='off'
                    type="text"
                    disabled
                    name="inputSearch"
                    className={`form-control form-control-icon py-2 `}
                    id="inputSearch"
                    placeholder="Cari Peserta"
                  />

                  <i className="">
                    <SearchIcon/>
                  </i>
                </div>
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
              </div>

            </section>

            {/*//! table */}
            <section className="table-responsive mt-2">
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
        </section>
    </section>
  )
}

export default SectionClassParticipants